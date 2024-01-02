package display

import (
	"fmt"
	"os/exec"
)

var (
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

var DisplayInfo string

func Display() {
	display, err := exec.Command("bash", "-c", "xrandr | grep '*' | awk '{ print $1 }'").Output()
	if err != nil {
		fmt.Println("Error! display")
	}
	DisplayInfo = fmt.Sprintf(Red+Bold+"Display: "+Reset+"%s", display)

}
