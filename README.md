# signal
A ~~small~~ tiny package for signal handling.

Each invocation iterates over the list of provided signals to act upon.  
Signals are iterated in the order provided and only the first signal that matches is
handled.  
Calls to each handler will be ignored until previous invocations are finished.

### Usage:
Install with:
```shell
go get github.com/dnsinogeorgos/signal
```

Sample code:
```go
package main

import (
	"log"
	"syscall"
	"time"

	"github.com/dnsinogeorgos/signal"
)

func main() {
	signals := []*signal.Signal{
		{
			Signal:  syscall.SIGHUP,
			Msg:     "received SIGHUP... reloading configuration",
			Handler: HandleReload,
		},
		{
			Signal:  syscall.SIGTERM,
			Msg:     "received SIGTERM... shutting down",
			Exit:    true,
			Handler: HandleStop,
		},
	}
	signal.Handle(signals)

	time.Sleep(1 * time.Hour)
}

func HandleReload() {
	// Configuration reload code here
	log.Printf("configuration reloaded")
}

func HandleStop() {
	// Graceful shutdown code here
	log.Printf("shut down complete")
}
```
