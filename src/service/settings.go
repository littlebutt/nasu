package service

import (
	"nasu/src/db"
	"strconv"
)

func ChangeHashPrefix(hashPrefix int) bool {
	if hashPrefix < 1 || hashPrefix > 32 {
		return false
	}
	return db.NasuMetaRepo.UpdateNasuMetaByType("HASH_PREFIX", strconv.Itoa(hashPrefix))
}

func ChangeMaxFileSize(size int) bool {
	if size <= 0 {
		return false
	}
	return db.NasuMetaRepo.UpdateNasuMetaByType("MAX_FILE_SIZE", strconv.Itoa(size))
}
