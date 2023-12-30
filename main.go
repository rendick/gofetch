package main

import (
	"fmt"
	"gofetch/cmd/cpu"
	"gofetch/cmd/display"
	"gofetch/cmd/distro"
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
	check, err := exec.Command("uname", "-o").Output()
	if err != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check)) == "GNU/Linux" {
		user.User()
		hostname.Hostname()
		distro.Distro()
		kernel.Kernel()
		shell.Shell()
		display.Display()
		uptime.Uptime()
		manager.Manager()
		terminal.Terminal()
		cpu.CPU()
		server.Server()
		memory.Memory()
		fmt.Printf(Logo,
			strings.Replace(user.UserInfo, "\n", " ", -1),
			strings.Replace(hostname.HostnameInfo, "\n", " ", -1),
			strings.Replace(distro.DistroInfo, "\n", " ", -1),
			strings.Replace(kernel.KernelInfo, "\n", " ", -1),
			strings.Replace(shell.ShellType, "\n", " ", -1),
			strings.Replace(display.DisplayInfo, "\n", " ", -1),
			strings.Replace(uptime.UptimeInfo, "\n", " ", -1),
			strings.Replace(manager.ManagerInfo, "\n", " ", -1),
			strings.Replace(terminal.TerminalInfo, "\n", " ", -1),
			strings.Replace(cpu.CpuInfo, "\n", " ", -1),
			strings.Replace(server.ServerInfo, "\n", " ", -1),
			strings.Replace(memory.MemoryInfo, "\n", " ", -1))

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
