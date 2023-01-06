package db

import (
	"nasu/src/context"
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

func InsertNasuMetaIfNotExistedByMetaType(nasuMeta *NasuMeta) bool {
	exist, err := context.NasuContext.XormEngine.Exist(&NasuMeta{
		MetaType: nasuMeta.MetaType,
	})
	if err != nil {
		return false
	}
	if !exist {
		_, err := context.NasuContext.XormEngine.Insert(nasuMeta)
		if err != nil {
			return false
		}
	}
	return true
}

func QueryNasuMetaByType(metaType string) *NasuMeta {
	var nasuMeta *NasuMeta = &NasuMeta{}
	res, err := context.NasuContext.XormEngine.Where("meta_type = ?", metaType).Get(nasuMeta)
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to get nasu_meta, metaType: ", metaType, ", err: ", err.Error())
		return nil
	}
	if res {
		return nasuMeta
	} else {
		return nil
	}
}

func QueryNasuMetasByType(metaType string) []NasuMeta {
	var nasuMetas []NasuMeta = make([]NasuMeta, 0)
	err := context.NasuContext.XormEngine.Where("meta_type = ?", metaType).Find(&nasuMetas)
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to query nasu_metas, meta_type: ",
			metaType, " err: ", err.Error())
	}
	return nasuMetas
}

func UpdateNasuMetaByType(metaType string, metaValue string) bool {
	var nasuMeta NasuMeta = NasuMeta{}
	nasuMeta.MetaType = metaType
	nasuMeta.MetaValue = metaValue
	updated, err := context.NasuContext.XormEngine.Update(&nasuMeta, &NasuMeta{MetaType: metaType})
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to update nasu_meta, updated: ", updated,
			", err: ", err.Error(), ", metaType: ", metaType)
		return false
	}
	return true
}

func InsertNasuMeta(nasuMeta *NasuMeta) bool {
	_, err := context.NasuContext.XormEngine.Insert(nasuMeta)
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to insert nasu_meta, err: ", err.Error())
		return false
	}
	return true
}

func DeleteNasuMetaByMetaTypeAndMetaValue(metaType string, metaValue string) bool {
	_, err := context.NasuContext.XormEngine.Where("meta_type = ?", metaType).
		And("meta_value = ?", metaValue).Delete(&NasuMeta{})
	if err != nil {
		context.NasuContext.Logger.Warn("[Nasu-db] Fail to delete nasu_meta, err: ", err.Error())
		return false
	}
	return true
}
