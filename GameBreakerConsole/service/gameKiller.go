package service

type GameKiller interface {
	NewGameKiller(int, int, string) *GameKillerImpl
	KillGames()
}

type GameKillerImpl struct {
	onBreak        bool
	countdownTime  int
	cooldownTime   int
	processName    string
	countDownTimer CountdownTimerImpl
	processManager ProcessManagerImpl
}

func NewGameKiller(countdownTime int, cooldownTime int, processName string) *GameKillerImpl {

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
	g.processManager.KillProcesses(g.processName)
	g.countDownTimer.countDown(g.cooldownTime)
	g.onBreak = false
}
