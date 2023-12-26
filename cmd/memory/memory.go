package memory

import (
	"fmt"
	"os"
	"syscall"
)

func Memory() {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		os.Exit(0)
	}

	var mb int = int(in.Totalram) / 1000

	fmt.Printf("Memory: %d MB\n", mb)
}
