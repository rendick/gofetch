package manager

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Manager() {
	cmd := exec.Command("sh", "-c", "echo $DESKTOP_SESSION")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	} else {
		fmt.Printf(Red+"WM: "+Reset+"%s", output)
	}
}
