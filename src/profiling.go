package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	cpuFile, err := os.Create("data/cpu.out")
	if err != nil {
		fmt.Println("Error creating CPU profile: ", err)
		return
	}

	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		fmt.Println("Error starting CPU profile: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		fmt.Println("I value --> ", i)
	}

	// To view profiling code status
	// type: go tool pprof data/cpu.out
}
