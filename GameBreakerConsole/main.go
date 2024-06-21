package main

import (
	service "GameBreakerConsole/service"
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

//go:embed resources/ascii.txt
var b []byte

func main() {

	if !amAdmin() {
		runMeElevated()
	}
	time.Sleep(5 * time.Second)

	processManager := service.NewProcessManager()
	reader := bufio.NewReader(os.Stdin)

	var processToKill string
	var countDownTime string
	var coolDownTime string
	var gameKiller *service.GameKillerImpl

	for {
		//Inital State of da app
		if processToKill == "" || countDownTime == "" || coolDownTime == "" {
			banner()

			fmt.Println("Now Printing Process List:")
			spinnerWee(28)

			fmt.Println(processManager.ListProcesses())

			fmt.Println("Please enter the process you want to kill:")
			processToKill, _ = reader.ReadString('\n')
			processToKill = strings.TrimSpace(processToKill)

			fmt.Println("Please enter the time you want to play the game for, format like (0h0m0s):")
			countDownTime, _ = reader.ReadString('\n')
			countDownTime = strings.TrimSpace(countDownTime)

			fmt.Println("Please enter how long you want the playing cooldown for, format like (0h0m0s):")
			coolDownTime, _ = reader.ReadString('\n')
			coolDownTime = strings.TrimSpace(coolDownTime)

			gameKiller = service.NewGameKiller(countDownTime, coolDownTime, processToKill)
			gameKiller.KillGames()
		}
		gameKiller.KillGames()

	}
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		fmt.Println("Open as admin to continue")
		return false
	}
	fmt.Println("Opening as admin")
	spinnerWee(28)
	return true
}

func banner() {
	fmt.Println(string(b))
	spinnerWee(50)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func spinnerWee(duration int) {
	spinner := []rune{'|', '/', '-', '\\'}
	for i := 0; i < duration; i++ {
		fmt.Printf("\r%c", spinner[i%len(spinner)])
		time.Sleep(100 * time.Millisecond)
	}
}
