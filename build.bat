@echo off
go build -o GoAT.exe -ldflags "-s -H windowsgui" GoAT.go
echo GoAT.exe has been built. Press any key to close the window...
pause >nul
