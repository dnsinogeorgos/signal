# signal
A small package for signal handling.

Each signal iterates over the list of provided signals to act upon.  
Signals are iterated in the order provided and only the first signal that matches is
handled.

### Usage:
install with:
```shell
go get github.com/dnsinogeorgos/signal
```

example:
```go
	signals := []signal.Signal{
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
```
