// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc"
)

type FakeProperties struct {
	SetStub        func(handle string, key string, value string)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		handle string
		key    string
		value  string
	}
	GetStub        func(handle string, key string) (string, bool)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		handle string
		key    string
	}
	getReturns struct {
		result1 string
		result2 bool
	}
}

func (fake *FakeProperties) Set(handle string, key string, value string) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		handle string
		key    string
		value  string
	}{handle, key, value})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(handle, key, value)
	}
}

func (fake *FakeProperties) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeProperties) SetArgsForCall(i int) (string, string, string) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].handle, fake.setArgsForCall[i].key, fake.setArgsForCall[i].value
}

func (fake *FakeProperties) Get(handle string, key string) (string, bool) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		handle string
		key    string
	}{handle, key})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(handle, key)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeProperties) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeProperties) GetArgsForCall(i int) (string, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].handle, fake.getArgsForCall[i].key
}

func (fake *FakeProperties) GetReturns(result1 string, result2 bool) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

var _ rundmc.Properties = new(FakeProperties)