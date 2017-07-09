package streamlimit

import (
	"errors"
	"strconv"
	"time"
)

var noDataError = errors.New("No data available")

// Streamlimiter is a reader and writer that outputs an input stream on a specific speed
type Streamlimiter struct {
	pool      []byte
	released  []byte
	bytelimit int
	byterate  int
	timerate  int
	every     time.Duration
}

// New creates a new Streamlimiter
func New(byterate, timerate int) Streamlimiter {
	every, _ := time.ParseDuration(strconv.Itoa(1000/timerate) + "ms")
	return Streamlimiter{
		bytelimit: byterate / timerate,
		byterate:  byterate,
		timerate:  timerate,
		every:     every,
	}
}

func (s *Streamlimiter) Read(p []byte) (n int, err error) {
	for len(s.released) == 0 {
		time.Sleep(time.Millisecond)
	}

	for n < len(p) {
		if len(s.released) <= 0 {
			break
		}

		p[n] = s.released[0]
		s.released = s.released[1:]
		n++
	}

	return
}

func (s *Streamlimiter) Write(p []byte) (n int, err error) {
	n = len(p)
	s.pool = append(s.pool, p...)

	return
}

// Start starts the output of the stream
func (s *Streamlimiter) Start() {
	go s.releaseLoop()
}

// RemainingTime tells how many seconds it will take for the pool to be empty
func (s *Streamlimiter) RemainingTime() int {
	return len(s.pool) / s.byterate
}

func (s *Streamlimiter) releaseLoop() {
	for {
		if len(s.pool) > s.bytelimit {
			s.released = append(s.released, s.pool[:s.bytelimit]...)
			s.pool = s.pool[s.bytelimit:]
		} else if len(s.pool) > 0 {
			s.released = append(s.released, s.pool...)
			s.pool = []byte{}
		}

		time.Sleep(s.every)
	}
}
