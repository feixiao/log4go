package main

import (
	l4g "github.com/feixiao/log4go"
	"time"
)

func main() {
	// Load the configuration (isn't this easy?)
	l4g.LoadConfiguration("example.xml")

	// And now we're ready!
	l4g.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
	time.Sleep(10*time.Millisecond)
	l4g.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
	l4g.Info("About that time, eh chaps?")

	time.Sleep(1*time.Second)
	l4g.Close()
}
