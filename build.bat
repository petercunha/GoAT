@echo off
go build -o GoAT.exe -ldflags "-s -H windowsgui" GoAT.go
echo Press any key to launch GoAT, or close the window.
pause >nul
GoAT.exe