GoAT
===
**Golang Advanced Trojan** v0.5 Beta

This is a trojan made in Go, using Twitter as a the C&C server. 

===
### Commands
```!echo <message>``` - Logs message to slave console

```!quit``` - Closes GoAT

```!clear``` - Does nothing. Use this command if you don't want slaves to execute anything upon connecting.


### Compilation
Compile with  ```go build -o GoAT.exe -ldflags "-H windowsgui" "C:\GoAT.go"```	to have no console show.


### Under development
While I keep teaching myself Go, I will be updating this project. If at anytime someone would like to add to the project they can. Just post in the 'issues' section and I will add your code and credits to the main project.

===
## TODO:
* Check for >1 running instance
* Commands
  * DDoS
  * Send messagebox
  * Uninstall
  * Shutdown/Restart


