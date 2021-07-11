package signal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Signal communicates the signal handling behaviour to the Handle function.
type Signal struct {
	Signal  syscall.Signal
	Msg     string
	Exit    bool
	Code    int
	Handler func()
	sem     chan struct{}
}

// Handle sets up the wiring and starts a goroutine to listen in the background. It
// receives a slice of signals to handle. Signals are iterated in the order provided and
// only the first signal that matches is handled.
func Handle(ss []*Signal) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh)

	go listen(signalCh, ss)
}

// listen iterates over the list of provided signals to act upon on each invocation.
func listen(c chan os.Signal, ss []*Signal) {
	for _, s := range ss {
		s.sem = make(chan struct{}, 1)
	}

	for {
		rs := <-c

		for _, s := range ss {
			if rs == s.Signal {
				if len(s.sem) != 0 {
					log.Printf("ignoring %s, a handler is in progress", s.Signal.String())
					break
				}

				s.sem <- struct{}{}
				go handleSignal(s)
				break
			}
		}
	}
}

// handleSignal executes the behaviour described in the Signal value.
func handleSignal(s *Signal) {
	log.Printf(s.Msg)
	if s.Handler != nil {
		s.Handler()
	}
	if s.Exit {
		os.Exit(s.Code)
	}
	<-s.sem
}
