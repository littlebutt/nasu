package service

import (
	"github.com/littlebutt/nasu/src/context"
	"github.com/littlebutt/nasu/src/db"
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

func ChangeTokenTtl(tokenTtl int64) bool {
	if tokenTtl <= 0 {
		return false
	}
	context.NasuContext.TokenTTL = tokenTtl
	return db.NasuMetaRepo.UpdateNasuMetaByType("TOKEN_TTL", strconv.Itoa(int(tokenTtl)))
}
