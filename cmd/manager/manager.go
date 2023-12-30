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

var ManagerInfo string

func Manager() {
	manager, err := exec.Command("sh", "-c", "echo $XDG_SESSION_DESKTOP").Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	} else {
		ManagerInfo = fmt.Sprintf(Red+"WM: "+Reset+"%s", manager)
	}
}
