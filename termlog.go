// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"io"
	"os"
	"fmt"
)



// This is the standard writer that prints to standard output.
//type ConsoleLogWriter  chan *LogRecord

type ConsoleLogWriter struct {
	logChan  chan *LogRecord
	wg       WaitGroupWrapper
	stdout   io.Writer
}

// This creates a new ConsoleLogWriter
func NewConsoleLogWriter() *ConsoleLogWriter {
	w := &ConsoleLogWriter{
		logChan: make(chan *LogRecord, LogBufferLength),
		stdout: os.Stdout,
	}
	w.wg.Wrap(func(){w.run()})
	return w
}

func (w *ConsoleLogWriter) run() {

	var timestr string
	var timestrAt int64

	out := w.stdout

	for rec := range w.logChan {
		if at := rec.Created.UnixNano() / 1e9; at != timestrAt {
			timestr, timestrAt = rec.Created.Format("01/02/06 15:04:05"), at
		}
		fmt.Fprint(out, "[", timestr, "] [", levelStrings[rec.Level], "] ", rec.Message, "\n")
	}
}

// This is the ConsoleLogWriter's output method.  This will block if the output
// buffer is full.
func (w *ConsoleLogWriter) LogWrite(rec *LogRecord) {
	w.logChan <- rec
}

// Close stops the logger from sending messages to standard output.  Attempts to
// send log messages to this logger after a Close have undefined behavior.
func (w *ConsoleLogWriter) Close() {
	close(w.logChan)
	w.wg.Wait()
}
