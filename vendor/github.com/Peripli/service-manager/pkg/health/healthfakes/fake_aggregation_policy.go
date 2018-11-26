// Code generated by counterfeiter. DO NOT EDIT.
package healthfakes

import (
	sync "sync"

	health "github.com/Peripli/service-manager/pkg/health"
)

type FakeAggregationPolicy struct {
	ApplyStub        func(map[string]*health.Health) *health.Health
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 map[string]*health.Health
	}
	applyReturns struct {
		result1 *health.Health
	}
	applyReturnsOnCall map[int]struct {
		result1 *health.Health
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAggregationPolicy) Apply(arg1 map[string]*health.Health) *health.Health {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 map[string]*health.Health
	}{arg1})
	fake.recordInvocation("Apply", []interface{}{arg1})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.applyReturns
	return fakeReturns.result1
}

func (fake *FakeAggregationPolicy) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeAggregationPolicy) ApplyCalls(stub func(map[string]*health.Health) *health.Health) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeAggregationPolicy) ApplyArgsForCall(i int) map[string]*health.Health {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAggregationPolicy) ApplyReturns(result1 *health.Health) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 *health.Health
	}{result1}
}

func (fake *FakeAggregationPolicy) ApplyReturnsOnCall(i int, result1 *health.Health) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 *health.Health
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 *health.Health
	}{result1}
}

func (fake *FakeAggregationPolicy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAggregationPolicy) recordInvocation(key string, args []interface{}) {
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

var _ health.AggregationPolicy = new(FakeAggregationPolicy)
