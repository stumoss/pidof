package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const proc_dir = "/proc"

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		os.Exit(1)
	}

	proc_name := os.Args[1]
	if os.Chdir(proc_dir) != nil {
		fmt.Println("/proc unavailable.")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("unable to read /proc directory.")
	}

	for _, file := range files {
		checkProcName(proc_name, file)
	}
}

func checkProcName(proc_name string, file os.FileInfo) {
	// Ignore files, we only care about directories
	if !file.IsDir() {
		return
	}

	// Our directory name should convert to integer
	// if it's a PID
	pid, err := strconv.Atoi(file.Name())
	if err != nil {
		return
	}

	// Open the /proc/xxx/stat file to read the name
	f, err := os.Open(file.Name() + "/stat")
	if err != nil {
		fmt.Println("unable to open", file.Name())
	}
	defer f.Close()

	r := bufio.NewReader(f)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), proc_name) {
			fmt.Println(pid)
			return
		}
	}
}
