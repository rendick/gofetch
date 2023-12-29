package cpu

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

func CPU() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		cpu_linux, err_linux := exec.Command("sh", "-c", "lscpu | grep 'Model name' | cut -f 2 -d ':' | awk '{$1=$1}1'").Output()
		if err_linux != nil {
			os.Exit(0)
		} else {
			fmt.Printf(Red+"CPU: "+Reset+"%s", cpu_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		cpu_android, err_android := exec.Command("sh", "-c", "lscpu | grep 'Model name' | cut -f 2 -d ':' | awk '{$1=$1}1'").Output()

		cpuAndroid := strings.ReplaceAll(string(cpu_android), "\"", "")
		if err_android != nil {
			os.Exit(0)
		} else {
			fmt.Printf(Red+"CPU: "+Reset+"%s", strings.Replace(cpuAndroid, "\n", " ", -1))
		}
	}
}
