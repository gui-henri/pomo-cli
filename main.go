package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"

	"gopkg.in/toast.v1"
)

func main() {
	fmt.Println("POMO CLI")

	minutes := flag.Int("min", 5, "The number of minutes counted in this pomodoro \n")
	flag.Parse()
	remaining := *minutes

	for i := 0; i < *minutes; i++ {
		exec.Command("clear")
		remaining = remaining - 1
		fmt.Println("Time remaining: ", remaining)
		time.Sleep(1 * time.Minute)
	}

	notification := toast.Notification{
		AppID: "Pomo CLI",
		Title: "Your pomodoro has ended",
		Message: "Start a new pomodoro or close",
	}

	err := notification.Push()

	if err != nil {
		log.Fatal(err)
	}
}
