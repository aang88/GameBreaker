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

func (p *CountdownTimerImpl) countDown(setTime int) {
	fmt.Printf("Starting countdown timer for %d seconds...\n", setTime)

	timer := time.NewTimer(time.Duration(setTime) * time.Second)

	// Block until the timer channel receives a value
	<-timer.C

	fmt.Println("Timer expired! Continue with the rest of the program.")
}

func (p *CountdownTimerImpl) coolDown(setTime int, done chan<- bool) {
	fmt.Printf("Starting countdown timer for %d seconds...\n", setTime)

	timer := time.NewTimer(time.Duration(setTime) * time.Second)

	go func() {
		select {
		case <-timer.C:
			fmt.Println("Timer expired! Continue with the rest of the program.")
			done <- true
		}
	}()

}
