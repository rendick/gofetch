package main

import (
	"gofetch/cmd/hostname"
	"gofetch/cmd/kernel"
	"gofetch/cmd/manager"
	"gofetch/cmd/uptime"
)

func main() {
	hostname.Hostname()
	uptime.Uptime()
	kernel.Kernel()
	manager.Manager()
}
