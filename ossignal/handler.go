package ossignal

import (
	"os"
	"os/signal"
)

type SignalHandleFunc func()

type SignalHandler struct {
	regSigMap  map[os.Signal]SignalHandleFunc
	listenChan chan os.Signal
}

func (t *SignalHandler) InitSignalHandler(listenChanSize uint32) error {
	t.regSigMap = make(map[os.Signal]SignalHandleFunc)
	t.listenChan = make(chan os.Signal, listenChanSize)
	return nil
}

func (t *SignalHandler) RegisterSignal(signal os.Signal, handleFunc SignalHandleFunc) bool {
	if _, exist := t.regSigMap[signal]; exist {
		return false
	}

	t.regSigMap[signal] = handleFunc
	return true
}

func (t *SignalHandler) CloseSignalHandler() {
	if t.listenChan == nil {
		return
	}
	close(t.listenChan)
}

func (t *SignalHandler) ListenSignal() {
	var sigs []os.Signal
	for sig := range t.regSigMap {
		sigs = append(sigs, sig)
	}
	signal.Notify(t.listenChan, sigs...)

	for {
		sig, ok := <-t.listenChan
		if !ok {
			break
		}

		handleFunc, exist := t.regSigMap[sig]
		if !exist {
			continue
		}

		handleFunc()
	}

}
