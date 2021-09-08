package time

import (
	"edge-px4/pkg/edge/domain/common"
	"time"
)

// NewTicker .
func NewTicker(d time.Duration) common.Ticker {
	return &tickerTime{
		t: time.NewTicker(d),
	}
}

type tickerTime struct {
	t *time.Ticker
}

func (t *tickerTime) Tick() <-chan time.Time {
	return t.t.C
}

func (t *tickerTime) Stop() {
	t.t.Stop()
}
