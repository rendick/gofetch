package uptime

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

var UptimeInfo string

func Uptime() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		uptime_linux, err_linux := exec.Command("uptime", "-p").Output()
		if err_linux != nil {
			os.Exit(0)
		} else {
			UptimeInfo = fmt.Sprintf(Red+"Uptime: "+Reset+"%s", uptime_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		uptime_android, err_android := exec.Command("uptime", "-p").Output()
		if err_android != nil {
			os.Exit(0)
		} else {
			UptimeInfo = fmt.Sprintf(Red+"Uptime: "+Reset+"%s", uptime_android)
		}
	}
}
