package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"

	"github.com/EpicStep/async-logger-test-task/alog"
)

type appConfig struct {
	workersCount  int
	logsPerWorker int
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var cfg appConfig

	flag.IntVar(&cfg.workersCount, "workersCount", runtime.NumCPU(), "workers count")
	flag.IntVar(&cfg.logsPerWorker, "logsPerWorker", runtime.NumCPU()*10, "logs count")
	flag.Parse()

	logger := alog.NewLogger(os.Stdout, cfg.workersCount, 1000)
	defer logger.Stop()

	for i := 0; i < cfg.logsPerWorker; i++ {
		logger.Infoln(i)
	}

	<-c
}
