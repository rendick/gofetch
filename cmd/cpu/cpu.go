package cpu

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func CPU() {
	cpu, err := exec.Command("sh", "-c", "lscpu | grep 'Model name' | cut -f 2 -d ':' | awk '{$1=$1}1'").Output()
	if err != nil {
		os.Exit(0)
	} else {
		fmt.Printf(Red+"CPU:"+Reset+"%s", cpu)
	}
}
