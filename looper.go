package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func looper() {
	for {
	}
}

func env(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		ival, _ := strconv.Atoi(val)
		return ival
	}
	return def
}

func main() {

	var wg sync.WaitGroup

	numCPUS := flag.Int("cpus", env("LOOPER_CPUS", 1), "Number of CPUs to occupy")

	flag.Parse()

	runtime.GOMAXPROCS(2048)

	fmt.Printf("Yum, CPU cycles...")

	for i := 0; i < *numCPUS; i++ {
		wg.Add(1)
		go looper()
	}

	wg.Wait()
}
