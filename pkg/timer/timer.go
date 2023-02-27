package timer

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
)

type Timer struct {
	Start        time.Time
	InWorkMode   bool
	WorkDuration int
	RestDuration int
}

func (t Timer) getElapsedTimeInSeconds() int {
	return int(time.Since(t.Start).Seconds())
}

func (t *Timer) switchMode() {
	t.Start = time.Now()
	t.InWorkMode = !t.InWorkMode
}

func (t Timer) alert() {
	message := "Take a break"
	if !t.InWorkMode {
		message = "Back to work"
	}
	fmt.Println(message)
	// os := runtime.GOOS
	// if os == "darwin" && !isQuietTime() {
	// 	go exec.Command("say", message).Output()
	// }
	beeep.Notify(message, message, "assets/information.png")
}

func (t Timer) shouldSwitchMode(elapsed int) bool {
	return elapsed == t.getDuration()
}

func (t Timer) getDuration() int {
	duration := t.WorkDuration
	if !t.InWorkMode {
		duration = t.RestDuration
	}
	return duration
}
func (t Timer) getMode() string {
	mode := "Work"
	if !t.InWorkMode {
		mode = "Rest"
	}
	return mode
}

func (t Timer) printTimeRemaining(elapsed int, clock *widget.Label) {
	timeRemaining := t.getDuration() - elapsed
	minutes := timeRemaining / 60
	seconds := timeRemaining - minutes*60
	updatedTime := fmt.Sprintf("\r%v: %02d:%02d", t.getMode(), minutes, seconds)
	clock.SetText(updatedTime)
}

func (t Timer) UpdateTime(clock *widget.Label) {
	prevElapsed := 0
	for {
		elapsed := t.getElapsedTimeInSeconds()
		if elapsed != prevElapsed {
			t.printTimeRemaining(elapsed, clock)
			prevElapsed = elapsed
			if t.shouldSwitchMode(elapsed) {
				t.alert()
				t.switchMode()
			}
		}
	}
}
