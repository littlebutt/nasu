package db

import (
	"github.com/littlebutt/nasu/src/log"
	"xorm.io/xorm"
)

//go:generate mockgen -source=./nasu_meta_store.go -destination=./nasu_meta_store_mock.go -package=db
type NasuMetaStore interface {
	InsertNasuMetaIfNotExistedByMetaType(nasuMeta *NasuMeta) bool
	QueryNasuMetaByType(metaType string) *NasuMeta
	QueryNasuMetasByType(metaType string) []NasuMeta
	UpdateNasuMetaByType(metaType string, metaValue string) bool
	InsertNasuMeta(nasuMeta *NasuMeta) bool
	DeleteNasuMetaByMetaTypeAndMetaValue(metaType string, metaValue string) bool
}

var NasuMetaRepo NasuMetaStore

type nasuMetaRepo struct {
	x *xorm.Engine
}

func NewNasuMetaRepo(engine *xorm.Engine) NasuMetaStore {
	return &nasuMetaRepo{x: engine}
}

func (db *nasuMetaRepo) InsertNasuMetaIfNotExistedByMetaType(nasuMeta *NasuMeta) bool {
	exist, err := db.x.Exist(&NasuMeta{
		MetaType: nasuMeta.MetaType,
	})
	if err != nil {
		return false
	}
	if !exist {
		_, err := db.x.Insert(nasuMeta)
		if err != nil {
			return false
		}
	}
	return true
}

func (db *nasuMetaRepo) QueryNasuMetaByType(metaType string) *NasuMeta {
	var nasuMeta *NasuMeta = &NasuMeta{}
	res, err := db.x.Where("meta_type = ?", metaType).Get(nasuMeta)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to get nasu_meta, metaType: %s, err: %s", metaType, err.Error())
		return nil
	}
	if res {
		return nasuMeta
	} else {
		return nil
	}
}

func (db *nasuMetaRepo) QueryNasuMetasByType(metaType string) []NasuMeta {
	var nasuMetas []NasuMeta = make([]NasuMeta, 0)
	err := db.x.Where("meta_type = ?", metaType).Find(&nasuMetas)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to query nasu_metas, meta_type: %s, err: %s",
			metaType, err.Error())
	}
	return nasuMetas
}

func (db *nasuMetaRepo) UpdateNasuMetaByType(metaType string, metaValue string) bool {
	var nasuMeta NasuMeta = NasuMeta{}
	nasuMeta.MetaType = metaType
	nasuMeta.MetaValue = metaValue
	_, err := db.x.Update(&nasuMeta, &NasuMeta{MetaType: metaType})
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to update nasu_meta, metaType: %s, err: %s",
			metaType, err.Error())
		return false
	}
	return true
}

func (db *nasuMetaRepo) InsertNasuMeta(nasuMeta *NasuMeta) bool {
	_, err := db.x.Insert(nasuMeta)
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to insert nasu_meta, err: %s", err.Error())
		return false
	}
	return true
}

func (db *nasuMetaRepo) DeleteNasuMetaByMetaTypeAndMetaValue(metaType string, metaValue string) bool {
	_, err := db.x.Where("meta_type = ?", metaType).
		And("meta_value = ?", metaValue).Delete(&NasuMeta{})
	if err != nil {
		log.Log.Warn("[Nasu-db] Fail to delete nasu_meta, err: %s", err.Error())
		return false
	}
	return true
}
