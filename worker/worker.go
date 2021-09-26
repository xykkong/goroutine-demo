package worker

import (
	"strings"
	"sync"
	"time"

	"github.com/xykkong/goroutine-example/stream"
)

type Status int

const (
	Success Status = iota
	Timeout
)

func (s Status) String() string {
	return []string{"SUCCESS", "TIMEOUT"}[s]
}

type Message struct {
	ElapsedTime uint64
	BytesRead   uint64
	Status      string
}

type Worker struct{}

func (w *Worker) Run(stream *stream.Stream, timeoutMs uint64, wg *sync.WaitGroup, done chan Message) {
	defer wg.Done()

	startTime := time.Now()
	found := false

	for {
		str, bytesRead := stream.GetData(10)
		if strings.Contains(str, "Demo") {
			found = true
		}

		elapsed := uint64(time.Since(startTime).Milliseconds())

		if found {
			done <- Message{
				ElapsedTime: elapsed,
				BytesRead:   bytesRead,
				Status:      Success.String(),
			}
			break
		}

		if elapsed >= timeoutMs {
			done <- Message{
				Status: Timeout.String(),
			}
			break
		}
	}
}
