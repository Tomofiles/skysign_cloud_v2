package common

import "time"

// Ticker .
type Ticker interface {
	Tick() <-chan time.Time
	Stop()
}
