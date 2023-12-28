package main

import (
	"fmt"
	"gofetch/cmd/cpu"
	"gofetch/cmd/display"
	"gofetch/cmd/distro"
	"gofetch/cmd/gpu"
	"gofetch/cmd/hostname"
	"gofetch/cmd/kernel"
	"gofetch/cmd/manager"
	"gofetch/cmd/memory"
	"gofetch/cmd/server"
	"gofetch/cmd/shell"
	"gofetch/cmd/terminal"
	"gofetch/cmd/uptime"
	"gofetch/cmd/user"
	"os"
	"os/exec"
	"strings"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func main() {
	check, err := exec.Command("uname", "-o").Output()
	if err != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check)) == "GNU/Linux" {
		fmt.Printf(Red+"%s", string(check)+Reset+"\n")

		user.User()
		distro.Distro()
		hostname.Hostname()
		uptime.Uptime()
		display.Display()
		cpu.CPU()
		gpu.GPU()
		server.Server()
		shell.Shell()
		memory.Memory()
		kernel.Kernel()
		terminal.Terminal()
		// weather.Weather()
		manager.Manager()
	} else if strings.TrimSpace(string(check)) == "Android" {
		fmt.Printf(Red+"%s", string(check)+Reset+"\n")

		user.User()
		distro.Distro()
		cpu.CPU()
		server.Server()
		shell.Shell()
		uptime.Uptime()
	}
}
