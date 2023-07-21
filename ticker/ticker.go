package ticker

import (
	"bispy-agent/worker"
	"time"
)

func Start() {
	ticker := time.NewTicker(1000 * time.Millisecond)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				worker.Supply(&t)
			}
		}
	}()

	select {}
}
