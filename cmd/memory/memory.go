package memory

import (
	"fmt"
	"os"
	"syscall"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

func Memory() {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		os.Exit(0)
	}

	total_ram := int64(in.Totalram) / 1000
	free_ram := int64(in.Freeram) / 1000
	fmt.Printf(Red+"Memory: "+Reset+"%d MB / %d MB\n", free_ram / 1000, total_ram / 1000)
}
