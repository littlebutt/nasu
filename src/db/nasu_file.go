package db

import (
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

func (nasuFile NasuFile) GetLabels() []string {
	return strings.Split(nasuFile.Labels, ",")
}

func (nasuFile NasuFile) GetTags() []string {
	return strings.Split(nasuFile.Tags, ",")
}
