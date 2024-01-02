package user

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

var UserInfo string

func User() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		user_linux, err_linux := exec.Command("sh", "-c", "echo $USER").Output()
		if err_linux != nil {
			fmt.Println("User: error")
		} else {
			UserInfo = fmt.Sprintf(Red+Bold+"User: "+Reset+"%s", user_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		user_android, err_android := exec.Command("whoami").Output()
		if err_android != nil {
			os.Exit(0)
		} else {
			UserInfo = fmt.Sprintf(Red+Bold+"User: "+Reset+"%s", user_android)
		}
	} else {
		os.Exit(0)
	}
}
