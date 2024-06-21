package service

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessManager interface {
	KillProcesses(string) error
	ListProcesses() string
}

type ProcessManagerImpl struct{}

func NewProcessManager() *ProcessManagerImpl {
	processManager := &ProcessManagerImpl{}
	return processManager
}

func (p *ProcessManagerImpl) KillProcesses(name string) error {
	processList, err := process.Processes()
	if err != nil {
		fmt.Println("process.Processes() Failed, are you using windows?")
		return nil
	}

	for _, p := range processList {
		n, err := p.Name()
		if err != nil {
			return err
		}

		if n == name {
			fmt.Println("Killing App")
			return p.Kill()
		}
	}

	return fmt.Errorf("process not found")
}

func (p *ProcessManagerImpl) ListProcesses() string {
	processList, err := process.Processes()
	if err != nil {
		fmt.Println("process.Processes() Failed, are you using windows?")
		return ""
	}

	processListString := ""

	// map ages
	for x := range processList {
		var process process.Process
		process = *processList[x]
		name, _ := process.Name()
		if name != "" {
			processListString = processListString + name + ", \n"
		}
	}

	return processListString

}
