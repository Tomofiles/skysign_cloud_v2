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
		DefaultOriginalID = DefaultID
		DefaultNewID      = DefaultID + "-new"
	)

	fleet := Fleet{
		id:           DefaultOriginalID,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultNewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultOriginalID).Return(&fleet, nil)
	repo.On("Save", mock.Anything).Return(nil)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	expectFleet := Fleet{
		id:           DefaultNewID,
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
		DefaultOriginalID = DefaultID
		DefaultNewID      = DefaultID + "-new"
	)

	fleet := Fleet{
		id:           DefaultNewID,
		isCarbonCopy: CarbonCopy,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultNewID).Return(&fleet, nil)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	a.Nil(ret)
}

// Fleetをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFleetの取得がエラーとなった場合、
// コピーが失敗し、エラーが返却されることを検証する。
func TestGetErrorWhenCarbonCopyFleetService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultOriginalID = DefaultID
		DefaultNewID      = DefaultID + "-new"
	)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultNewID).Return(nil, ErrGet)
	repo.On("GetByID", DefaultOriginalID).Return(nil, ErrGet)
	repo.On("Save", mock.Anything).Return(nil)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	a.Equal(ret, ErrGet)
}

// Fleetをカーボンコピーするドメインサービスをテストする。
// 指定されたIDのFleetの取得がエラーとなった場合、
// コピーが失敗し、エラーが返却されることを検証する。
func TestGetError2WhenCarbonCopyFleetService(t *testing.T) {
	a := assert.New(t)

	ctx := context.Background()

	var (
		DefaultOriginalID = DefaultID
		DefaultNewID      = DefaultID + "-new"
	)

	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultNewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultOriginalID).Return(nil, ErrGet)
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
		DefaultOriginalID = DefaultID
		DefaultNewID      = DefaultID + "-new"
	)

	fleet := Fleet{
		id:           DefaultOriginalID,
		isCarbonCopy: Original,
		version:      DefaultVersion,
		newVersion:   DefaultVersion,
		gen:          nil,
	}
	gen := &generatorMock{}
	repo := &repositoryMock{}
	repo.On("GetByID", DefaultNewID).Return(nil, ErrNotFound)
	repo.On("GetByID", DefaultOriginalID).Return(&fleet, nil)
	repo.On("Save", mock.Anything).Return(ErrSave)

	ret := CarbonCopyFleet(ctx, gen, repo, nil, DefaultOriginalID, DefaultNewID)

	a.Equal(ret, ErrSave)
}
