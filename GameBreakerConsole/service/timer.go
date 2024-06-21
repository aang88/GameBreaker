package service

import (
	"fmt"
	"time"
)

type CountDownTimer interface {
	countDown(int)
	coolDown(setTime int, done chan<- bool)
}

type CountdownTimerImpl struct {
}

func NewCountdownTimer() *CountdownTimerImpl {
	countdownTimer := &CountdownTimerImpl{}
	return countdownTimer
}

func (p *CountdownTimerImpl) countDown(setTime string) {
	fmt.Printf("Starting countdown timer for %s ...\n", setTime)
	timeDuration, err := time.ParseDuration(setTime)
	if err != nil {
		fmt.Println("Enter a right time! ex.5h30m40s")
		return
	}
	timer := time.NewTimer(timeDuration)

	// Block until the timer channel receives a value
	<-timer.C

	fmt.Println("Timer expired! Continue with the rest of the program.")
}

func (p *CountdownTimerImpl) coolDown(setTime string, done chan<- bool) {
	fmt.Printf("Starting countdown timer for %s ...\n", setTime)
	timeDuration, err := time.ParseDuration(setTime)
	if err != nil {
		fmt.Println("Enter a right time! ex.5h30m40s")
		return
	}
	timer := time.NewTimer(time.Duration(timeDuration) * time.Second)

	go func() {
		select {
		case <-timer.C:
			fmt.Println("Timer expired! Continue with the rest of the program.")
			done <- true
		}
	}()

}
