package sighandler

import (
	"syscall"
	"testing"
	"time"
)

func TestSigHupChan(t *testing.T) {
	s := NewSigHandler()
	passed := false

	go func() {
		s.ListenForSignals(func() {
			passed = true
		})
	}()
	s.hupChan <- syscall.SIGHUP

	time.Sleep(time.Second * 1)
	if !passed {
		t.Error("Exit wasn't triggered")
	}

}

func TestSigTermChan(t *testing.T) {
	s := NewSigHandler()
	passed := false

	go func() {
		s.ListenForSignals(func() {
			passed = true
		})
	}()
	s.termChan <- syscall.SIGINT

	time.Sleep(time.Second * 1)
	if !passed {
		t.Error("Exit wasn't triggered")
	}

}
