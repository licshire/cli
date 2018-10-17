// Code generated by counterfeiter. DO NOT EDIT.
package v3actionfakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
)

type FakeUAAClient struct {
	GetSSHPasscodeStub        func(string, string) (string, error)
	getSSHPasscodeMutex       sync.RWMutex
	getSSHPasscodeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSSHPasscodeReturns struct {
		result1 string
		result2 error
	}
	getSSHPasscodeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAAClient) GetSSHPasscode(arg1 string, arg2 string) (string, error) {
	fake.getSSHPasscodeMutex.Lock()
	ret, specificReturn := fake.getSSHPasscodeReturnsOnCall[len(fake.getSSHPasscodeArgsForCall)]
	fake.getSSHPasscodeArgsForCall = append(fake.getSSHPasscodeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSSHPasscode", []interface{}{arg1, arg2})
	fake.getSSHPasscodeMutex.Unlock()
	if fake.GetSSHPasscodeStub != nil {
		return fake.GetSSHPasscodeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSSHPasscodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) GetSSHPasscodeCallCount() int {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	return len(fake.getSSHPasscodeArgsForCall)
}

func (fake *FakeUAAClient) GetSSHPasscodeCalls(stub func(string, string) (string, error)) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = stub
}

func (fake *FakeUAAClient) GetSSHPasscodeArgsForCall(i int) (string, string) {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	argsForCall := fake.getSSHPasscodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) GetSSHPasscodeReturns(result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	fake.getSSHPasscodeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetSSHPasscodeReturnsOnCall(i int, result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	if fake.getSSHPasscodeReturnsOnCall == nil {
		fake.getSSHPasscodeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getSSHPasscodeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAAClient) recordInvocation(key string, args []interface{}) {
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

var _ v3action.UAAClient = new(FakeUAAClient)
