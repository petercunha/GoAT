package command

import (
	"strings"
	"fmt"
)

func Parse(cmd string) {
	if strings.Contains(cmd, "!echo") {
		Echo(strings.Split(cmd, "!echo ")[1])
	} else if strings.Contains(cmd, "!quit") {
		Quit()
	} else {
		fmt.Println("Unknown command!")
	}
}