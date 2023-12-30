package display

import (
	"fmt"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

var DisplayInfo string

func Display() {
	display, err := exec.Command("bash", "-c", "xrandr | grep '*' | awk '{ print $1 }'").Output()
	if err != nil {
		fmt.Println("Error! display")
	}
	DisplayInfo = fmt.Sprintf(Red+"Display: "+Reset+"%s", display)

}
