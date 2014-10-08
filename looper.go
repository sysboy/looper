package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

type LoopMessage struct {
	instance int
	loops    uint64
}

func looper(instance int, ch chan LoopMessage) {
	m := LoopMessage{instance, 0}
	t := time.Now().Unix()
	for {
		if time.Now().Unix() == t {
			m.loops++
		} else {
			ch <- m
			m.loops = 0
			t = time.Now().Unix()
		}
	}
}

func env(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		ival, _ := strconv.Atoi(val)
		return ival
	}
	return def
}

func printTitle(threads int) {
	fmt.Printf("%5s", "Time")
	for i := 0; i < threads; i++ {
		fmt.Printf("     T%02d  ", i)
	}
	fmt.Printf("%10s\n", "Sum")
}

func printData(ms []LoopMessage, elapsed int) {
	var sum uint64 = 0
	fmt.Printf("%05d", elapsed)
	for i := 0; i < len(ms); i++ {
		fmt.Printf("%10d", ms[i].loops)
		sum = sum + ms[i].loops
	}
	fmt.Printf("%10d\n", sum)
}

func main() {

	var messages_received int = 0
	var elapsed int = 0

	numThreads := flag.Int("threads", env("LOOPER_THREADS", 1), "Number of threads to launch")

	flag.Parse()

	runtime.GOMAXPROCS(2048)

	ch := make(chan LoopMessage)
	ms := make([]LoopMessage, *numThreads)

	for i := 0; i < *numThreads; i++ {
		go looper(i, ch)
	}

	printTitle(*numThreads)

	for {
		v := <-ch
		messages_received++
		ms[v.instance] = v
		if messages_received == *numThreads {
			printData(ms, elapsed)
			messages_received = 0
			elapsed++
		}
	}
}
