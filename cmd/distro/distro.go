package distro

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

func Distro() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	}

	if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		distro_linux, err := exec.Command("lsb_release", "-si").Output()
		if err != nil {
			os.Exit(0)
		}
		fmt.Printf(Red+"Distribution: "+Reset+"%s", distro_linux)
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		distro_android, err_android := exec.Command("uname", "-o").Output()
		if err_android != nil {
			os.Exit(0)
		}
		fmt.Printf(Red+"Distribution: "+Reset+"%s", distro_android)
	}
}
