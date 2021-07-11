package signal

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Signal communicates the signal handling behaviour to the Handle function.
type Signal struct {
	mu         sync.Mutex
	Signal     syscall.Signal
	Msg        string
	Exit       bool
	Code       int
	Handler    func()
	inProgress bool
}

// Handle sets up the wiring and starts a goroutine to listen in the background. It
// receives a slice of signals to handle. Signals are iterated in the order provided and
// only the first signal that matches is handled.
func Handle(ss []*Signal) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c)

		for {
			rs := <-c

			for _, s := range ss {
				if rs == s.Signal {
					s.mu.Lock()
					if s.inProgress {
						log.Printf("ignoring %s, handler is in progress", s.Signal)
						break
					}
					s.inProgress = true
					go handleSignal(s)
					s.mu.Unlock()
					break
				}
			}
		}
	}()
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

	s.inProgress = false
}
