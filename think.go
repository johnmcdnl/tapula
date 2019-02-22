package tapula

import "time"

func thinkTime(d *Duration) {
	if true {
		return
	}
	time.Sleep(d.Duration)
}
