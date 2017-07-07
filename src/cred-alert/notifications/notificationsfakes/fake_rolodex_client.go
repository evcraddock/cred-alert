// Code generated by counterfeiter. DO NOT EDIT.
package notificationsfakes

import (
	"cred-alert/notifications"
	"rolodex/rolodexpb"
	"sync"

	netcontext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FakeRolodexClient struct {
	GetOwnersStub        func(ctx netcontext.Context, in *rolodexpb.GetOwnersRequest, opts ...grpc.CallOption) (*rolodexpb.GetOwnersResponse, error)
	getOwnersMutex       sync.RWMutex
	getOwnersArgsForCall []struct {
		ctx  netcontext.Context
		in   *rolodexpb.GetOwnersRequest
		opts []grpc.CallOption
	}
	getOwnersReturns struct {
		result1 *rolodexpb.GetOwnersResponse
		result2 error
	}
	getOwnersReturnsOnCall map[int]struct {
		result1 *rolodexpb.GetOwnersResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRolodexClient) GetOwners(ctx netcontext.Context, in *rolodexpb.GetOwnersRequest, opts ...grpc.CallOption) (*rolodexpb.GetOwnersResponse, error) {
	fake.getOwnersMutex.Lock()
	ret, specificReturn := fake.getOwnersReturnsOnCall[len(fake.getOwnersArgsForCall)]
	fake.getOwnersArgsForCall = append(fake.getOwnersArgsForCall, struct {
		ctx  netcontext.Context
		in   *rolodexpb.GetOwnersRequest
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("GetOwners", []interface{}{ctx, in, opts})
	fake.getOwnersMutex.Unlock()
	if fake.GetOwnersStub != nil {
		return fake.GetOwnersStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getOwnersReturns.result1, fake.getOwnersReturns.result2
}

func (fake *FakeRolodexClient) GetOwnersCallCount() int {
	fake.getOwnersMutex.RLock()
	defer fake.getOwnersMutex.RUnlock()
	return len(fake.getOwnersArgsForCall)
}

func (fake *FakeRolodexClient) GetOwnersArgsForCall(i int) (netcontext.Context, *rolodexpb.GetOwnersRequest, []grpc.CallOption) {
	fake.getOwnersMutex.RLock()
	defer fake.getOwnersMutex.RUnlock()
	return fake.getOwnersArgsForCall[i].ctx, fake.getOwnersArgsForCall[i].in, fake.getOwnersArgsForCall[i].opts
}

func (fake *FakeRolodexClient) GetOwnersReturns(result1 *rolodexpb.GetOwnersResponse, result2 error) {
	fake.GetOwnersStub = nil
	fake.getOwnersReturns = struct {
		result1 *rolodexpb.GetOwnersResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRolodexClient) GetOwnersReturnsOnCall(i int, result1 *rolodexpb.GetOwnersResponse, result2 error) {
	fake.GetOwnersStub = nil
	if fake.getOwnersReturnsOnCall == nil {
		fake.getOwnersReturnsOnCall = make(map[int]struct {
			result1 *rolodexpb.GetOwnersResponse
			result2 error
		})
	}
	fake.getOwnersReturnsOnCall[i] = struct {
		result1 *rolodexpb.GetOwnersResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRolodexClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOwnersMutex.RLock()
	defer fake.getOwnersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRolodexClient) recordInvocation(key string, args []interface{}) {
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

var _ notifications.RolodexClient = new(FakeRolodexClient)
