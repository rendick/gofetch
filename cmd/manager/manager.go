package manager

import (
	"fmt"
	"os"
	"os/exec"
)

func Manager() {
	cmd := exec.Command("sh", "-c", "echo $DESKTOP_SESSION")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	} else {
		fmt.Printf("WM: %s", output)
	}
}
