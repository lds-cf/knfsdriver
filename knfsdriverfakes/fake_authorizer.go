// This file was generated by counterfeiter
package knfsdriverfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/lds-cf/knfsdriver/authorizer"
)

type FakeAuthorizer struct {
	AuthorizeStub        func(logger lager.Logger, mountpath string, mountmode authorizer.MountMode, principal, keytab string) error
	authorizeMutex       sync.RWMutex
	authorizeArgsForCall []struct {
		logger    lager.Logger
		mountpath string
		mountmode authorizer.MountMode
		principal string
		keytab    string
	}
	authorizeReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthorizer) Authorize(logger lager.Logger, mountpath string, mountmode authorizer.MountMode, principal string, keytab string) error {
	fake.authorizeMutex.Lock()
	fake.authorizeArgsForCall = append(fake.authorizeArgsForCall, struct {
		logger    lager.Logger
		mountpath string
		mountmode authorizer.MountMode
		principal string
		keytab    string
	}{logger, mountpath, mountmode, principal, keytab})
	fake.recordInvocation("Authorize", []interface{}{logger, mountpath, mountmode, principal, keytab})
	fake.authorizeMutex.Unlock()
	if fake.AuthorizeStub != nil {
		return fake.AuthorizeStub(logger, mountpath, mountmode, principal, keytab)
	} else {
		return fake.authorizeReturns.result1
	}
}

func (fake *FakeAuthorizer) AuthorizeCallCount() int {
	fake.authorizeMutex.RLock()
	defer fake.authorizeMutex.RUnlock()
	return len(fake.authorizeArgsForCall)
}

func (fake *FakeAuthorizer) AuthorizeArgsForCall(i int) (lager.Logger, string, authorizer.MountMode, string, string) {
	fake.authorizeMutex.RLock()
	defer fake.authorizeMutex.RUnlock()
	return fake.authorizeArgsForCall[i].logger, fake.authorizeArgsForCall[i].mountpath, fake.authorizeArgsForCall[i].mountmode, fake.authorizeArgsForCall[i].principal, fake.authorizeArgsForCall[i].keytab
}

func (fake *FakeAuthorizer) AuthorizeReturns(result1 error) {
	fake.AuthorizeStub = nil
	fake.authorizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthorizer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authorizeMutex.RLock()
	defer fake.authorizeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAuthorizer) recordInvocation(key string, args []interface{}) {
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

var _ authorizer.Authorizer = new(FakeAuthorizer)
