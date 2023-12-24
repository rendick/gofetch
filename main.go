package main

import (
	"fmt"
	"gofetch/cmd/cpu"
	"gofetch/cmd/display"
	"gofetch/cmd/distro"
	"gofetch/cmd/hostname"
	"gofetch/cmd/kernel"
	"gofetch/cmd/manager"
	"gofetch/cmd/terminal"
	"gofetch/cmd/uptime"
	"gofetch/cmd/user"
	"gofetch/cmd/weather"
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
		os.Exit(1)
	} else if strings.TrimSpace(string(check)) == "GNU/Linux" {
		fmt.Printf(Red + "GNU/Linux" + Reset + "\n\n")

		user.User()
		distro.Distro()
		hostname.Hostname()
		uptime.Uptime()
		display.Display()
		cpu.CPU()
		kernel.Kernel()
		terminal.Terminal()
		weather.Weather()
		manager.Manager()
	} else if strings.TrimSpace(string(check)) == "Android" {
		fmt.Println("Android")
	}
}
