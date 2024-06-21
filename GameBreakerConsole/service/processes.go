package service

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessManager interface {
	KillProcesses(string) error
	ListProcesses() string
}

type ProcessManagerImpl struct {
	processList []*process.Process
}

func NewProcessManager() *ProcessManagerImpl {
	processList, err := process.Processes()
	if err != nil {
		fmt.Println("process.Processes() Failed, are you using windows?")
	}
	processManager := &ProcessManagerImpl{
		processList: processList,
	}
	return processManager
}

func (p *ProcessManagerImpl) KillProcesses(name string) error {

	for _, process := range p.processList {
		n, _ := process.Name()

		if n == name {
			fmt.Println("Killing App")
			err := process.Terminate()
			fmt.Println(err)
			return err
		}
	}

	return fmt.Errorf("process not found")
}

func (p *ProcessManagerImpl) ListProcesses() string {

	processListString := ""

	// map ages
	for x := range p.processList {
		var process process.Process
		process = *p.processList[x]
		name, _ := process.Name()
		if name != "" {
			processListString = processListString + name + ", \n"
		}
	}

	return processListString

}

func (p *ProcessManagerImpl) refreshProcessList() {
	p.processList, _ = process.Processes()
}
