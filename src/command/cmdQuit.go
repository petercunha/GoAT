package command

import (
	"log"
)

func Quit() {
	log.Fatal("Commander has sent !quit command! Shutting down GoAT...")
}
