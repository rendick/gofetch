package terminal

import (
	"fmt"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Terminal() {
	terminal, err := exec.Command("sh", "-c", "echo $TERM").Output()
	if err != nil {
		fmt.Println("Error! Terminal")
	}
	fmt.Printf(Red+"Terminal: "+Reset+"%s", terminal)
}
