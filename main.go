package main

import (
	"gofetch/cmd/hostname"
	"gofetch/cmd/kernel"
	"gofetch/cmd/manager"
	"gofetch/cmd/uptime"
	"gofetch/cmd/user"
	"gofetch/cmd/weather"
)

func main() {
	user.User()
	hostname.Hostname()
	uptime.Uptime()
	kernel.Kernel()
	weather.Weather()
	manager.Manager()
}
