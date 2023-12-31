package memory

import (
	"fmt"
	"os"
	"os/exec" // Add this import for the exec package
	"strconv"
	"strings"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

var MemoryInfo string

func Memory() {
	// Commented out code, remove if not needed
	// in := &syscall.Sysinfo_t{}
	// err := syscall.Sysinfo(in)
	// if err != nil {
	// 	os.Exit(0)
	// }

	// total_ram := int64(in.Totalram) / 1024
	// free_ram := int64(in.Freeram) / 1024
	// MemoryInfo = fmt.Sprintf(Red+"Memory: "+Reset+"%d MB / %d MB\n", free_ram/1024, total_ram/1024)

	// Run the command to get memory information
	memory_linux, err := exec.Command("sh", "-c", "cat /proc/meminfo | grep 'MemAvailable:' | awk '{print $2}'").Output()
	if err != nil {
		os.Exit(0)
	}

	// Convert the byte slice to string and remove leading/trailing whitespaces
	memoryInfo := strings.TrimSpace(string(memory_linux))
	i, err_convert := strconv.Atoi(memoryInfo)
	if err_convert != nil {
		panic(err)
	}

	// Print the memory information
	MemoryInfo = fmt.Sprintf(Red+"Memory: "+Reset+"%d", i/1024)
}
