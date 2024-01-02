package hostname

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

var HostnameInfo string

func Hostname() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		hostname_linux, err_linux := exec.Command("cat", "/etc/hostname").Output()
		if err_linux != nil {
			os.Exit(0)
		} else {
			HostnameInfo = fmt.Sprintf(Red+Bold+"Hostname: "+Reset+"%s", hostname_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		hostname_android, err_android := exec.Command("hostname").Output()
		if err_android != nil {
			os.Exit(0)
		} else {
			HostnameInfo = fmt.Sprintf(Red+Bold+"Hostname: "+Reset+"%s", hostname_android)
		}
	}
}
