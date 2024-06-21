package main

import (
	service "GameBreakerConsole/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
