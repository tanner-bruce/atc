package ssm_test

import (
	"os"

	"github.com/concourse/atc/creds/ssm"
	flags "github.com/jessevdk/go-flags"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("SsmManager", func() {
	var manager ssm.SsmManager

	Describe("IsConfigured()", func() {
		JustBeforeEach(func() {
			_, err := flags.ParseArgs(&manager, []string{})
			Expect(err).To(BeNil())
		})

		It("failes on empty SsmManager", func() {
			Expect(manager.IsConfigured()).To(BeFalse())
		})

		It("passes if AwsRegion is set", func() {
			manager.AwsRegion = "test-region"
			Expect(manager.IsConfigured()).To(BeTrue())
		})

		It("passes if AWS_REGION environment is set", func() {
			os.Setenv("AWS_REGION", "env-region")
			Expect(manager.IsConfigured()).To(BeTrue())
		})
	})

	Describe("Validate()", func() {
		JustBeforeEach(func() {
			manager = ssm.SsmManager{AwsRegion: "test-region"}
			_, err := flags.ParseArgs(&manager, []string{})
			Expect(err).To(BeNil())
			Expect(manager.SecretTemplate).To(Equal(ssm.DefaultSecretTemplate))
			Expect(manager.FallbackTemplate).To(Equal(ssm.DefaultFallbackTemplate))
		})

		It("passes on default parameters", func() {
			Expect(manager.Validate()).To(BeNil())
		})

		It("passes if all aws credentials are specified", func() {
			manager.AwsAccessKeyID = "access"
			manager.AwsSecretAccessKey = "secret"
			manager.AwsSessionToken = "token"
			Expect(manager.Validate()).To(BeNil())
		})

		DescribeTable("fails on partial AWS credentials",
			func(accessKey, secretKey, sessionToken string) {
				manager.AwsAccessKeyID = accessKey
				manager.AwsSecretAccessKey = secretKey
				manager.AwsSessionToken = sessionToken
				Expect(manager.Validate()).ToNot(BeNil())
			},
			Entry("only access", "access", "", ""),
			Entry("access & secret", "access", "secret", ""),
			Entry("access & token", "access", "", "token"),
			Entry("only secret", "", "secret", ""),
			Entry("secret & token", "", "secret", "token"),
			Entry("only token", "", "", "token"),
		)

		It("passes on secret template containing less specialization", func() {
			manager.SecretTemplate = "{{.Secret}}"
			Expect(manager.Validate()).To(BeNil())
		})

		It("passes on secret template containing no specialization", func() {
			manager.SecretTemplate = "var"
			Expect(manager.Validate()).To(BeNil())
		})

		It("fails on empty secret template", func() {
			manager.FallbackTemplate = ""
			Expect(manager.Validate()).To(BeNil())
		})

		It("fails on secret template containing invalid parameters", func() {
			manager.SecretTemplate = "{{.Teams}}"
			Expect(manager.Validate()).ToNot(BeNil())
		})

		It("passes on fallback template containing less specialization", func() {
			manager.FallbackTemplate = "{{.Secret}}"
			Expect(manager.Validate()).To(BeNil())
		})

		It("passes on empty fallback template", func() {
			manager.FallbackTemplate = ""
			Expect(manager.Validate()).To(BeNil())
		})

		It("fails on fallback template containing invalid parameters", func() {
			manager.FallbackTemplate = "{{.Teams}}"
			Expect(manager.Validate()).ToNot(BeNil())
		})
	})
})
