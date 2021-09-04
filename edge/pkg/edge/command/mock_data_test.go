package command

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrSend = errors.New("send error")
)

type supportMock struct {
	messages []string
}

func (m *supportMock) NotifyInfo(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

func (m *supportMock) NotifyError(format string, args ...interface{}) {
	m.messages = append(m.messages, fmt.Sprintf(format, args...))
}

// WaitCommandSender .
func WaitCommandSender(ret *CommandStream) (bool, bool, bool, bool, bool, bool, bool) {
	var wg sync.WaitGroup
	var resArm, resDisarm, resStart, resPause, resTakeOff, resLand, resReturn bool

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-ret.ArmStream:
				if !ok {
					return
				}
				resArm = true
			case _, ok := <-ret.DisarmStream:
				if !ok {
					return
				}
				resDisarm = true
			case _, ok := <-ret.StartStream:
				if !ok {
					return
				}
				resStart = true
			case _, ok := <-ret.PauseStream:
				if !ok {
					return
				}
				resPause = true
			case _, ok := <-ret.TakeOffStream:
				if !ok {
					return
				}
				resTakeOff = true
			case _, ok := <-ret.LandStream:
				if !ok {
					return
				}
				resLand = true
			case _, ok := <-ret.ReturnStream:
				if !ok {
					return
				}
				resReturn = true
			}
		}
	}()

	wg.Wait()

	return resArm, resDisarm, resStart, resPause, resTakeOff, resLand, resReturn
}
