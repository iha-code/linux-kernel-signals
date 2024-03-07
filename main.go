package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"slices"
	"strconv"
	"strings"
	"syscall"
)

func main() {

	kernel_info, err := ioutil.ReadFile("/proc/version")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Print kernel information
	str1 := ""
	str2 := "-"
	for i := 0; i < len(kernel_info); i++ {
		str1 = str1 + str2
	}

	fmt.Println(str1)
	fmt.Println("Linux Kernel Info:")
	fmt.Println(strings.TrimSpace(string(kernel_info)))
	fmt.Println(str1)
	// Get all the signals
	signals := []os.Signal{
		syscall.SIGABRT,
		syscall.SIGALRM,
		syscall.SIGBUS,
		syscall.SIGCHLD,
		syscall.SIGCONT,
		syscall.SIGFPE,
		syscall.SIGHUP,
		syscall.SIGILL,
		syscall.SIGINT,
		syscall.SIGIO,
		syscall.SIGIOT,
		syscall.SIGKILL,
		syscall.SIGPIPE,
		syscall.SIGPOLL,
		syscall.SIGPROF,
		syscall.SIGPWR,
		syscall.SIGQUIT,
		syscall.SIGSEGV,
		syscall.SIGSTKFLT,
		syscall.SIGSTOP,
		syscall.SIGSYS,
		syscall.SIGTERM,
		syscall.SIGTRAP,
		syscall.SIGTSTP,
		syscall.SIGTTIN,
		syscall.SIGTTOU,
		syscall.SIGUNUSED,
		syscall.SIGURG,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGVTALRM,
		syscall.SIGWINCH,
		syscall.SIGXCPU,
		syscall.SIGXFSZ,
	}
	signalMap := make(map[int]os.Signal)

	for _, s := range signals {
		res := fmt.Sprintf("%d", s)
		sigcode, err := strconv.Atoi(res)

		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		signalMap[sigcode] = s
	}
	keys := make([]int, 0, len(signalMap))

	// Populate the slice with keys from the map
	for key := range signalMap {
		keys = append(keys, key)
	}

	// Sort the keys slice
	slices.Sort(keys)

	// Iterate over the sorted keys and access corresponding values from the map
	fmt.Println("List of Linux Kernel syscall signals sorted by code:")
	fmt.Println("----------------------------------------------------")
	for _, key := range keys {
		fmt.Printf("Signal:(%d) %s\n", key, signalMap[key])
	}

	// Wait for interrupt signal (Ctrl+C)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Exiting...")
}
