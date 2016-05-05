package libs

import (
	"fmt"
	. "gitlab.caishuo.com/ruby/go-data-client/global"
	"os"
	"os/signal"
)

func RegisterSignal() {

	signals := make(chan os.Signal)
	signal.Notify(signals)

	SetSignals(signals)
	fmt.Printf("Init Signal Ok\n")
}
