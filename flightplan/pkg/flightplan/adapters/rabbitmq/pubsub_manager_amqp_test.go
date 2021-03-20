package rabbitmq

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetChannel(t *testing.T) {
	a := assert.New(t)

	connMock := &connectionMockCommon{}
	chMock := &channelMockGetChannel{}
	connMock.On("GetChannel").Return(chMock, nil)

	ch, close, err := getChannel(connMock)

	close()

	a.Equal(chMock, ch)
	a.True(chMock.isClose)
	a.Nil(err)
}

func TestErrorWhenGetChannel(t *testing.T) {
	a := assert.New(t)

	connMock := &connectionMockCommon{}
	errCh := errors.New("channel error")
	connMock.On("GetChannel").Return(nil, errCh)

	ch, close, err := getChannel(connMock)

	a.Nil(ch)
	a.Nil(close)
	a.Equal(err, errCh)
}

func TestSetConsumer(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	excName := "exchange-name"
	queName := "queue-name"
	msg := []byte("receive-message")
	var msgResult []byte

	connMock := &connectionMockCommon{}
	chMock := &channelMockSetConsumer{}
	msgCh := make(chan Message)
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("QueueDeclareAndBind", mock.Anything).Return(nil)
	chMock.On("Consume", mock.Anything, mock.Anything).Return(msgCh, nil)

	psm := NewPubSubManager(connMock)
	err := psm.SetConsumer(
		ctx,
		excName,
		queName,
		func(b []byte) {
			msgResult = b
		},
	)

	msgCh <- msg

	a.Nil(err)
	a.Equal(msg, msgResult)
	a.False(chMock.isClose)
}

func TestChannelCloseWhenMessageChClose(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	excName := "exchange-name"
	queName := "queue-name"
	var msgResult []byte

	var wg sync.WaitGroup

	connMock := &connectionMockCommon{}
	chMock := &channelMockSetConsumer{
		wg: &wg,
	}
	msgCh := make(chan Message)
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("QueueDeclareAndBind", mock.Anything).Return(nil)
	chMock.On("Consume", mock.Anything, mock.Anything).Return(msgCh, nil)

	psm := NewPubSubManager(connMock)

	wg.Add(1)
	err := psm.SetConsumer(
		ctx,
		excName,
		queName,
		func(b []byte) {
			msgResult = b
		},
	)

	close(msgCh)

	wg.Wait()

	a.Nil(err)
	a.Nil(msgResult)
	a.True(chMock.isClose)
}

func TestChannelCloseWhenContextCancel(t *testing.T) {
	a := assert.New(t)

	ctx, cancel := context.WithCancel(context.Background())

	excName := "exchange-name"
	queName := "queue-name"
	var msgResult []byte

	var wg sync.WaitGroup

	connMock := &connectionMockCommon{}
	chMock := &channelMockSetConsumer{
		wg: &wg,
	}
	msgCh := make(chan Message)
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("QueueDeclareAndBind", mock.Anything).Return(nil)
	chMock.On("Consume", mock.Anything, mock.Anything).Return(msgCh, nil)

	psm := NewPubSubManager(connMock)

	wg.Add(1)
	err := psm.SetConsumer(
		ctx,
		excName,
		queName,
		func(b []byte) {
			msgResult = b
		},
	)

	cancel()

	wg.Wait()

	a.Nil(err)
	a.Nil(msgResult)
	a.True(chMock.isClose)
}

func TestFanoutExchangeDeclareErrorWhenSetConsumer(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	excName := "exchange-name"
	queName := "queue-name"
	var msgResult []byte
	errSC := errors.New("set consumer error")

	connMock := &connectionMockCommon{}
	chMock := &channelMockSetConsumer{}
	msgCh := make(chan Message)
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(errSC)
	chMock.On("QueueDeclareAndBind", mock.Anything).Return(nil)
	chMock.On("Consume", mock.Anything, mock.Anything).Return(msgCh, nil)

	psm := NewPubSubManager(connMock)
	err := psm.SetConsumer(
		ctx,
		excName,
		queName,
		func(b []byte) {
			msgResult = b
		},
	)

	a.Equal(err, errSC)
	a.Nil(msgResult)
	a.True(chMock.isClose)
}

func TestQueueDeclareAndBindErrorWhenSetConsumer(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	excName := "exchange-name"
	queName := "queue-name"
	var msgResult []byte
	errSC := errors.New("set consumer error")

	connMock := &connectionMockCommon{}
	chMock := &channelMockSetConsumer{}
	msgCh := make(chan Message)
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("QueueDeclareAndBind", mock.Anything).Return(errSC)
	chMock.On("Consume", mock.Anything, mock.Anything).Return(msgCh, nil)

	psm := NewPubSubManager(connMock)
	err := psm.SetConsumer(
		ctx,
		excName,
		queName,
		func(b []byte) {
			msgResult = b
		},
	)

	a.Equal(err, errSC)
	a.Nil(msgResult)
	a.True(chMock.isClose)
}

func TestConsumeErrorWhenSetConsumer(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	excName := "exchange-name"
	queName := "queue-name"
	var msgResult []byte
	errSC := errors.New("set consumer error")

	connMock := &connectionMockCommon{}
	chMock := &channelMockSetConsumer{}
	connMock.On("GetChannel").Return(chMock, nil)
	chMock.On("FanoutExchangeDeclare", mock.Anything).Return(nil)
	chMock.On("QueueDeclareAndBind", mock.Anything).Return(nil)
	chMock.On("Consume", mock.Anything, mock.Anything).Return(nil, errSC)

	psm := NewPubSubManager(connMock)
	err := psm.SetConsumer(
		ctx,
		excName,
		queName,
		func(b []byte) {
			msgResult = b
		},
	)

	a.Equal(err, errSC)
	a.Nil(msgResult)
	a.True(chMock.isClose)
}

type connectionMockCommon struct {
	mock.Mock
}

func (conn *connectionMockCommon) GetChannel() (Channel, error) {
	ret := conn.Called()
	var ch Channel
	if ret.Get(0) == nil {
		ch = nil
	} else {
		ch = ret.Get(0).(Channel)
	}
	return ch, ret.Error(1)
}

func (conn *connectionMockCommon) Close() error {
	panic("implement me")
}

type channelMockGetChannel struct {
	isClose bool
}

func (ch *channelMockGetChannel) FanoutExchangeDeclare(exchange string) error {
	panic("implement me")
}

func (ch *channelMockGetChannel) QueueDeclareAndBind(exchange, queue string) error {
	panic("implement me")
}

func (ch *channelMockGetChannel) Publish(queue string, message Message) error {
	panic("implement me")
}

func (ch *channelMockGetChannel) Consume(ctx context.Context, queue string) (<-chan Message, error) {
	panic("implement me")
}

func (ch *channelMockGetChannel) Close() error {
	ch.isClose = true
	return nil
}

type channelMockSetConsumer struct {
	mock.Mock
	wg      *sync.WaitGroup
	isClose bool
}

func (ch *channelMockSetConsumer) FanoutExchangeDeclare(exchange string) error {
	ret := ch.Called()
	return ret.Error(0)
}

func (ch *channelMockSetConsumer) QueueDeclareAndBind(exchange, queue string) error {
	ret := ch.Called()
	return ret.Error(0)
}

func (ch *channelMockSetConsumer) Publish(queue string, message Message) error {
	panic("implement me")
}

func (ch *channelMockSetConsumer) Consume(ctx context.Context, queue string) (<-chan Message, error) {
	ret := ch.Called()
	msgCh := make(chan Message)
	if ret.Get(0) == nil {
		msgCh = nil
	} else {
		msgCh = ret.Get(0).(chan Message)
	}
	return msgCh, ret.Error(1)
}

func (ch *channelMockSetConsumer) Close() error {
	defer func() {
		if ch.wg != nil {
			ch.wg.Done()
		}
	}()
	ch.isClose = true
	return nil
}
