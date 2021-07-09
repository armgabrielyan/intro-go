package sleep

import "time"

func Sleep(d time.Duration) {
	<-time.After(d)
}
