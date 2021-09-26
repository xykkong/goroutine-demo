package stream

import (
	"math/rand"
	"time"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Stream struct {
	bytesRead uint64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (s *Stream) GetData(n uint64) (string, uint64) {
	b := make([]byte, n)
	for i := range b {
		b[i] = letter[rand.Int63()%int64(len(letter))]
	}
	s.bytesRead += n
	return string(b), s.bytesRead
}
