package db

import (
	"nasu/src/context"
	"strings"
	"time"
)

type NasuFile struct {
	Id          int64
	GmtCreate   time.Time `xorm:"created 'gmt_create'"`
	GmtModified time.Time `xorm:"updated 'gmt_modified'"`
	Filename    string    `xorm:"unique 'filename'"`
	Labels      string    `xorm:"labels"`
	Tags        string    `xorm:"tags"`
	Location    string    `xorm:"location"`
	Size        string    `xorm:"varchar(128) 'size'"`
	UploadTime  time.Time `xorm:"upload_time"`
	Extension   string    `xorm:"varchar(128) 'extension'"`
	Hash        string    `xorm:"unique 'hash'"`
}

func (nasuFile NasuFile) TableName() string {
	return "nasu_file"
}

func InsertNasuFile(nasuFile NasuFile) bool {
	inserted, err := context.NasuContext.XormEngine.Insert(&nasuFile)
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to insert nasu_file, nasu_file:",
			nasuFile, " err: ", err.Error())
		return false
	}
	return inserted != 0
}

func QueryNasuFileByHash(hash string) *NasuFile {
	var nasuFile NasuFile = NasuFile{}
	res, err := context.NasuContext.XormEngine.Where("hash = ?", hash).Get(&nasuFile)
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to query nasu_file by hash, hash: ", hash,
			" err: ", err.Error())
		return nil
	}
	if res {
		return &nasuFile
	} else {
		return nil
	}
}

func QueryNasuFiles() []NasuFile {
	var nasuFiles []NasuFile = make([]NasuFile, 0)
	err := context.NasuContext.XormEngine.Find(&nasuFiles)
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to query all nasu_file, err: ", err.Error())
	}
	return nasuFiles
}

func QueryNasuFilesByCondition(filename string, extension string, labels []string, tags []string,
	startTime string, endTime string, pageSize int, pageNum int) []NasuFile {
	nasuFiles := make([]NasuFile, 0)
	var session = context.NasuContext.XormEngine.Where("1 = 1")
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
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to query nasu_file by condition, filename: ", filename,
			" extension: ", extension, " labels: ", strings.Join(labels, ","), " tags: ", strings.Join(tags, ","),
			" startTime: ", startTime, " endTime: ", endTime, " err: ", err.Error())
	}
	return nasuFiles
}

func QueryNasuFileById(id int64) *NasuFile {
	var nasuFile NasuFile = NasuFile{}
	res, err := context.NasuContext.XormEngine.Where("id = ?", id).Get(&nasuFile)
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to find target nasu_file, id: ", id, " err: ", err.Error())
	}
	if res {
		return &nasuFile
	} else {
		return nil
	}
}

func UpdateNasuFile(nasuFile *NasuFile) bool {
	_, err := context.NasuContext.XormEngine.Update(nasuFile, &NasuFile{Id: nasuFile.Id})
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to update nasu_file, nasu_file: ", nasuFile, " err: ", err.Error())
		return false
	}
	return true
}
