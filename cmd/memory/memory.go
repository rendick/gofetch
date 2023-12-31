package memory

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

var MemoryInfo string

func Memory() {
	avail_memory_linux, err_avail := exec.Command("sh", "-c", "cat /proc/meminfo | grep 'MemAvailable:' | awk '{print $2}'").Output()
	if err_avail != nil {
		os.Exit(0)
	}
	total_memory_linux, err_total := exec.Command("sh", "-c", "cat /proc/meminfo | grep 'MemTotal:' | awk '{print $2}'").Output()
	if err_total != nil {
		os.Exit(0)
	}

	avail_convert := strings.TrimSpace(string(avail_memory_linux))
	avail, err_convert_avail := strconv.Atoi(avail_convert)
	if err_convert_avail != nil {
		panic(err_avail)
	}

	total_convert := strings.TrimSpace(string(total_memory_linux))
	total, err_convert_total := strconv.Atoi(total_convert)
	if err_convert_total != nil {
		panic(err_total)
	}
	fmt.Printf("%d and %T", total, total)
	fmt.Printf("%d and %T %d", avail, avail, (total-avail)/1024)

	MemoryInfo = fmt.Sprintf(Red+"Memory: "+Reset+"%d MB / %d MB", (total-avail)/1024, total/1024)
}
