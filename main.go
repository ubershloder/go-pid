package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	pid := readinput()
	proc, err := os.FindProcess(int(pid))
	if err != nil {
		panic("no such process")
	} else {
		send_sig := proc.Signal(syscall.Signal(0))
		fmt.Printf("Pid %d returned: %v\n", pid, send_sig)

	}
}

func readinput() int {
	input := os.Args[1:2]
	if len(input) < 1 {
		panic("no args")
	}
	var input_str string = strings.Join(input, " ")
	pid, err := strconv.Atoi(input_str)
	if err != nil {
		panic(err)
	}
	return pid
}
