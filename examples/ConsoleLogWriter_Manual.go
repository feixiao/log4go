package main

import (
	//"time"
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
	// ConsoleLogWriter 使用带缓存的chan进行输出的管理，所以不进行sleep操作
	// 控制台没有输出，因为我们不知道什么时候channel会进行处理
	//time.Sleep(1*time.Second)
	log.Close()
}
