package misc

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

func IsPathOrFileExisted(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else {
		if os.IsNotExist(err) {
			return false, nil
		}
	}
	return false, err
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
