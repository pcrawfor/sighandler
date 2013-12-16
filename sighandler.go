package sighandler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Sighandler struct {
	hupChan  chan os.Signal
	termChan chan os.Signal
}

type ExitFunc func()

func NewSigHandler() *Sighandler {
	s := Sighandler{}
	s.hupChan = make(chan os.Signal, 1)
	s.termChan = make(chan os.Signal, 1)
	return &s
}

/*
ListenForSignals - takes an exit function and fires it when one of the watched signals is fired, this function starts a go routine to listen for the signals
*/
func (s *Sighandler) ListenForSignals(exitFunc ExitFunc) {
	signal.Notify(s.hupChan, syscall.SIGHUP)
	signal.Notify(s.termChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-s.termChan:
				fmt.Println("Exiting term interrupt")
				exitFunc()
			case <-s.hupChan:
				fmt.Println("Exiting hup interrupt")
				exitFunc()
			}
		}
	}()
}
