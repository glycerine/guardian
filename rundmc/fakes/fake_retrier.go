// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc"
)

type FakeRetrier struct {
	RunStub        func(fn func() error) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		fn func() error
	}
	runReturns struct {
		result1 error
	}
}

func (fake *FakeRetrier) Run(fn func() error) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		fn func() error
	}{fn})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(fn)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeRetrier) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeRetrier) RunArgsForCall(i int) func() error {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].fn
}

func (fake *FakeRetrier) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

var _ rundmc.Retrier = new(FakeRetrier)
