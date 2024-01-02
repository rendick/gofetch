package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

var ShellInfo string

func Shell() {
	check_distro, err := exec.Command("uname", "-o").Output()
	if err != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		shell_linux, err_linux := exec.Command("sh", "-c", "echo $SHELL").Output()
		if err_linux != nil {
			os.Exit(0)
		} else {
			ShellInfo = fmt.Sprintf(Red+Bold+"Shell: "+Reset+"%s", shell_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		shell_android, err_android := exec.Command("sh", "-c", "echo $SHELL").Output()
		if err_android != nil {
			os.Exit(0)
		} else {
			ShellInfo = fmt.Sprintf(Red+Bold+"Shell: "+Reset+"%s", shell_android)
		}

	}
}
