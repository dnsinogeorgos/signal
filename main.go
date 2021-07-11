package signal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Signal struct {
	Signal  syscall.Signal
	Msg     string
	Exit    bool
	Code    int
	Handler func()
}

func Handle(ss []Signal) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh)

	go Listen(signalCh, ss)
}

func Listen(c chan os.Signal, ss []Signal) {
	for {
		rs := <-c

		for _, sig := range ss {
			if rs == sig.Signal {
				HandleSignal(sig)
				break
			}
		}
	}
}

func HandleSignal(s Signal) {
	log.Printf(s.Msg)
	if s.Handler != nil {
		s.Handler()
	}
	if s.Exit {
		os.Exit(s.Code)
	}
}
