package manager

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

var ManagerInfo string

func Manager() {
	manager, err := exec.Command("sh", "-c", "echo $XDG_SESSION_DESKTOP").Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	} else {
		ManagerInfo = fmt.Sprintf(Red+Bold+"WM: "+Reset+"%s", manager)
	}
}
