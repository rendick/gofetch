package kernel

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

var KernelInfo string

func Kernel() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		kernel_linux, err_linux := exec.Command("uname", "--kernel-release").Output()
		if err_linux != nil {
			os.Exit(0)
		} else {

			KernelInfo = fmt.Sprintf(Red+Bold+"Kernel: "+Reset+"%s", kernel_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		kernel_android, err_android := exec.Command("uname", "--kernel-release").Output()
		if err_android != nil {
			os.Exit(0)
		} else {
			KernelInfo = fmt.Sprintf(Red+Bold+"Kernel: "+Reset+"%s", kernel_android)
		}
	}
}
