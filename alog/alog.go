package alog

import (
	"fmt"
	"io"
)

// Logger is a async logger impl.
type Logger struct {
	out io.Writer

	workers []*Worker
	ch      chan string
}

// NewLogger returns new async Logger.
func NewLogger(out io.Writer, workersCount, logChanSize int) *Logger {
	l := Logger{
		out:     out,
		workers: make([]*Worker, 0, workersCount),
		ch:      make(chan string, logChanSize),
	}

	l.runWorkers(workersCount)

	return &l
}

func (l *Logger) runWorkers(count int) {
	for i := 0; i < count; i++ {
		l.workers = append(l.workers, NewWorker(l.out, l.ch))
		go l.workers[i].StartBackground()
	}
}

// Stop async Logger.
func (l *Logger) Stop() {
	for _, w := range l.workers {
		w.Stop()
	}
}

// Info write param.
func (l *Logger) Info(v ...any) {
	l.ch <- fmt.Sprint(v...)
}

// Infof write params by format.
func (l *Logger) Infof(format string, v ...any) {
	l.ch <- fmt.Sprintf(format, v...)
}

// Infoln write params separated by space with new line.
func (l *Logger) Infoln(v ...any) {
	l.ch <- fmt.Sprintln(v...)
}
