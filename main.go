package main

import (
	"gofetch/cmd/display"
	"gofetch/cmd/hostname"
	"gofetch/cmd/kernel"
	"gofetch/cmd/manager"
	"gofetch/cmd/terminal"
	"gofetch/cmd/uptime"
	"gofetch/cmd/user"
	"gofetch/cmd/weather"
)

func main() {
	user.User()
	hostname.Hostname()
	uptime.Uptime()
	kernel.Kernel()
	terminal.Terminal()
	weather.Weather()
	manager.Manager()
	display.Display()
}
