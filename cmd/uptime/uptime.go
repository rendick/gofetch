package uptime

import (
	"fmt"
	"os/exec"
)

func Uptime() {
	uptime, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		fmt.Println("Error!")
	}

	fmt.Printf("Uptime: %s", uptime)
}
