package server

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

var ServerInfo string

func Server() {
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	}

	if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {
		server_linux, err_linux := exec.Command("sh", "-c", "echo $XDG_SESSION_TYPE").Output()
		if err_linux != nil {
			os.Exit(0)
		} else {
			ServerInfo = fmt.Sprintf(Red+"Server: "+Reset+"%s", server_linux)
		}
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		server_android, err_android := exec.Command("sh", "-c", "echo binder").Output()
		if err_android != nil {
			os.Exit(0)
		} else {
			ServerInfo = fmt.Sprintf(Red+"Server: "+Reset+"%s", server_android)
		}
	}
}
