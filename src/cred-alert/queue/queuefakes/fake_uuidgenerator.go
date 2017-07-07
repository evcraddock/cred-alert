// Code generated by counterfeiter. DO NOT EDIT.
package queuefakes

import (
	"cred-alert/queue"
	"sync"
)

type FakeUUIDGenerator struct {
	GenerateStub        func() string
	generateMutex       sync.RWMutex
	generateArgsForCall []struct{}
	generateReturns     struct {
		result1 string
	}
	generateReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUUIDGenerator) Generate() string {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct{}{})
	fake.recordInvocation("Generate", []interface{}{})
	fake.generateMutex.Unlock()
	if fake.GenerateStub != nil {
		return fake.GenerateStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.generateReturns.result1
}

func (fake *FakeUUIDGenerator) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *FakeUUIDGenerator) GenerateReturns(result1 string) {
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUUIDGenerator) GenerateReturnsOnCall(i int, result1 string) {
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUUIDGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUUIDGenerator) recordInvocation(key string, args []interface{}) {
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

var _ queue.UUIDGenerator = new(FakeUUIDGenerator)
