package db

import (
	"time"
)

type NasuMeta struct {
	Id          int64
	GmtCreate   time.Time `xorm:"created 'gmt_create'"`
	GmtModified time.Time `xorm:"updated 'gmt_modified'"`
	MetaType    string    `xorm:"unique(uk_type_value) 'meta_type'"`
	MetaValue   string    `xorm:"unique(uk_type_value) 'meta_value'"`
}

func (nasuMeta NasuMeta) TableName() string {
	return "nasu_meta"
}
