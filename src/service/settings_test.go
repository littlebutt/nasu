package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"nasu/src/db"
	"testing"
)

func TestChangeHashPrefix(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuMetaStore(ctl)
	mockDB.EXPECT().UpdateNasuMetaByType(gomock.Eq("HASH_PREFIX"), gomock.Any()).Return(true)
	db.NasuMetaRepo = mockDB
	res := ChangeHashPrefix(1)
	assert.True(t, res)
}

func TestChangeMaxFileSize(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuMetaStore(ctl)
	mockDB.EXPECT().UpdateNasuMetaByType(gomock.Eq("MAX_FILE_SIZE"), gomock.Any()).Return(true)
	db.NasuMetaRepo = mockDB
	res := ChangeMaxFileSize(1)
	assert.True(t, res)
}

func TestChangeTokenTtl(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuMetaStore(ctl)
	mockDB.EXPECT().UpdateNasuMetaByType(gomock.Eq("TOKEN_TTL"), gomock.Any()).Return(true)
	db.NasuMetaRepo = mockDB
	res := ChangeTokenTtl(1)
	assert.True(t, res)
}
