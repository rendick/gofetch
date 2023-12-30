package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

var TerminalInfo string

func Terminal() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		terminal_linux, err_linux := exec.Command("sh", "-c", "echo $TERM").Output()
		if err_linux != nil {
			fmt.Println("Error! Terminal")
		} else {
			TerminalInfo = fmt.Sprintf(Red+"Terminal: "+Reset+"%s", terminal_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		terminal_android, err_android := exec.Command("sh", "-c", "echo $TERM").Output()
		if err_android != nil {
			fmt.Println("Error! Terminal")
		} else {
			TerminalInfo = fmt.Sprintf(Red+"Terminal: "+Reset+"%s", terminal_android)
		}
	}

}
