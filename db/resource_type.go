package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/cloudfoundry/bosh-cli/director/template"
	"github.com/concourse/atc"
	"github.com/concourse/atc/creds"
)

type ResourceTypeNotFoundError struct {
	Name string
}

func (e ResourceTypeNotFoundError) Error() string {
	return fmt.Sprintf("resource type not found: %s", e.Name)
}

//go:generate counterfeiter . ResourceType

type ResourceType interface {
	ID() int
	Name() string
	Type() string
	Privileged() bool
	Source() atc.Source

	Version() atc.Version
	SaveVersion(atc.Version) error

	Reload() (bool, error)
}

type ResourceTypes []ResourceType

func (resourceTypes ResourceTypes) Deserialize(variablesSource template.Variables) (atc.VersionedResourceTypes, error) {
	var versionedResourceTypes atc.VersionedResourceTypes

	for _, t := range resourceTypes {
		source, err := creds.NewSource(variablesSource, t.Source()).Evaluate()
		if err != nil {
			return nil, err
		}

		versionedResourceTypes = append(versionedResourceTypes, atc.VersionedResourceType{
			ResourceType: atc.ResourceType{
				Name:       t.Name(),
				Type:       t.Type(),
				Source:     source,
				Privileged: t.Privileged(),
			},
			Version: t.Version(),
		})
	}

	return versionedResourceTypes, nil
}

var resourceTypesQuery = psql.Select("id, name, type, config, version, nonce").
	From("resource_types").
	Where(sq.Eq{"active": true})

type resourceType struct {
	id         int
	name       string
	type_      string
	privileged bool
	source     atc.Source
	version    atc.Version

	conn Conn
}

func (t *resourceType) ID() int            { return t.id }
func (t *resourceType) Name() string       { return t.name }
func (t *resourceType) Type() string       { return t.type_ }
func (t *resourceType) Privileged() bool   { return t.privileged }
func (t *resourceType) Source() atc.Source { return t.source }

func (t *resourceType) Version() atc.Version { return t.version }
func (t *resourceType) SaveVersion(version atc.Version) error {
	versionJSON, err := json.Marshal(version)
	if err != nil {
		return err
	}

	result, err := psql.Update("resource_types").
		Where(sq.Eq{"id": t.id}).
		Set("version", versionJSON).
		RunWith(t.conn).
		Exec()
	if err != nil {
		return err
	}

	num, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if num == 0 {
		return ResourceTypeNotFoundError{t.name}
	}

	return nil
}

func (t *resourceType) Reload() (bool, error) {
	row := resourceTypesQuery.Where(sq.Eq{"id": t.id}).RunWith(t.conn).QueryRow()

	err := scanResourceType(t, row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func scanResourceType(t *resourceType, row scannable) error {
	var (
		configJSON     []byte
		version, nonce sql.NullString
	)

	err := row.Scan(&t.id, &t.name, &t.type_, &configJSON, &version, &nonce)
	if err != nil {
		return err
	}

	if version.Valid {
		err := json.Unmarshal([]byte(version.String), &t.version)
		if err != nil {
			return err
		}
	}

	es := t.conn.EncryptionStrategy()

	var noncense *string
	if nonce.Valid {
		noncense = &nonce.String
	}

	decryptedConfig, err := es.Decrypt(string(configJSON), noncense)
	if err != nil {
		return err
	}

	var config atc.ResourceType
	err = json.Unmarshal(decryptedConfig, &config)
	if err != nil {
		return err
	}

	t.source = config.Source
	t.privileged = config.Privileged

	return nil
}
