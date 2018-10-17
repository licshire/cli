// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeScaleActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetApplicationByNameAndSpaceStub        func(string, string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationSummaryByNameAndSpaceStub        func(string, string, bool) (v3action.ApplicationSummary, v3action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}
	PollStartStub        func(string, chan<- v3action.Warnings) error
	pollStartMutex       sync.RWMutex
	pollStartArgsForCall []struct {
		arg1 string
		arg2 chan<- v3action.Warnings
	}
	pollStartReturns struct {
		result1 error
	}
	pollStartReturnsOnCall map[int]struct {
		result1 error
	}
	ScaleProcessByApplicationStub        func(string, v3action.Process) (v3action.Warnings, error)
	scaleProcessByApplicationMutex       sync.RWMutex
	scaleProcessByApplicationArgsForCall []struct {
		arg1 string
		arg2 v3action.Process
	}
	scaleProcessByApplicationReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	scaleProcessByApplicationReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	StartApplicationStub        func(string) (v3action.Application, v3action.Warnings, error)
	startApplicationMutex       sync.RWMutex
	startApplicationArgsForCall []struct {
		arg1 string
	}
	startApplicationReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	startApplicationReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	StopApplicationStub        func(string) (v3action.Warnings, error)
	stopApplicationMutex       sync.RWMutex
	stopApplicationArgsForCall []struct {
		arg1 string
	}
	stopApplicationReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	stopApplicationReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScaleActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeScaleActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeScaleActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeScaleActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeScaleActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeScaleActor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeScaleActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeScaleActor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v3action.Application, v3action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeScaleActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeScaleActor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScaleActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScaleActor) GetApplicationSummaryByNameAndSpace(arg1 string, arg2 string, arg3 bool) (v3action.ApplicationSummary, v3action.Warnings, error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)]
	fake.getApplicationSummaryByNameAndSpaceArgsForCall = append(fake.getApplicationSummaryByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetApplicationSummaryByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationSummaryByNameAndSpaceStub != nil {
		return fake.GetApplicationSummaryByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationSummaryByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeScaleActor) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeScaleActor) GetApplicationSummaryByNameAndSpaceCalls(stub func(string, string, bool) (v3action.ApplicationSummary, v3action.Warnings, error)) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = stub
}

func (fake *FakeScaleActor) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string, bool) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationSummaryByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeScaleActor) GetApplicationSummaryByNameAndSpaceReturns(result1 v3action.ApplicationSummary, result2 v3action.Warnings, result3 error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScaleActor) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v3action.ApplicationSummary, result2 v3action.Warnings, result3 error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.ApplicationSummary
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.ApplicationSummary
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScaleActor) PollStart(arg1 string, arg2 chan<- v3action.Warnings) error {
	fake.pollStartMutex.Lock()
	ret, specificReturn := fake.pollStartReturnsOnCall[len(fake.pollStartArgsForCall)]
	fake.pollStartArgsForCall = append(fake.pollStartArgsForCall, struct {
		arg1 string
		arg2 chan<- v3action.Warnings
	}{arg1, arg2})
	fake.recordInvocation("PollStart", []interface{}{arg1, arg2})
	fake.pollStartMutex.Unlock()
	if fake.PollStartStub != nil {
		return fake.PollStartStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pollStartReturns
	return fakeReturns.result1
}

func (fake *FakeScaleActor) PollStartCallCount() int {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return len(fake.pollStartArgsForCall)
}

func (fake *FakeScaleActor) PollStartCalls(stub func(string, chan<- v3action.Warnings) error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = stub
}

func (fake *FakeScaleActor) PollStartArgsForCall(i int) (string, chan<- v3action.Warnings) {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	argsForCall := fake.pollStartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeScaleActor) PollStartReturns(result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	fake.pollStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScaleActor) PollStartReturnsOnCall(i int, result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	if fake.pollStartReturnsOnCall == nil {
		fake.pollStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pollStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScaleActor) ScaleProcessByApplication(arg1 string, arg2 v3action.Process) (v3action.Warnings, error) {
	fake.scaleProcessByApplicationMutex.Lock()
	ret, specificReturn := fake.scaleProcessByApplicationReturnsOnCall[len(fake.scaleProcessByApplicationArgsForCall)]
	fake.scaleProcessByApplicationArgsForCall = append(fake.scaleProcessByApplicationArgsForCall, struct {
		arg1 string
		arg2 v3action.Process
	}{arg1, arg2})
	fake.recordInvocation("ScaleProcessByApplication", []interface{}{arg1, arg2})
	fake.scaleProcessByApplicationMutex.Unlock()
	if fake.ScaleProcessByApplicationStub != nil {
		return fake.ScaleProcessByApplicationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.scaleProcessByApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeScaleActor) ScaleProcessByApplicationCallCount() int {
	fake.scaleProcessByApplicationMutex.RLock()
	defer fake.scaleProcessByApplicationMutex.RUnlock()
	return len(fake.scaleProcessByApplicationArgsForCall)
}

func (fake *FakeScaleActor) ScaleProcessByApplicationCalls(stub func(string, v3action.Process) (v3action.Warnings, error)) {
	fake.scaleProcessByApplicationMutex.Lock()
	defer fake.scaleProcessByApplicationMutex.Unlock()
	fake.ScaleProcessByApplicationStub = stub
}

func (fake *FakeScaleActor) ScaleProcessByApplicationArgsForCall(i int) (string, v3action.Process) {
	fake.scaleProcessByApplicationMutex.RLock()
	defer fake.scaleProcessByApplicationMutex.RUnlock()
	argsForCall := fake.scaleProcessByApplicationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeScaleActor) ScaleProcessByApplicationReturns(result1 v3action.Warnings, result2 error) {
	fake.scaleProcessByApplicationMutex.Lock()
	defer fake.scaleProcessByApplicationMutex.Unlock()
	fake.ScaleProcessByApplicationStub = nil
	fake.scaleProcessByApplicationReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeScaleActor) ScaleProcessByApplicationReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.scaleProcessByApplicationMutex.Lock()
	defer fake.scaleProcessByApplicationMutex.Unlock()
	fake.ScaleProcessByApplicationStub = nil
	if fake.scaleProcessByApplicationReturnsOnCall == nil {
		fake.scaleProcessByApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.scaleProcessByApplicationReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeScaleActor) StartApplication(arg1 string) (v3action.Application, v3action.Warnings, error) {
	fake.startApplicationMutex.Lock()
	ret, specificReturn := fake.startApplicationReturnsOnCall[len(fake.startApplicationArgsForCall)]
	fake.startApplicationArgsForCall = append(fake.startApplicationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StartApplication", []interface{}{arg1})
	fake.startApplicationMutex.Unlock()
	if fake.StartApplicationStub != nil {
		return fake.StartApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.startApplicationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeScaleActor) StartApplicationCallCount() int {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return len(fake.startApplicationArgsForCall)
}

func (fake *FakeScaleActor) StartApplicationCalls(stub func(string) (v3action.Application, v3action.Warnings, error)) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = stub
}

func (fake *FakeScaleActor) StartApplicationArgsForCall(i int) string {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	argsForCall := fake.startApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeScaleActor) StartApplicationReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = nil
	fake.startApplicationReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScaleActor) StartApplicationReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = nil
	if fake.startApplicationReturnsOnCall == nil {
		fake.startApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.startApplicationReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScaleActor) StopApplication(arg1 string) (v3action.Warnings, error) {
	fake.stopApplicationMutex.Lock()
	ret, specificReturn := fake.stopApplicationReturnsOnCall[len(fake.stopApplicationArgsForCall)]
	fake.stopApplicationArgsForCall = append(fake.stopApplicationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StopApplication", []interface{}{arg1})
	fake.stopApplicationMutex.Unlock()
	if fake.StopApplicationStub != nil {
		return fake.StopApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.stopApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeScaleActor) StopApplicationCallCount() int {
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	return len(fake.stopApplicationArgsForCall)
}

func (fake *FakeScaleActor) StopApplicationCalls(stub func(string) (v3action.Warnings, error)) {
	fake.stopApplicationMutex.Lock()
	defer fake.stopApplicationMutex.Unlock()
	fake.StopApplicationStub = stub
}

func (fake *FakeScaleActor) StopApplicationArgsForCall(i int) string {
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	argsForCall := fake.stopApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeScaleActor) StopApplicationReturns(result1 v3action.Warnings, result2 error) {
	fake.stopApplicationMutex.Lock()
	defer fake.stopApplicationMutex.Unlock()
	fake.StopApplicationStub = nil
	fake.stopApplicationReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeScaleActor) StopApplicationReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.stopApplicationMutex.Lock()
	defer fake.stopApplicationMutex.Unlock()
	fake.StopApplicationStub = nil
	if fake.stopApplicationReturnsOnCall == nil {
		fake.stopApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.stopApplicationReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeScaleActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	fake.scaleProcessByApplicationMutex.RLock()
	defer fake.scaleProcessByApplicationMutex.RUnlock()
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	fake.stopApplicationMutex.RLock()
	defer fake.stopApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScaleActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.ScaleActor = new(FakeScaleActor)
