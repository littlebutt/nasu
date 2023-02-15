package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func IsPathOrFileExisted(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func GetFileMd5(file multipart.File) string {
	md5Hash := md5.New()
	io.Copy(md5Hash, file)
	return hex.EncodeToString(md5Hash.Sum(nil))
}

func TransformSizeToString(size *int64) string {
	level := []string{"B", "KB", "MB", "GB"}
	counter := 0
	for {
		if *size < 1024 || counter == 3 {
			break
		}
		counter += 1
		*size /= 1024
	}
	return strconv.Itoa(int(*size)) + level[counter]
}

func TransformFromPathToLocation(absPath string) string {
	absPath = strings.Replace(absPath, "\\", "/", -1)
	loactionSlice := strings.Split(absPath, "/")
	return "/file/upload/" + strings.Join(loactionSlice[len(loactionSlice)-2:], "/")
}
