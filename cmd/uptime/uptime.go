package uptime

import (
	"fmt"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Uptime() {
	uptime, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		fmt.Println("Error!")
	}

	fmt.Printf(Red+"Uptime: "+Reset+"%s", uptime)
}
