package fleet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Fleetをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFleetを、指定されたIDでコピーされたことを検証する。
func TestCarbonCopyFleetService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultID1        = DefaultID + "-1"
		DefaultID2        = DefaultID + "-2"
		DefaultOriginalID = DefaultFlightplanID
		DefaultNewID      = DefaultFlightplanID + "-new"
	)

	fleet := Fleet{
		id:           DefaultID1,
		flightplanID: DefaultOriginalID,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{
		id: DefaultID2,
	}
	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultNewID).Return(nil, ErrNotFound)
	repo.On("GetByFlightplanID", DefaultOriginalID).Return(&fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	expectFleet := Fleet{
		id:           DefaultID2,
		flightplanID: DefaultNewID,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          gen,
	}

	a.Equal(repo.fleet, &expectFleet)
	a.Nil(ret)
}

// Fleetをカーボンコピーするドメインサービスをテストする。
// コピー後のIDのFleetのがすでに存在する場合、コピーを行わず
// 正常終了することを検証する。
func TestCopySuccessWhenAlreadyExistsFleetWhenCarbonCopyFleetService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultOriginalID = DefaultFlightplanID
		DefaultNewID      = DefaultFlightplanID + "-new"
	)

	fleet := Fleet{
		id:           DefaultID,
		flightplanID: DefaultNewID,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultNewID).Return(&fleet, nil)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	a.Nil(ret)
}

// Fleetをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFleetの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenCarbonCopyFleetService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultOriginalID = DefaultFlightplanID
		DefaultNewID      = DefaultFlightplanID + "-new"
	)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultNewID).Return(nil, ErrGet)
	repo.On("GetByFlightplanID", DefaultOriginalID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	a.Equal(ret, ErrGet)
}

// Fleetをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFleetの取得がエラーとなった場合、
// 削除が失敗し、エラーが返却されることを検証する。
func TestGetError2WhenCarbonCopyFleetService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultOriginalID = DefaultFlightplanID
		DefaultNewID      = DefaultFlightplanID + "-new"
	)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultNewID).Return(nil, ErrNotFound)
	repo.On("GetByFlightplanID", DefaultOriginalID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	a.Equal(ret, ErrGet)
}

// Fleetをカーボンコピーするドメインサービスをテストする。
// コピーを保存時にリポジトリがエラーとなった場合、
// 保存が失敗し、エラーが返却されることを検証する。
func TestSaveErrorWhenCarbonCopyFleetService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultID1        = DefaultID + "-1"
		DefaultID2        = DefaultID + "-2"
		DefaultOriginalID = DefaultFlightplanID
		DefaultNewID      = DefaultFlightplanID + "-new"
	)

	fleet := Fleet{
		id:           DefaultID1,
		flightplanID: DefaultOriginalID,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{
		id: DefaultID2,
	}
	repo := &repositoryMock{}
	repo.On("GetByFlightplanID", DefaultNewID).Return(nil, ErrNotFound)
	repo.On("GetByFlightplanID", DefaultOriginalID).Return(&fleet, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	a.Equal(ret, ErrSave)
}
