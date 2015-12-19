package main

import (
	"time"
)

import l4g "github.com/feixiao/log4go"

func main() {
	log := l4g.NewDefaultLogger(l4g.DEBUG)
	log.Finest("Finest")
	log.Fine("Fine")
	log.Debug("Debug")
	log.Trace("Trace")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	log.Critical("Critical")
	time.Sleep(1*time.Second)
	log.Close()
}
