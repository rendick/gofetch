package display

import (
	"fmt"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Display() {
	display, err := exec.Command("bash", "-c", "xrandr | grep '*' | awk '{ print $1 }'").Output()
	if err != nil {
		fmt.Println("Error! display")
	}
	fmt.Printf(Red+"Display: "+Reset+"%s", display)

}
