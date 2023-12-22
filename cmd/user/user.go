package user

import (
	"fmt"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func User() {
	user, err := exec.Command("sh", "-c", "echo $USER").Output()
	if err != nil {
		fmt.Println("User: error")
	}
	fmt.Printf(Red+"User: "+Reset+"%s", user)
}
