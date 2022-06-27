package alog

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWorker(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer

	ch := make(chan string, 10)
	w := NewWorker(&b, ch)

	go w.StartBackground()
	defer w.Stop()

	ch <- "test"
	time.Sleep(time.Second)
	s := b.String()

	logText := getLogText(t, s)
	require.Equal(t, "test", logText)
}
