package distro

import (
	"fmt"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Distro() {
	distro, err := exec.Command("lsb_release", "-si").Output()
	if err != nil {
		fmt.Println("Error! Distro")
	}
	fmt.Printf(Red+"Distribution: "+Reset+"%s", distro)
}
