package service

import (
	"nasu/src/db"
	"strconv"
	"time"
)

const DEFAULT_PASSWORD = "21232f297a57a5a743894a0e4a801fc3"

func Login(password string) (success bool, isFirst bool, token string) {
	var passwordFromDb string
	if nasuMeta := db.QueryNasuMetaByType("PASSWORD"); nasuMeta != nil {
		passwordFromDb = nasuMeta.MetaValue
	}
	if passwordFromDb == password {
		success = true
		token = password + "+" + strconv.Itoa(int(time.Now().Unix()))
		isFirst = password == DEFAULT_PASSWORD
		return
	} else {
		success = false
		return
	}
}

func ChangePassword(oldPassword string, newPassword string) (success bool) {
	var passwordFromDb string
	if nasuMeta := db.QueryNasuMetaByType("PASSWORD"); nasuMeta != nil {
		passwordFromDb = nasuMeta.MetaValue
	}
	if passwordFromDb == oldPassword {
		return db.UpdateNasuMetaByType("PASSWORD", newPassword)
	} else {
		return false
	}
}
