// This file was generated by counterfeiter
package fakes

import (
	"os/exec"
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc"
)

type FakeRuncBinary struct {
	ExecCommandStub        func(id, processJSONPath, pidFilePath string) *exec.Cmd
	execCommandMutex       sync.RWMutex
	execCommandArgsForCall []struct {
		id              string
		processJSONPath string
		pidFilePath     string
	}
	execCommandReturns struct {
		result1 *exec.Cmd
	}
	EventsCommandStub        func(id string) *exec.Cmd
	eventsCommandMutex       sync.RWMutex
	eventsCommandArgsForCall []struct {
		id string
	}
	eventsCommandReturns struct {
		result1 *exec.Cmd
	}
	StateCommandStub        func(id, logFile string) *exec.Cmd
	stateCommandMutex       sync.RWMutex
	stateCommandArgsForCall []struct {
		id      string
		logFile string
	}
	stateCommandReturns struct {
		result1 *exec.Cmd
	}
	StatsCommandStub        func(id, logFile string) *exec.Cmd
	statsCommandMutex       sync.RWMutex
	statsCommandArgsForCall []struct {
		id      string
		logFile string
	}
	statsCommandReturns struct {
		result1 *exec.Cmd
	}
	KillCommandStub        func(id, signal, logFile string) *exec.Cmd
	killCommandMutex       sync.RWMutex
	killCommandArgsForCall []struct {
		id      string
		signal  string
		logFile string
	}
	killCommandReturns struct {
		result1 *exec.Cmd
	}
}

func (fake *FakeRuncBinary) ExecCommand(id string, processJSONPath string, pidFilePath string) *exec.Cmd {
	fake.execCommandMutex.Lock()
	fake.execCommandArgsForCall = append(fake.execCommandArgsForCall, struct {
		id              string
		processJSONPath string
		pidFilePath     string
	}{id, processJSONPath, pidFilePath})
	fake.execCommandMutex.Unlock()
	if fake.ExecCommandStub != nil {
		return fake.ExecCommandStub(id, processJSONPath, pidFilePath)
	} else {
		return fake.execCommandReturns.result1
	}
}

func (fake *FakeRuncBinary) ExecCommandCallCount() int {
	fake.execCommandMutex.RLock()
	defer fake.execCommandMutex.RUnlock()
	return len(fake.execCommandArgsForCall)
}

func (fake *FakeRuncBinary) ExecCommandArgsForCall(i int) (string, string, string) {
	fake.execCommandMutex.RLock()
	defer fake.execCommandMutex.RUnlock()
	return fake.execCommandArgsForCall[i].id, fake.execCommandArgsForCall[i].processJSONPath, fake.execCommandArgsForCall[i].pidFilePath
}

func (fake *FakeRuncBinary) ExecCommandReturns(result1 *exec.Cmd) {
	fake.ExecCommandStub = nil
	fake.execCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeRuncBinary) EventsCommand(id string) *exec.Cmd {
	fake.eventsCommandMutex.Lock()
	fake.eventsCommandArgsForCall = append(fake.eventsCommandArgsForCall, struct {
		id string
	}{id})
	fake.eventsCommandMutex.Unlock()
	if fake.EventsCommandStub != nil {
		return fake.EventsCommandStub(id)
	} else {
		return fake.eventsCommandReturns.result1
	}
}

func (fake *FakeRuncBinary) EventsCommandCallCount() int {
	fake.eventsCommandMutex.RLock()
	defer fake.eventsCommandMutex.RUnlock()
	return len(fake.eventsCommandArgsForCall)
}

func (fake *FakeRuncBinary) EventsCommandArgsForCall(i int) string {
	fake.eventsCommandMutex.RLock()
	defer fake.eventsCommandMutex.RUnlock()
	return fake.eventsCommandArgsForCall[i].id
}

func (fake *FakeRuncBinary) EventsCommandReturns(result1 *exec.Cmd) {
	fake.EventsCommandStub = nil
	fake.eventsCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeRuncBinary) StateCommand(id string, logFile string) *exec.Cmd {
	fake.stateCommandMutex.Lock()
	fake.stateCommandArgsForCall = append(fake.stateCommandArgsForCall, struct {
		id      string
		logFile string
	}{id, logFile})
	fake.stateCommandMutex.Unlock()
	if fake.StateCommandStub != nil {
		return fake.StateCommandStub(id, logFile)
	} else {
		return fake.stateCommandReturns.result1
	}
}

func (fake *FakeRuncBinary) StateCommandCallCount() int {
	fake.stateCommandMutex.RLock()
	defer fake.stateCommandMutex.RUnlock()
	return len(fake.stateCommandArgsForCall)
}

func (fake *FakeRuncBinary) StateCommandArgsForCall(i int) (string, string) {
	fake.stateCommandMutex.RLock()
	defer fake.stateCommandMutex.RUnlock()
	return fake.stateCommandArgsForCall[i].id, fake.stateCommandArgsForCall[i].logFile
}

func (fake *FakeRuncBinary) StateCommandReturns(result1 *exec.Cmd) {
	fake.StateCommandStub = nil
	fake.stateCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeRuncBinary) StatsCommand(id string, logFile string) *exec.Cmd {
	fake.statsCommandMutex.Lock()
	fake.statsCommandArgsForCall = append(fake.statsCommandArgsForCall, struct {
		id      string
		logFile string
	}{id, logFile})
	fake.statsCommandMutex.Unlock()
	if fake.StatsCommandStub != nil {
		return fake.StatsCommandStub(id, logFile)
	} else {
		return fake.statsCommandReturns.result1
	}
}

func (fake *FakeRuncBinary) StatsCommandCallCount() int {
	fake.statsCommandMutex.RLock()
	defer fake.statsCommandMutex.RUnlock()
	return len(fake.statsCommandArgsForCall)
}

func (fake *FakeRuncBinary) StatsCommandArgsForCall(i int) (string, string) {
	fake.statsCommandMutex.RLock()
	defer fake.statsCommandMutex.RUnlock()
	return fake.statsCommandArgsForCall[i].id, fake.statsCommandArgsForCall[i].logFile
}

func (fake *FakeRuncBinary) StatsCommandReturns(result1 *exec.Cmd) {
	fake.StatsCommandStub = nil
	fake.statsCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeRuncBinary) KillCommand(id string, signal string, logFile string) *exec.Cmd {
	fake.killCommandMutex.Lock()
	fake.killCommandArgsForCall = append(fake.killCommandArgsForCall, struct {
		id      string
		signal  string
		logFile string
	}{id, signal, logFile})
	fake.killCommandMutex.Unlock()
	if fake.KillCommandStub != nil {
		return fake.KillCommandStub(id, signal, logFile)
	} else {
		return fake.killCommandReturns.result1
	}
}

func (fake *FakeRuncBinary) KillCommandCallCount() int {
	fake.killCommandMutex.RLock()
	defer fake.killCommandMutex.RUnlock()
	return len(fake.killCommandArgsForCall)
}

func (fake *FakeRuncBinary) KillCommandArgsForCall(i int) (string, string, string) {
	fake.killCommandMutex.RLock()
	defer fake.killCommandMutex.RUnlock()
	return fake.killCommandArgsForCall[i].id, fake.killCommandArgsForCall[i].signal, fake.killCommandArgsForCall[i].logFile
}

func (fake *FakeRuncBinary) KillCommandReturns(result1 *exec.Cmd) {
	fake.KillCommandStub = nil
	fake.killCommandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

var _ runrunc.RuncBinary = new(FakeRuncBinary)