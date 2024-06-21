package service
import "github.com/shirou/gopsutil/v3/process"

type processManager struct{}

func newProcessManager()processManager{
	processManager := &processManager{}
	return processManager
}

func(p *processManager) killProcesses(name string){
	processList, err := ps.Processes()
	for _, p := range processes {
        n, err := p.Name()
        if err != nil {
            return err
        }
        if n == name {
            return p.Kill()
        }
    }
}

func(p *processManager) listProcesses(processName string)processListString string{
	processList, err := ps.Processes()
    if err != nil {
        log.Println("ps.Processes() Failed, are you using windows?")
        return
    }

	processListString := ""

    // map ages
    for x := range processList {
        var process ps.Process
        process = processList[x]
        name, _ := process.Name()
        processListString = processListString + name + ", \n"
    }

}
