// sighandler is a simple library for handling interrupt signals
package sighandler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Sighandler looks for hup and term signals
type Sighandler struct {
	hupChan  chan os.Signal
	termChan chan os.Signal
}

// ExitFunc is function that takes no parameter and expects no return value which is called when an interrupt hup or term signal is detected
type ExitFunc func()

// NewSigHandler returns an instance of SigHandler
func NewSigHandler() *Sighandler {
	s := Sighandler{}
	s.hupChan = make(chan os.Signal, 1)
	s.termChan = make(chan os.Signal, 1)
	return &s
}

// ListenForSignals - takes an exit function and fires it when one of the watched signals is fired, this function starts a go routine to listen for the signals
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
