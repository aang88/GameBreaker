package service

import (
	"fmt"
	"time"
)

type GameKiller interface {
	NewGameKiller(int, int, string) *GameKillerImpl
	KillGames()
}

type GameKillerImpl struct {
	onBreak        bool
	countdownTime  string
	cooldownTime   string
	processName    string
	countDownTimer CountdownTimerImpl
	processManager ProcessManagerImpl
}

func NewGameKiller(countdownTime string, cooldownTime string, processName string) *GameKillerImpl {

	gameKiller := &GameKillerImpl{
		onBreak:        false,
		countdownTime:  countdownTime,
		cooldownTime:   cooldownTime,
		processName:    processName,
		countDownTimer: *NewCountdownTimer(),
		processManager: *NewProcessManager(),
	}
	return gameKiller
}

func (g *GameKillerImpl) KillGames() {
	if !g.onBreak {
		g.countDownTimer.countDown(g.countdownTime)
		g.processManager.KillProcesses(g.processName)
		g.onBreak = true
		return
	}
	done := make(chan bool)
	g.processManager.refreshProcessList()
	g.countDownTimer.coolDown(g.cooldownTime, done)
	for {
		select {
		case <-done:
			fmt.Println("Countdown timer has completed.")
			g.onBreak = false
			return
		default:
			g.processManager.refreshProcessList()
			g.processManager.KillProcesses(g.processName)
			time.Sleep(1 * time.Second)
		}
	}
}
