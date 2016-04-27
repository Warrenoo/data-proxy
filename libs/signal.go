package libs

import (
	"os"
	"os/signal"
)

func RegisterSignal() chan os.Signal {
	signals := make(chan os.Signal)
	signal.Notify(signals)
	return signals
}
