package hostname

import (
	"fmt"
	"log"
	"os/exec"
)

func Hostname() {
	out, err := exec.Command("cat", "/etc/hostname").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Hostname: %s", out)
}
