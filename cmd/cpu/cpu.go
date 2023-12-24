package cpu

import (
	"fmt"
	"os/exec"
)

func CPU() {
	temp, err := exec.Command("sh", "-c", "lscpu | grep 'Model name' | cut -f 2 -d ':' | awk '{$1=$1}1'").Output()
	if err != nil {
		fmt.Println("error! temp")
	}

	fmt.Printf("Temperature: %s", temp)
}
