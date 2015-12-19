package main

import (
	"time"
	l4g "github.com/feixiao/log4go"
)

func main() {
	// 虽然源代码里面说NewLogger()已经弃用，但是在不需要控制台输出的时候，还是有用的
	log := l4g.NewLogger()
	log.AddFilter("network", l4g.FINEST, l4g.NewSocketLogWriter("udp", "localhost:12124"))

	// Run `nc -u -l -p 12124` or similar before you run this to see the following message
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))

	// This makes sure the output stream buffer is written
	time.Sleep(1*time.Second)
	log.Close()
}
