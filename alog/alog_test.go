package alog

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	logger := NewLogger(os.Stdout, 1, 5)
	require.Len(t, logger.workers, 1)
	logger.Stop()
}

func TestLogger_Info(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer
	l := getLogger(&b)
	defer l.Stop()

	l.Info("test")
	time.Sleep(time.Second)
	logText := getLogText(t, b.String())
	require.Equal(t, "test", logText)
}

func TestLogger_Infoln(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer
	l := getLogger(&b)
	defer l.Stop()

	l.Infoln("test")
	time.Sleep(time.Second)
	logText := getLogText(t, b.String())
	require.Equal(t, "test\n", logText)
}

func TestLogger_Infof(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer
	l := getLogger(&b)
	defer l.Stop()

	l.Infof("%s %d", "some number", 10)
	time.Sleep(time.Second)
	logText := getLogText(t, b.String())
	require.Equal(t, "some number 10", logText)
}

func getLogger(out io.Writer) *Logger {
	return NewLogger(out, 1, 5)
}

func getLogText(t *testing.T, log string) string {
	splitted := strings.Split(log, "] ")
	require.Len(t, splitted, 2)

	return splitted[1]
}
