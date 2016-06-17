// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc"
	"github.com/pivotal-golang/lager"
)

type FakeWaitWatcher struct {
	OnExitStub        func(log lager.Logger, process runrunc.Waiter, onExit runrunc.Runner)
	onExitMutex       sync.RWMutex
	onExitArgsForCall []struct {
		log     lager.Logger
		process runrunc.Waiter
		onExit  runrunc.Runner
	}
}

func (fake *FakeWaitWatcher) OnExit(log lager.Logger, process runrunc.Waiter, onExit runrunc.Runner) {
	fake.onExitMutex.Lock()
	fake.onExitArgsForCall = append(fake.onExitArgsForCall, struct {
		log     lager.Logger
		process runrunc.Waiter
		onExit  runrunc.Runner
	}{log, process, onExit})
	fake.onExitMutex.Unlock()
	if fake.OnExitStub != nil {
		fake.OnExitStub(log, process, onExit)
	}
}

func (fake *FakeWaitWatcher) OnExitCallCount() int {
	fake.onExitMutex.RLock()
	defer fake.onExitMutex.RUnlock()
	return len(fake.onExitArgsForCall)
}

func (fake *FakeWaitWatcher) OnExitArgsForCall(i int) (lager.Logger, runrunc.Waiter, runrunc.Runner) {
	fake.onExitMutex.RLock()
	defer fake.onExitMutex.RUnlock()
	return fake.onExitArgsForCall[i].log, fake.onExitArgsForCall[i].process, fake.onExitArgsForCall[i].onExit
}

var _ runrunc.WaitWatcher = new(FakeWaitWatcher)