// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
)

type FakeResourceConfigFactory struct {
	FindOrCreateResourceConfigStub        func(logger lager.Logger, user dbng.ResourceUser, resourceType string, source atc.Source, pipelineID int, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error)
	findOrCreateResourceConfigMutex       sync.RWMutex
	findOrCreateResourceConfigArgsForCall []struct {
		logger        lager.Logger
		user          dbng.ResourceUser
		resourceType  string
		source        atc.Source
		pipelineID    int
		resourceTypes atc.ResourceTypes
	}
	findOrCreateResourceConfigReturns struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}
	findOrCreateResourceConfigReturnsOnCall map[int]struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}
	CleanConfigUsesForFinishedBuildsStub        func() error
	cleanConfigUsesForFinishedBuildsMutex       sync.RWMutex
	cleanConfigUsesForFinishedBuildsArgsForCall []struct{}
	cleanConfigUsesForFinishedBuildsReturns     struct {
		result1 error
	}
	cleanConfigUsesForFinishedBuildsReturnsOnCall map[int]struct {
		result1 error
	}
	CleanConfigUsesForInactiveResourceTypesStub        func() error
	cleanConfigUsesForInactiveResourceTypesMutex       sync.RWMutex
	cleanConfigUsesForInactiveResourceTypesArgsForCall []struct{}
	cleanConfigUsesForInactiveResourceTypesReturns     struct {
		result1 error
	}
	cleanConfigUsesForInactiveResourceTypesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanConfigUsesForInactiveResourcesStub        func() error
	cleanConfigUsesForInactiveResourcesMutex       sync.RWMutex
	cleanConfigUsesForInactiveResourcesArgsForCall []struct{}
	cleanConfigUsesForInactiveResourcesReturns     struct {
		result1 error
	}
	cleanConfigUsesForInactiveResourcesReturnsOnCall map[int]struct {
		result1 error
	}
	CleanUselessConfigsStub        func() error
	cleanUselessConfigsMutex       sync.RWMutex
	cleanUselessConfigsArgsForCall []struct{}
	cleanUselessConfigsReturns     struct {
		result1 error
	}
	cleanUselessConfigsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfig(logger lager.Logger, user dbng.ResourceUser, resourceType string, source atc.Source, pipelineID int, resourceTypes atc.ResourceTypes) (*dbng.UsedResourceConfig, error) {
	fake.findOrCreateResourceConfigMutex.Lock()
	ret, specificReturn := fake.findOrCreateResourceConfigReturnsOnCall[len(fake.findOrCreateResourceConfigArgsForCall)]
	fake.findOrCreateResourceConfigArgsForCall = append(fake.findOrCreateResourceConfigArgsForCall, struct {
		logger        lager.Logger
		user          dbng.ResourceUser
		resourceType  string
		source        atc.Source
		pipelineID    int
		resourceTypes atc.ResourceTypes
	}{logger, user, resourceType, source, pipelineID, resourceTypes})
	fake.recordInvocation("FindOrCreateResourceConfig", []interface{}{logger, user, resourceType, source, pipelineID, resourceTypes})
	fake.findOrCreateResourceConfigMutex.Unlock()
	if fake.FindOrCreateResourceConfigStub != nil {
		return fake.FindOrCreateResourceConfigStub(logger, user, resourceType, source, pipelineID, resourceTypes)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateResourceConfigReturns.result1, fake.findOrCreateResourceConfigReturns.result2
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigCallCount() int {
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	return len(fake.findOrCreateResourceConfigArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigArgsForCall(i int) (lager.Logger, dbng.ResourceUser, string, atc.Source, int, atc.ResourceTypes) {
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	return fake.findOrCreateResourceConfigArgsForCall[i].logger, fake.findOrCreateResourceConfigArgsForCall[i].user, fake.findOrCreateResourceConfigArgsForCall[i].resourceType, fake.findOrCreateResourceConfigArgsForCall[i].source, fake.findOrCreateResourceConfigArgsForCall[i].pipelineID, fake.findOrCreateResourceConfigArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigReturns(result1 *dbng.UsedResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigStub = nil
	fake.findOrCreateResourceConfigReturns = struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigReturnsOnCall(i int, result1 *dbng.UsedResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigStub = nil
	if fake.findOrCreateResourceConfigReturnsOnCall == nil {
		fake.findOrCreateResourceConfigReturnsOnCall = make(map[int]struct {
			result1 *dbng.UsedResourceConfig
			result2 error
		})
	}
	fake.findOrCreateResourceConfigReturnsOnCall[i] = struct {
		result1 *dbng.UsedResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuilds() error {
	fake.cleanConfigUsesForFinishedBuildsMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForFinishedBuildsReturnsOnCall[len(fake.cleanConfigUsesForFinishedBuildsArgsForCall)]
	fake.cleanConfigUsesForFinishedBuildsArgsForCall = append(fake.cleanConfigUsesForFinishedBuildsArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForFinishedBuilds", []interface{}{})
	fake.cleanConfigUsesForFinishedBuildsMutex.Unlock()
	if fake.CleanConfigUsesForFinishedBuildsStub != nil {
		return fake.CleanConfigUsesForFinishedBuildsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForFinishedBuildsReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuildsCallCount() int {
	fake.cleanConfigUsesForFinishedBuildsMutex.RLock()
	defer fake.cleanConfigUsesForFinishedBuildsMutex.RUnlock()
	return len(fake.cleanConfigUsesForFinishedBuildsArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuildsReturns(result1 error) {
	fake.CleanConfigUsesForFinishedBuildsStub = nil
	fake.cleanConfigUsesForFinishedBuildsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForFinishedBuildsReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForFinishedBuildsStub = nil
	if fake.cleanConfigUsesForFinishedBuildsReturnsOnCall == nil {
		fake.cleanConfigUsesForFinishedBuildsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForFinishedBuildsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypes() error {
	fake.cleanConfigUsesForInactiveResourceTypesMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall[len(fake.cleanConfigUsesForInactiveResourceTypesArgsForCall)]
	fake.cleanConfigUsesForInactiveResourceTypesArgsForCall = append(fake.cleanConfigUsesForInactiveResourceTypesArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForInactiveResourceTypes", []interface{}{})
	fake.cleanConfigUsesForInactiveResourceTypesMutex.Unlock()
	if fake.CleanConfigUsesForInactiveResourceTypesStub != nil {
		return fake.CleanConfigUsesForInactiveResourceTypesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForInactiveResourceTypesReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypesCallCount() int {
	fake.cleanConfigUsesForInactiveResourceTypesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourceTypesMutex.RUnlock()
	return len(fake.cleanConfigUsesForInactiveResourceTypesArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypesReturns(result1 error) {
	fake.CleanConfigUsesForInactiveResourceTypesStub = nil
	fake.cleanConfigUsesForInactiveResourceTypesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourceTypesReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForInactiveResourceTypesStub = nil
	if fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall == nil {
		fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForInactiveResourceTypesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResources() error {
	fake.cleanConfigUsesForInactiveResourcesMutex.Lock()
	ret, specificReturn := fake.cleanConfigUsesForInactiveResourcesReturnsOnCall[len(fake.cleanConfigUsesForInactiveResourcesArgsForCall)]
	fake.cleanConfigUsesForInactiveResourcesArgsForCall = append(fake.cleanConfigUsesForInactiveResourcesArgsForCall, struct{}{})
	fake.recordInvocation("CleanConfigUsesForInactiveResources", []interface{}{})
	fake.cleanConfigUsesForInactiveResourcesMutex.Unlock()
	if fake.CleanConfigUsesForInactiveResourcesStub != nil {
		return fake.CleanConfigUsesForInactiveResourcesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanConfigUsesForInactiveResourcesReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourcesCallCount() int {
	fake.cleanConfigUsesForInactiveResourcesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourcesMutex.RUnlock()
	return len(fake.cleanConfigUsesForInactiveResourcesArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourcesReturns(result1 error) {
	fake.CleanConfigUsesForInactiveResourcesStub = nil
	fake.cleanConfigUsesForInactiveResourcesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanConfigUsesForInactiveResourcesReturnsOnCall(i int, result1 error) {
	fake.CleanConfigUsesForInactiveResourcesStub = nil
	if fake.cleanConfigUsesForInactiveResourcesReturnsOnCall == nil {
		fake.cleanConfigUsesForInactiveResourcesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanConfigUsesForInactiveResourcesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigs() error {
	fake.cleanUselessConfigsMutex.Lock()
	ret, specificReturn := fake.cleanUselessConfigsReturnsOnCall[len(fake.cleanUselessConfigsArgsForCall)]
	fake.cleanUselessConfigsArgsForCall = append(fake.cleanUselessConfigsArgsForCall, struct{}{})
	fake.recordInvocation("CleanUselessConfigs", []interface{}{})
	fake.cleanUselessConfigsMutex.Unlock()
	if fake.CleanUselessConfigsStub != nil {
		return fake.CleanUselessConfigsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUselessConfigsReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigsCallCount() int {
	fake.cleanUselessConfigsMutex.RLock()
	defer fake.cleanUselessConfigsMutex.RUnlock()
	return len(fake.cleanUselessConfigsArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigsReturns(result1 error) {
	fake.CleanUselessConfigsStub = nil
	fake.cleanUselessConfigsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanUselessConfigsReturnsOnCall(i int, result1 error) {
	fake.CleanUselessConfigsStub = nil
	if fake.cleanUselessConfigsReturnsOnCall == nil {
		fake.cleanUselessConfigsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUselessConfigsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	fake.cleanConfigUsesForFinishedBuildsMutex.RLock()
	defer fake.cleanConfigUsesForFinishedBuildsMutex.RUnlock()
	fake.cleanConfigUsesForInactiveResourceTypesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourceTypesMutex.RUnlock()
	fake.cleanConfigUsesForInactiveResourcesMutex.RLock()
	defer fake.cleanConfigUsesForInactiveResourcesMutex.RUnlock()
	fake.cleanUselessConfigsMutex.RLock()
	defer fake.cleanUselessConfigsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResourceConfigFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dbng.ResourceConfigFactory = new(FakeResourceConfigFactory)
