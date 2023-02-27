package main

import (
	"fmt"
	"localhost/bptimer/pkg/timer"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const defaultWorkDurationInMins = 25
const defaultRestDurationInMins = 5

func main() {
	a := app.New()
	win := a.NewWindow("Pomodoro")
	clock := widget.NewLabel("")
	workDurationInMins := defaultWorkDurationInMins
	restDurationInMins := defaultRestDurationInMins
	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of arguments, using default values of 25minutes working and 5minutes resting")
	}
	if len(os.Args) == 3 {
		wdur, err1 := strconv.Atoi(os.Args[1])
		rdur, err2 := strconv.Atoi(os.Args[2])
		workDurationInMins = wdur
		restDurationInMins = rdur
		if err1 != nil || err2 != nil {
			fmt.Println("Problem setting interval values, using default values of 25minutes working and 5minutes resting")
		}
	}

	t := timer.Timer {
		Start:        time.Now(),
		InWorkMode:   true,
		WorkDuration: workDurationInMins * 60,
		RestDuration: restDurationInMins * 60,
	}

	go t.UpdateTime(clock)
	win.SetContent(clock)
	win.ShowAndRun()
}
