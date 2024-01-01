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
	check_distro, err_distro := exec.Command("uname", "-o").Output()
	if err_distro != nil {
		os.Exit(0)
	} else if strings.TrimSpace(string(check_distro)) == "GNU/Linux" {

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

		MemoryInfo = fmt.Sprintf(Red+"Memory: "+Reset+"%d MB / %d MB", (total-avail)/1024, total/1024)
	} else if strings.TrimSpace(string(check_distro)) == "Android" {
		avail_memory_android, err_avail_android := exec.Command("sh", "-c", "cat /proc/meminfo | grep 'MemAvailable:' | awk '{print $2}'").Output()
		if err_avail_android != nil {
			os.Exit(0)
		}
		total_memory_android, err_total_android := exec.Command("sh", "-c", "cat /proc/meminfo | grep 'MemTotal:' | awk '{print $2}'").Output()
		if err_total_android != nil {
			os.Exit(0)
		}

		avail_convert_android := strings.TrimSpace(string(avail_memory_android))
		avail_android, err_convert_avail_android := strconv.Atoi(avail_convert_android)
		if err_convert_avail_android != nil {
			panic(err_avail_android)
		}

		total_convert_android := strings.TrimSpace(string(total_memory_android))
		total_android, err_convert_total := strconv.Atoi(total_convert_android)
		if err_convert_total != nil {
			panic(err_total_android)
		}
		MemoryInfo = fmt.Sprintf(Red+"Memory: "+Reset+"%d MB / %d MB", (total_android-avail_android)/1024, total_android/1024)

	}
}
