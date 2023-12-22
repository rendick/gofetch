package hostname

import (
	"fmt"
	"log"
	"os/exec"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Hostname() {
	out, err := exec.Command("cat", "/etc/hostname").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(Red+"Hostname: "+Reset+"%s", out)
}
