package db

import "xorm.io/xorm"

func NewEngine(path string) *xorm.Engine {
	engine, err := xorm.NewEngine("sqlite3", path)
	if err != nil {
		return nil
	} else {
		_ = engine.Sync(new(NasuMeta), new(NasuFile))
		return engine
	}
}
