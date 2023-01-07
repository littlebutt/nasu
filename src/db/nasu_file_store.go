package db

import (
	"nasu/src/log"
	"strings"
	"xorm.io/xorm"
)

type NasuFileStore interface {
	InsertNasuFile(nasuFile NasuFile) bool
	QueryNasuFileByHash(hash string) *NasuFile
	QueryNasuFiles() []NasuFile
	QueryNasuFilesByCondition(filename string, extension string, labels []string, tags []string,
		startTime string, endTime string, pageSize int, pageNum int) []NasuFile
	QueryNasuFileById(id int64) *NasuFile
	UpdateNasuFile(nasuFile *NasuFile) bool
	DeleteNasuFileByFilename(filename string) bool
}

var NasuFileRepo NasuFileStore

type nasuFileRepo struct {
	x *xorm.Engine
}

func NewNasuFileRepo(engine *xorm.Engine) NasuFileStore {
	return &nasuFileRepo{x: engine}
}

func (db *nasuFileRepo) InsertNasuFile(nasuFile NasuFile) bool {
	inserted, err := db.x.Insert(&nasuFile)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to insert nasu_file, nasu_file: %s, err: %s",
			nasuFile, err.Error())
		return false
	}
	return inserted != 0
}

func (db *nasuFileRepo) QueryNasuFileByHash(hash string) *NasuFile {
	var nasuFile NasuFile = NasuFile{}
	res, err := db.x.Where("hash = ?", hash).Get(&nasuFile)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to query nasu_file by hash, hash: %s, err: %s", hash, err.Error())
		return nil
	}
	if res {
		return &nasuFile
	} else {
		return nil
	}
}

func (db *nasuFileRepo) QueryNasuFiles() []NasuFile {
	var nasuFiles []NasuFile = make([]NasuFile, 0)
	err := db.x.Find(&nasuFiles)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to query all nasu_file, err: %s", err.Error())
	}
	return nasuFiles
}

func (db *nasuFileRepo) QueryNasuFilesByCondition(filename string, extension string, labels []string, tags []string,
	startTime string, endTime string, pageSize int, pageNum int) []NasuFile {
	nasuFiles := make([]NasuFile, 0)
	var session = db.x.Where("1 = 1")
	if filename != "" {
		session = session.And("filename like ?", "%"+filename+"%")
	}
	if extension != "" {
		session = session.And("extension = ?", extension)
	}
	if len(labels) != 0 {
		for _, label := range labels {
			session = session.And("labels like ?", "%"+label+"%")
		}
	}
	if len(tags) != 0 {
		for _, tag := range tags {
			session = session.And("tags like ?", "%"+tag+"%")
		}
	}
	if startTime != "" {
		session = session.And("upload_time > ?", startTime)
	}
	if endTime != "" {
		session = session.And("upload_time < ?", endTime)
	}
	err := session.Limit(pageSize, pageSize*(pageNum-1)).Find(&nasuFiles)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to query nasu_file by condition, filename: %s, extension: %s, labels %s, "+
			"tags: %s, startTime: %s, endTime: %s, err: %s",
			filename, extension, strings.Join(labels, ","), strings.Join(tags, ","),
			startTime, endTime, err.Error())
	}
	return nasuFiles
}

func (db *nasuFileRepo) QueryNasuFileById(id int64) *NasuFile {
	var nasuFile NasuFile = NasuFile{}
	res, err := db.x.Where("id = ?", id).Get(&nasuFile)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to find target nasu_file, id: %d, err: %s", id, err.Error())
	}
	if res {
		return &nasuFile
	} else {
		return nil
	}
}

func (db *nasuFileRepo) UpdateNasuFile(nasuFile *NasuFile) bool {
	_, err := db.x.Update(nasuFile, &NasuFile{Id: nasuFile.Id})
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to update nasu_file, nasu_file: %s, err: %s", nasuFile, err.Error())
		return false
	}
	return true
}

func (db *nasuFileRepo) DeleteNasuFileByFilename(filename string) bool {
	_, err := db.x.Where("filename = ?", filename).Delete(&NasuFile{})
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to delete nasu_file, filename: %d, err: %s", filename, err.Error())
		return false
	}
	return true
}
