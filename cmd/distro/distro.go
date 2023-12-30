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

var DistroInfo string

func Distro() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		distro_linux, err_linux := exec.Command("lsb_release", "-sd").Output()
		architecture_linux, err_linux_architecture := exec.Command("uname", "-m").Output()
		if err_linux != nil {
			os.Exit(0)
		} else if err_linux_architecture != nil {
			os.Exit(0)
		} else {
			distroLinux := strings.ReplaceAll(string(distro_linux), "\"", "")
			architectureLinux := strings.ReplaceAll(string(architecture_linux), "\"", "")

			DistroInfo = fmt.Sprintf(Red+"Distribution: "+Reset+"%s%s", strings.Replace(distroLinux, "\n", " ", -1), architectureLinux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		distro_android, err_android := exec.Command("uname", "-o").Output()
		architecture_android, err_android_architecture := exec.Command("uname", "-m").Output()
		if err_android != nil {
			os.Exit(0)
		} else if err_android_architecture != nil {
			os.Exit(0)
		} else {
			distroAndroid := strings.ReplaceAll(string(distro_android), "\"", "")
			architectureAndroid := strings.ReplaceAll(string(architecture_android), "\"", "")

			DistroInfo = fmt.Sprintf(Red+"Distribution: "+Reset+"%s%s", strings.Replace(distroAndroid, "\n", " ", -1), architectureAndroid)
		}
	}
}
