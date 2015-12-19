package main

import (
	//"bufio"
	//"fmt"
	//"io"
	//"os"
	"time"
	l4g "github.com/feixiao/log4go"
)


const (
	filename = "flw.log"
)

func main() {
	// 默认输出在控制台,分析源码我们知道，他支持多个目标输出，添加Filter即可
	// 一共提供了三个Filter分别在filelog.go,termlog.go,socklog.go中实现
	log := l4g.NewDefaultLogger(l4g.FINEST)

	log.Close()

	/* Can also specify manually via the following: (these are the defaults) */
	flw := l4g.NewFileLogWriter(filename, false)
	flw.SetFormat("[%D %T] [%L] (%S) %M")
	flw.SetRotate(false)
	flw.SetRotateSize(0)
	flw.SetRotateLines(0)
	flw.SetRotateDaily(false)
	log.AddFilter("file", l4g.FINE, flw)

	// Log some experimental messages
	log.Finest("Everything is created now (notice that I will not be printing to the file)")
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
	log.Critical("Time to close out!")


	time.Sleep(1*time.Second)  // 等待writer的channel被处理，否则我们看不到输出
	// Close the log
	log.Close()
}
