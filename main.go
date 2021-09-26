package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sync"

	"github.com/xykkong/goroutine-example/stream"
	"github.com/xykkong/goroutine-example/worker"
)

var NUM_THREAD = 10
var timeout uint64

func init() {
	flag.Uint64Var(&timeout, "timeout", 60, "Maximium time in seconds for a worker to run until finding the target string.")
	timeout *= 1000
	flag.Parse()
}

func print(done chan worker.Message) {
	for r := range done {
		b, _ := json.Marshal(r)
		fmt.Println(string(b))
	}
}

func main() {
	fmt.Printf("Process started with timeout = %d\n", timeout)
	done := make(chan worker.Message)
	stream_ := &stream.Stream{}
	var wg sync.WaitGroup

	for i := 0; i < NUM_THREAD; i++ {
		wg.Add(1)
		worker := &worker.Worker{}
		go worker.Run(stream_, timeout, &wg, done)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	print(done)
}
