package command

import (
	"context"
	"edge-px4/pkg/edge/domain/common"
	"edge-px4/pkg/edge/domain/model"
)

// CommandSender .
func CommandSender(
	ctx context.Context,
	support common.Support,
	stream <-chan struct{},
	adapter func() error,
	name string,
) <-chan struct{} {
	sendExit := make(chan struct{})

	go func() {
		defer close(sendExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("command %s done", name)
				return
			case _, ok := <-stream:
				if !ok {
					support.NotifyInfo("command %s close", name)
					return
				}
				if err := adapter(); err != nil {
					support.NotifyError("command %s error: %v", name, err)
				}
			}
		}
	}()

	return sendExit
}

// MissionSender .
func MissionSender(
	ctx context.Context,
	support common.Support,
	stream <-chan *model.Mission,
	adapter func(*model.Mission) error,
) <-chan struct{} {
	sendExit := make(chan struct{})

	go func() {
		defer close(sendExit)
		for {
			select {
			case <-ctx.Done():
				support.NotifyInfo("mission done")
				return
			case mission, ok := <-stream:
				if !ok {
					support.NotifyInfo("mission close")
					return
				}
				if err := adapter(mission); err != nil {
					support.NotifyError("mission error: %v", err)
				}
			}
		}
	}()

	return sendExit
}
