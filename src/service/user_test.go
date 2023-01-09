package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"nasu/src/db"
	"testing"
)

func TestLogin(t *testing.T) {
	type test struct {
		desc   string
		input  string
		output struct {
			success bool
			isFirst bool
			token   string
		}
	}
	tests := [...]test{
		{"wrong password", "foo", struct {
			success bool
			isFirst bool
			token   string
		}{success: false, isFirst: false, token: ""}},
		{"login", "bar", struct {
			success bool
			isFirst bool
			token   string
		}{success: true, isFirst: false, token: "bar"}},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuMetaStore(ctl)
	mockDB.EXPECT().QueryNasuMetaByType("PASSWORD").Return(&db.NasuMeta{MetaValue: "bar"}).AnyTimes()
	db.NasuMetaRepo = mockDB
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			success, isFirst, _ := Login(test.input)
			assert.Equal(t, test.output.success, success)
			assert.Equal(t, test.output.isFirst, isFirst)
		})
	}
}

func TestChangePassword(t *testing.T) {
	type test struct {
		desc   string
		input  [2]string
		output bool
	}
	tests := [...]test{
		{"wrong password", [...]string{"foo", "test"}, false},
		{"right password", [...]string{"bar", "test"}, true},
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuMetaStore(ctl)
	mockDB.EXPECT().QueryNasuMetaByType("PASSWORD").Return(&db.NasuMeta{MetaValue: "bar"}).AnyTimes()
	mockDB.EXPECT().UpdateNasuMetaByType("PASSWORD", gomock.Any()).Return(true).AnyTimes()
	db.NasuMetaRepo = mockDB
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.output, ChangePassword(test.input[0], test.input[1]))
		})
	}
}
