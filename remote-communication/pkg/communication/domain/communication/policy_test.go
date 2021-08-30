package communication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// CommandがTakeoffかつTelemetryがDisarmedの場合、ポリシーに合致することを検証する。
func TestTakeoffAndDisarmedWhenIsFollowArmCommandPushPolicy(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	ret := IsFollowArmCommandPushPolicy(CommandTypeTAKEOFF, communication)

	a.True(ret)
}

// CommandがTakeoffかつTelemetryがArmedの場合、ポリシーに合致しないことを検証する。
func TestTakeoffAndArmedWhenIsFollowArmCommandPushPolicy(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)
	communication.telemetry.armed = Armed

	ret := IsFollowArmCommandPushPolicy(CommandTypeTAKEOFF, communication)

	a.False(ret)
}

// CommandがStartかつTelemetryがDisarmedの場合、ポリシーに合致することを検証する。
func TestStartAndDisarmedWhenIsFollowArmCommandPushPolicy(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	ret := IsFollowArmCommandPushPolicy(CommandTypeSTART, communication)

	a.True(ret)
}

// CommandがStartかつTelemetryがArmedの場合、ポリシーに合致しないことを検証する。
func TestStartAndArmedWhenIsFollowArmCommandPushPolicy(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)
	communication.telemetry.armed = Armed

	ret := IsFollowArmCommandPushPolicy(CommandTypeSTART, communication)

	a.False(ret)
}

// CommandがTakeoffでもStartでもない、かつTelemetryがDisarmedの場合、ポリシーに合致しないことを検証する。
func TestOtherAndDisarmedWhenIsFollowArmCommandPushPolicy(t *testing.T) {
	a := assert.New(t)

	gen := &generatorMock{}
	communication := NewInstance(gen, DefaultID)

	ret := IsFollowArmCommandPushPolicy(CommandTypeARM, communication)

	a.False(ret)
}
