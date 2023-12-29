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

var (
	Logo = ` 
  __ _  ___   
 / _  |/ _ \   %s 
| (_| | (_) |  %s 
 \__, |\___/   %s 
 |___/    _    %s 
 / _| ___| |_  %s
| |_ / _ \ __| %s
|  _|  __/ |_  %s
|_|  \___|\__| %s
  ___| |__     %s 
 / __|  _ \    %s
| (__| | | |   %s 
 \___|_| |_|   %s 
`
)

func main() {
	user.User()
	distro.Distro()
	shell.Shell()
	uptime.Uptime()
	terminal.Terminal()
	fmt.Printf(Logo,
		strings.Replace(user.Username, "\n", " ", -1),
		strings.Replace(distro.Distribution, "\n", " ", -1),
		strings.Replace(shell.ShellType, "\n", " ", -1),
		strings.Replace(uptime.Time, "\n", " ", -1),
		strings.Replace(terminal.App, "\n", " ", -1))

	check, err := exec.Command("uname", "-o").Output()
	if err != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check)) == "GNU/Linux" {
		fmt.Printf(Red+"%s", string(check)+Reset+"\n")

		user.User()
		distro.Distro()
		kernel.Kernel()
		hostname.Hostname()
		uptime.Uptime()
		display.Display()
		terminal.Terminal()
		cpu.CPU()
		gpu.GPU()
		server.Server()
		shell.Shell()
		memory.Memory()
		// weather.Weather()
		manager.Manager()
	} else if strings.TrimSpace(string(check)) == "Android" {
		fmt.Printf(Red+"%s", string(check)+Reset+"\n")

		user.User()
		distro.Distro()
		kernel.Kernel()
		hostname.Hostname()
		uptime.Uptime()
		terminal.Terminal()
		cpu.CPU()
		server.Server()
		shell.Shell()
	} else {
		os.Exit(0)
	}
}
