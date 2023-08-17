// Code generated by counterfeiter. DO NOT EDIT.
package channelfakes

import (
	"sync"

	"github.com/phgermanov/notification-service/internal/channel"
)

type FakeSender struct {
	GetNameStub        func() string
	getNameMutex       sync.RWMutex
	getNameArgsForCall []struct {
	}
	getNameReturns struct {
		result1 string
	}
	getNameReturnsOnCall map[int]struct {
		result1 string
	}
	SendStub        func(string) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 string
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSender) GetName() string {
	fake.getNameMutex.Lock()
	ret, specificReturn := fake.getNameReturnsOnCall[len(fake.getNameArgsForCall)]
	fake.getNameArgsForCall = append(fake.getNameArgsForCall, struct {
	}{})
	stub := fake.GetNameStub
	fakeReturns := fake.getNameReturns
	fake.recordInvocation("GetName", []interface{}{})
	fake.getNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSender) GetNameCallCount() int {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	return len(fake.getNameArgsForCall)
}

func (fake *FakeSender) GetNameCalls(stub func() string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = stub
}

func (fake *FakeSender) GetNameReturns(result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	fake.getNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSender) GetNameReturnsOnCall(i int, result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	if fake.getNameReturnsOnCall == nil {
		fake.getNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSender) Send(arg1 string) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SendStub
	fakeReturns := fake.sendReturns
	fake.recordInvocation("Send", []interface{}{arg1})
	fake.sendMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSender) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *FakeSender) SendCalls(stub func(string) error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = stub
}

func (fake *FakeSender) SendArgsForCall(i int) string {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	argsForCall := fake.sendArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSender) SendReturns(result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) SendReturnsOnCall(i int, result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSender) recordInvocation(key string, args []interface{}) {
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

var _ channel.Sender = new(FakeSender)