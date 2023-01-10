package service

import (
	"github.com/golang/mock/gomock"
	"github.com/littlebutt/nasu/src/db"
	"testing"
	"time"
)

func TestOverallFileInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuMetaStore(ctl)
	mockDB.EXPECT().QueryNasuMetasByType(gomock.Eq("FILENAME")).Return([]db.NasuMeta{
		{GmtModified: time.Now(), MetaValue: "foo.txt"},
		{GmtModified: time.Now(), MetaValue: "bar.txt"},
	})

	db.NasuMetaRepo = mockDB
	res := OverallFileInfo()
	if len(res) != 2 {
		t.Fail()
	}

}

func TestOverallLabelInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuFileStore(ctl)
	mockDB.EXPECT().QueryNasuFiles().Return([]db.NasuFile{
		{Labels: "foo,bar"},
	})
	db.NasuFileRepo = mockDB
	res := OverallLabelInfo()
	if len(res) != 2 {
		t.Fail()
	}
}

func TestOverallTagInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuFileStore(ctl)
	mockDB.EXPECT().QueryNasuFiles().Return([]db.NasuFile{
		{Tags: "foo,bar"},
	})
	db.NasuFileRepo = mockDB
	res := OverallTagInfo()
	if len(res) != 2 {
		t.Fail()
	}
}

func TestOverallExtensionInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuFileStore(ctl)
	mockDB.EXPECT().QueryNasuFiles().Return([]db.NasuFile{
		{Extension: "foo"},
	})
	db.NasuFileRepo = mockDB
	res := OverallExtensionInfo()
	if len(res) != 1 {
		t.Fail()
	}
}

func TestUploadFile(t *testing.T) {
	t.Skip()
}

func TestListFilesByCondition(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB := db.NewMockNasuFileStore(ctl)
	mockDB.EXPECT().QueryNasuFilesByCondition(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return([]db.NasuFile{})
	db.NasuFileRepo = mockDB
	ListFilesByCondition("", "", []string{}, []string{}, "", "", 0, 0)
}

func TestModifyFile(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockDB1 := db.NewMockNasuFileStore(ctl)
	mockDB2 := db.NewMockNasuMetaStore(ctl)
	gomock.InOrder(
		mockDB1.EXPECT().QueryNasuFileById(gomock.Eq(int64(1))).Return(&db.NasuFile{
			Id:       1,
			Filename: "foo.txt",
		}),
		mockDB1.EXPECT().UpdateNasuFile(gomock.Any()).Return(true),
		mockDB2.EXPECT().DeleteNasuMetaByMetaTypeAndMetaValue(gomock.Eq("FILENAME"), gomock.Any()).Return(true),
		mockDB2.EXPECT().InsertNasuMeta(gomock.Any()).Return(true),
	)
	db.NasuFileRepo = mockDB1
	db.NasuMetaRepo = mockDB2
	res := ModifyFile(db.NasuFile{Id: 1, Filename: "bar.txt"})
	if !res {
		t.Fail()
	}
}

func TestDeleteFile(t *testing.T) {
	t.Skip()
}
