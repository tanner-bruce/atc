// This file was generated by counterfeiter
package maxinflightfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/scheduler/maxinflight"
)

type FakeUpdater struct {
	UpdateMaxInFlightReachedStub        func(logger lager.Logger, jobConfig atc.JobConfig, buildID int) (bool, error)
	updateMaxInFlightReachedMutex       sync.RWMutex
	updateMaxInFlightReachedArgsForCall []struct {
		logger    lager.Logger
		jobConfig atc.JobConfig
		buildID   int
	}
	updateMaxInFlightReachedReturns struct {
		result1 bool
		result2 error
	}
	updateMaxInFlightReachedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUpdater) UpdateMaxInFlightReached(logger lager.Logger, jobConfig atc.JobConfig, buildID int) (bool, error) {
	fake.updateMaxInFlightReachedMutex.Lock()
	ret, specificReturn := fake.updateMaxInFlightReachedReturnsOnCall[len(fake.updateMaxInFlightReachedArgsForCall)]
	fake.updateMaxInFlightReachedArgsForCall = append(fake.updateMaxInFlightReachedArgsForCall, struct {
		logger    lager.Logger
		jobConfig atc.JobConfig
		buildID   int
	}{logger, jobConfig, buildID})
	fake.recordInvocation("UpdateMaxInFlightReached", []interface{}{logger, jobConfig, buildID})
	fake.updateMaxInFlightReachedMutex.Unlock()
	if fake.UpdateMaxInFlightReachedStub != nil {
		return fake.UpdateMaxInFlightReachedStub(logger, jobConfig, buildID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateMaxInFlightReachedReturns.result1, fake.updateMaxInFlightReachedReturns.result2
}

func (fake *FakeUpdater) UpdateMaxInFlightReachedCallCount() int {
	fake.updateMaxInFlightReachedMutex.RLock()
	defer fake.updateMaxInFlightReachedMutex.RUnlock()
	return len(fake.updateMaxInFlightReachedArgsForCall)
}

func (fake *FakeUpdater) UpdateMaxInFlightReachedArgsForCall(i int) (lager.Logger, atc.JobConfig, int) {
	fake.updateMaxInFlightReachedMutex.RLock()
	defer fake.updateMaxInFlightReachedMutex.RUnlock()
	return fake.updateMaxInFlightReachedArgsForCall[i].logger, fake.updateMaxInFlightReachedArgsForCall[i].jobConfig, fake.updateMaxInFlightReachedArgsForCall[i].buildID
}

func (fake *FakeUpdater) UpdateMaxInFlightReachedReturns(result1 bool, result2 error) {
	fake.UpdateMaxInFlightReachedStub = nil
	fake.updateMaxInFlightReachedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdater) UpdateMaxInFlightReachedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.UpdateMaxInFlightReachedStub = nil
	if fake.updateMaxInFlightReachedReturnsOnCall == nil {
		fake.updateMaxInFlightReachedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.updateMaxInFlightReachedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeUpdater) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateMaxInFlightReachedMutex.RLock()
	defer fake.updateMaxInFlightReachedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeUpdater) recordInvocation(key string, args []interface{}) {
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

var _ maxinflight.Updater = new(FakeUpdater)
