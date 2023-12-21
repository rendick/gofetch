package kernel

import (
	"fmt"
	"os/exec"
)

func Kernel() {
	kernel, err := exec.Command("uname", "--kernel-release").Output()
	if err != nil {
		fmt.Println("Error! kernel")
	}

	fmt.Printf("Kernel: %s", kernel)
}
