package alog

import (
	"fmt"
	"io"
	"time"
)

// Worker is a log writer worker.
type Worker struct {
	out io.Writer

	logChan chan string
	quit    chan bool
}

// NewWorker returns new log writer worker.
func NewWorker(out io.Writer, logChan chan string) *Worker {
	return &Worker{
		out:     out,
		logChan: logChan,
		quit:    make(chan bool),
	}
}

// StartBackground is a starter for log writer worker.
func (w *Worker) StartBackground() {
	for {
		select {
		case log := <-w.logChan:
			fmt.Fprintf(w.out, "[%d] %s", time.Now().Unix(), log)
		case <-w.quit:
			return
		}
	}
}

// Stop worker.
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
