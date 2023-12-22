package kernel

import (
	"fmt"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Kernel() {
	kernel, err := exec.Command("uname", "--kernel-release").Output()
	if err != nil {
		fmt.Println("Error! kernel")
	}

	fmt.Printf(Red+"Kernel: "+Reset+"%s", kernel)
}
