// This file was generated by counterfeiter
package knfsdriverfakes

import (
	"sync"

	"code.cloudfoundry.org/goshims/execshim"
	"code.cloudfoundry.org/lager"
	"github.com/lds-cf/knfsdriver/kerberizer"
)

type FakeKerberizer struct {
	LoginStub        func(lager.Logger, string, string) error
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}
	loginReturns struct {
		result1 error
	}
	LoginWithExecStub        func(lager.Logger, execshim.Exec, string, string) error
	loginWithExecMutex       sync.RWMutex
	loginWithExecArgsForCall []struct {
		arg1 lager.Logger
		arg2 execshim.Exec
		arg3 string
		arg4 string
	}
	loginWithExecReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKerberizer) Login(arg1 lager.Logger, arg2 string, arg3 string) error {
	fake.loginMutex.Lock()
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Login", []interface{}{arg1, arg2, arg3})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(arg1, arg2, arg3)
	} else {
		return fake.loginReturns.result1
	}
}

func (fake *FakeKerberizer) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeKerberizer) LoginArgsForCall(i int) (lager.Logger, string, string) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return fake.loginArgsForCall[i].arg1, fake.loginArgsForCall[i].arg2, fake.loginArgsForCall[i].arg3
}

func (fake *FakeKerberizer) LoginReturns(result1 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKerberizer) LoginWithExec(arg1 lager.Logger, arg2 execshim.Exec, arg3 string, arg4 string) error {
	fake.loginWithExecMutex.Lock()
	fake.loginWithExecArgsForCall = append(fake.loginWithExecArgsForCall, struct {
		arg1 lager.Logger
		arg2 execshim.Exec
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("LoginWithExec", []interface{}{arg1, arg2, arg3, arg4})
	fake.loginWithExecMutex.Unlock()
	if fake.LoginWithExecStub != nil {
		return fake.LoginWithExecStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.loginWithExecReturns.result1
	}
}

func (fake *FakeKerberizer) LoginWithExecCallCount() int {
	fake.loginWithExecMutex.RLock()
	defer fake.loginWithExecMutex.RUnlock()
	return len(fake.loginWithExecArgsForCall)
}

func (fake *FakeKerberizer) LoginWithExecArgsForCall(i int) (lager.Logger, execshim.Exec, string, string) {
	fake.loginWithExecMutex.RLock()
	defer fake.loginWithExecMutex.RUnlock()
	return fake.loginWithExecArgsForCall[i].arg1, fake.loginWithExecArgsForCall[i].arg2, fake.loginWithExecArgsForCall[i].arg3, fake.loginWithExecArgsForCall[i].arg4
}

func (fake *FakeKerberizer) LoginWithExecReturns(result1 error) {
	fake.LoginWithExecStub = nil
	fake.loginWithExecReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKerberizer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.loginWithExecMutex.RLock()
	defer fake.loginWithExecMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeKerberizer) recordInvocation(key string, args []interface{}) {
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

var _ kerberizer.Kerberizer = new(FakeKerberizer)
