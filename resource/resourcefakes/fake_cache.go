// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeCache struct {
	IsInitializedStub        func() (bool, error)
	isInitializedMutex       sync.RWMutex
	isInitializedArgsForCall []struct{}
	isInitializedReturns     struct {
		result1 bool
		result2 error
	}
	isInitializedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	InitializeStub        func() error
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct{}
	initializeReturns     struct {
		result1 error
	}
	initializeReturnsOnCall map[int]struct {
		result1 error
	}
	VolumeStub        func() worker.Volume
	volumeMutex       sync.RWMutex
	volumeArgsForCall []struct{}
	volumeReturns     struct {
		result1 worker.Volume
	}
	volumeReturnsOnCall map[int]struct {
		result1 worker.Volume
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCache) IsInitialized() (bool, error) {
	fake.isInitializedMutex.Lock()
	ret, specificReturn := fake.isInitializedReturnsOnCall[len(fake.isInitializedArgsForCall)]
	fake.isInitializedArgsForCall = append(fake.isInitializedArgsForCall, struct{}{})
	fake.recordInvocation("IsInitialized", []interface{}{})
	fake.isInitializedMutex.Unlock()
	if fake.IsInitializedStub != nil {
		return fake.IsInitializedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isInitializedReturns.result1, fake.isInitializedReturns.result2
}

func (fake *FakeCache) IsInitializedCallCount() int {
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	return len(fake.isInitializedArgsForCall)
}

func (fake *FakeCache) IsInitializedReturns(result1 bool, result2 error) {
	fake.IsInitializedStub = nil
	fake.isInitializedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCache) IsInitializedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsInitializedStub = nil
	if fake.isInitializedReturnsOnCall == nil {
		fake.isInitializedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isInitializedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCache) Initialize() error {
	fake.initializeMutex.Lock()
	ret, specificReturn := fake.initializeReturnsOnCall[len(fake.initializeArgsForCall)]
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct{}{})
	fake.recordInvocation("Initialize", []interface{}{})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		return fake.InitializeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.initializeReturns.result1
}

func (fake *FakeCache) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *FakeCache) InitializeReturns(result1 error) {
	fake.InitializeStub = nil
	fake.initializeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCache) InitializeReturnsOnCall(i int, result1 error) {
	fake.InitializeStub = nil
	if fake.initializeReturnsOnCall == nil {
		fake.initializeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCache) Volume() worker.Volume {
	fake.volumeMutex.Lock()
	ret, specificReturn := fake.volumeReturnsOnCall[len(fake.volumeArgsForCall)]
	fake.volumeArgsForCall = append(fake.volumeArgsForCall, struct{}{})
	fake.recordInvocation("Volume", []interface{}{})
	fake.volumeMutex.Unlock()
	if fake.VolumeStub != nil {
		return fake.VolumeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.volumeReturns.result1
}

func (fake *FakeCache) VolumeCallCount() int {
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	return len(fake.volumeArgsForCall)
}

func (fake *FakeCache) VolumeReturns(result1 worker.Volume) {
	fake.VolumeStub = nil
	fake.volumeReturns = struct {
		result1 worker.Volume
	}{result1}
}

func (fake *FakeCache) VolumeReturnsOnCall(i int, result1 worker.Volume) {
	fake.VolumeStub = nil
	if fake.volumeReturnsOnCall == nil {
		fake.volumeReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
		})
	}
	fake.volumeReturnsOnCall[i] = struct {
		result1 worker.Volume
	}{result1}
}

func (fake *FakeCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isInitializedMutex.RLock()
	defer fake.isInitializedMutex.RUnlock()
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	fake.volumeMutex.RLock()
	defer fake.volumeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCache) recordInvocation(key string, args []interface{}) {
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

var _ resource.Cache = new(FakeCache)
