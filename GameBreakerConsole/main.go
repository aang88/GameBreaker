package main

import (
	service "GameBreakerConsole/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

func main() {

	if !amAdmin() {
		runMeElevated()
	}
	time.Sleep(5 * time.Second)

	processManager := service.NewProcessManager()
	reader := bufio.NewReader(os.Stdin)

	var processToKill string
	var countDownTime int
	var coolDownTime int
	var gameKiller *service.GameKillerImpl

	for {
		if processToKill == "" || countDownTime == 0 || coolDownTime == 0 {
			fmt.Println("Process List:")
			fmt.Println(processManager.ListProcesses())

			fmt.Println("Please enter the process you want to kill:")
			processToKill, _ = reader.ReadString('\n')
			processToKill = strings.TrimSpace(processToKill)

			fmt.Println("Please enter the time you want to play the game for:")
			countDownStr, _ := reader.ReadString('\n')
			countDownStr = strings.TrimSpace(countDownStr)
			countDownTime, _ = strconv.Atoi(countDownStr)

			fmt.Println("Please enter how long you want the playing cooldown for:")
			coolDownStr, _ := reader.ReadString('\n')
			coolDownStr = strings.TrimSpace(coolDownStr)
			coolDownTime, _ = strconv.Atoi(coolDownStr)

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
	fmt.Println("Opening as admin...")
	return true
}
