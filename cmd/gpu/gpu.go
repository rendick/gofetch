package gpu

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

func GPU() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		gpu_linux, err_linux := exec.Command("sh", "-c", "lspci | grep VGA").Output()
		if err_linux != nil {
			os.Exit(0)
		} else {
			fmt.Printf(Red+"GPU: "+Reset+"%s", gpu_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		os.Exit(0)
	}
}
