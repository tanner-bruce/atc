// This file was generated by counterfeiter
package gcngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeContainerFactory struct {
	MarkContainersForDeletionStub        func() error
	markContainersForDeletionMutex       sync.RWMutex
	markContainersForDeletionArgsForCall []struct{}
	markContainersForDeletionReturns     struct {
		result1 error
	}
	markContainersForDeletionReturnsOnCall map[int]struct {
		result1 error
	}
	FindContainersMarkedForDeletionStub        func() ([]dbng.DestroyingContainer, error)
	findContainersMarkedForDeletionMutex       sync.RWMutex
	findContainersMarkedForDeletionArgsForCall []struct{}
	findContainersMarkedForDeletionReturns     struct {
		result1 []dbng.DestroyingContainer
		result2 error
	}
	findContainersMarkedForDeletionReturnsOnCall map[int]struct {
		result1 []dbng.DestroyingContainer
		result2 error
	}
	FindHijackedContainersForDeletionStub        func() ([]dbng.CreatedContainer, error)
	findHijackedContainersForDeletionMutex       sync.RWMutex
	findHijackedContainersForDeletionArgsForCall []struct{}
	findHijackedContainersForDeletionReturns     struct {
		result1 []dbng.CreatedContainer
		result2 error
	}
	findHijackedContainersForDeletionReturnsOnCall map[int]struct {
		result1 []dbng.CreatedContainer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerFactory) MarkContainersForDeletion() error {
	fake.markContainersForDeletionMutex.Lock()
	ret, specificReturn := fake.markContainersForDeletionReturnsOnCall[len(fake.markContainersForDeletionArgsForCall)]
	fake.markContainersForDeletionArgsForCall = append(fake.markContainersForDeletionArgsForCall, struct{}{})
	fake.recordInvocation("MarkContainersForDeletion", []interface{}{})
	fake.markContainersForDeletionMutex.Unlock()
	if fake.MarkContainersForDeletionStub != nil {
		return fake.MarkContainersForDeletionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.markContainersForDeletionReturns.result1
}

func (fake *FakeContainerFactory) MarkContainersForDeletionCallCount() int {
	fake.markContainersForDeletionMutex.RLock()
	defer fake.markContainersForDeletionMutex.RUnlock()
	return len(fake.markContainersForDeletionArgsForCall)
}

func (fake *FakeContainerFactory) MarkContainersForDeletionReturns(result1 error) {
	fake.MarkContainersForDeletionStub = nil
	fake.markContainersForDeletionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerFactory) MarkContainersForDeletionReturnsOnCall(i int, result1 error) {
	fake.MarkContainersForDeletionStub = nil
	if fake.markContainersForDeletionReturnsOnCall == nil {
		fake.markContainersForDeletionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.markContainersForDeletionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerFactory) FindContainersMarkedForDeletion() ([]dbng.DestroyingContainer, error) {
	fake.findContainersMarkedForDeletionMutex.Lock()
	ret, specificReturn := fake.findContainersMarkedForDeletionReturnsOnCall[len(fake.findContainersMarkedForDeletionArgsForCall)]
	fake.findContainersMarkedForDeletionArgsForCall = append(fake.findContainersMarkedForDeletionArgsForCall, struct{}{})
	fake.recordInvocation("FindContainersMarkedForDeletion", []interface{}{})
	fake.findContainersMarkedForDeletionMutex.Unlock()
	if fake.FindContainersMarkedForDeletionStub != nil {
		return fake.FindContainersMarkedForDeletionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findContainersMarkedForDeletionReturns.result1, fake.findContainersMarkedForDeletionReturns.result2
}

func (fake *FakeContainerFactory) FindContainersMarkedForDeletionCallCount() int {
	fake.findContainersMarkedForDeletionMutex.RLock()
	defer fake.findContainersMarkedForDeletionMutex.RUnlock()
	return len(fake.findContainersMarkedForDeletionArgsForCall)
}

func (fake *FakeContainerFactory) FindContainersMarkedForDeletionReturns(result1 []dbng.DestroyingContainer, result2 error) {
	fake.FindContainersMarkedForDeletionStub = nil
	fake.findContainersMarkedForDeletionReturns = struct {
		result1 []dbng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerFactory) FindContainersMarkedForDeletionReturnsOnCall(i int, result1 []dbng.DestroyingContainer, result2 error) {
	fake.FindContainersMarkedForDeletionStub = nil
	if fake.findContainersMarkedForDeletionReturnsOnCall == nil {
		fake.findContainersMarkedForDeletionReturnsOnCall = make(map[int]struct {
			result1 []dbng.DestroyingContainer
			result2 error
		})
	}
	fake.findContainersMarkedForDeletionReturnsOnCall[i] = struct {
		result1 []dbng.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerFactory) FindHijackedContainersForDeletion() ([]dbng.CreatedContainer, error) {
	fake.findHijackedContainersForDeletionMutex.Lock()
	ret, specificReturn := fake.findHijackedContainersForDeletionReturnsOnCall[len(fake.findHijackedContainersForDeletionArgsForCall)]
	fake.findHijackedContainersForDeletionArgsForCall = append(fake.findHijackedContainersForDeletionArgsForCall, struct{}{})
	fake.recordInvocation("FindHijackedContainersForDeletion", []interface{}{})
	fake.findHijackedContainersForDeletionMutex.Unlock()
	if fake.FindHijackedContainersForDeletionStub != nil {
		return fake.FindHijackedContainersForDeletionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findHijackedContainersForDeletionReturns.result1, fake.findHijackedContainersForDeletionReturns.result2
}

func (fake *FakeContainerFactory) FindHijackedContainersForDeletionCallCount() int {
	fake.findHijackedContainersForDeletionMutex.RLock()
	defer fake.findHijackedContainersForDeletionMutex.RUnlock()
	return len(fake.findHijackedContainersForDeletionArgsForCall)
}

func (fake *FakeContainerFactory) FindHijackedContainersForDeletionReturns(result1 []dbng.CreatedContainer, result2 error) {
	fake.FindHijackedContainersForDeletionStub = nil
	fake.findHijackedContainersForDeletionReturns = struct {
		result1 []dbng.CreatedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerFactory) FindHijackedContainersForDeletionReturnsOnCall(i int, result1 []dbng.CreatedContainer, result2 error) {
	fake.FindHijackedContainersForDeletionStub = nil
	if fake.findHijackedContainersForDeletionReturnsOnCall == nil {
		fake.findHijackedContainersForDeletionReturnsOnCall = make(map[int]struct {
			result1 []dbng.CreatedContainer
			result2 error
		})
	}
	fake.findHijackedContainersForDeletionReturnsOnCall[i] = struct {
		result1 []dbng.CreatedContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.markContainersForDeletionMutex.RLock()
	defer fake.markContainersForDeletionMutex.RUnlock()
	fake.findContainersMarkedForDeletionMutex.RLock()
	defer fake.findContainersMarkedForDeletionMutex.RUnlock()
	fake.findHijackedContainersForDeletionMutex.RLock()
	defer fake.findHijackedContainersForDeletionMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeContainerFactory) recordInvocation(key string, args []interface{}) {
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
