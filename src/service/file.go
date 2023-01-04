package service

import (
	"io"
	"mime/multipart"
	"nasu/src/context"
	"nasu/src/db"
	"nasu/src/utils"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func OverallMetaInfo() (resultMap map[string]any) {
	resultMap = make(map[string]any)
	nasuMetas := db.QueryNasuMetasByType("FILENAME")
	sort.SliceStable(nasuMetas, func(i, j int) bool {
		return (nasuMetas)[i].GmtModified.Unix() > (nasuMetas)[j].GmtModified.Unix()
	})
	filenames := make([]string, 0)
	for _, nasuMeta := range nasuMetas {
		filenames = append(filenames, nasuMeta.MetaValue)
	}
	resultMap["filenames"] = filenames

	nasuMetas = db.QueryNasuMetasByType("LABEL")
	labels := make([]string, 0)
	for _, nasuMeta := range nasuMetas {
		labels = append(labels, nasuMeta.MetaValue)
	}
	resultMap["labels"] = labels

	nasuMetas = db.QueryNasuMetasByType("TAG")
	tags := make([]string, 0)
	for _, nasuMeta := range nasuMetas {
		tags = append(tags, nasuMeta.MetaValue)
	}
	resultMap["tags"] = tags

	nasuMetas = db.QueryNasuMetasByType("EXTENSION")
	extensions := make([]string, 0)
	for _, nasuMeta := range nasuMetas {
		extensions = append(extensions, nasuMeta.MetaValue)
	}
	resultMap["extension"] = extensions

	return
}

func OverallLabelInfo() (resultMap map[string]int) {
	resultMap = make(map[string]int)
	nasuFiles := db.QueryNasuFiles()
	for _, nasuFile := range nasuFiles {
		labels := strings.Split(nasuFile.Labels, ",")
		for _, label := range labels {
			if _, ok := resultMap[label]; ok {
				resultMap[label] += 1
			} else {
				resultMap[label] = 1
			}
		}
	}
	return
}

func OverallTagInfo() (resultMap map[string]int) {
	resultMap = make(map[string]int)
	nasuFiles := db.QueryNasuFiles()
	for _, nasuFile := range nasuFiles {
		tags := strings.Split(nasuFile.Tags, ",")
		for _, tag := range tags {
			if _, ok := resultMap[tag]; ok {
				resultMap[tag] += 1
			} else {
				resultMap[tag] = 1
			}
		}
	}
	return
}

// UploadFile 上传一个文件到某个路径（路径由文件哈希值决定）。该方法包括预检验重名文件，拷贝文件，更新数据库
func UploadFile(file *multipart.FileHeader, filename string,
	labels []string, tags []string, uploadTime string, extension string) (success bool, reason string) {
	// pre-check filename duplication
	precheckeds := db.QueryNasuMetasByType("FILENAME")
	for _, prechecked := range precheckeds {
		if prechecked.MetaValue == filename {
			return false, "存在同名文件，请重命名后再次上传"
		}
	}
	_uploadTime, err := time.Parse("2006-01-02 15:04:05", uploadTime)
	if err != nil {
		return false, "时间格式错误，请遵守yyyy-MM-dd HH:mm:ss格式"
	}

	// upload file
	src, err := file.Open()
	defer src.Close()
	size := file.Size
	if err != nil {
		return false, "无法打开上传的文件: " + err.Error()
	}
	hash := utils.GetFileMd5(src)
	res := db.QueryNasuFileByHash(hash)
	if res != nil {
		return false, "重复上传相同的文件"
	}
	// TODO: customize hash prefix
	targetPath := filepath.Join(context.NasuContext.ResourcesDir, hash[:1])
	if existed := utils.IsPathOrFileExisted(targetPath); !existed {
		_ = os.Mkdir(targetPath, os.ModePerm)
	}
	location := filepath.Join(targetPath, filename)
	dst, _ := os.Create(location)
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return false, "上传过程中出现未知错误: " + err.Error()
	}

	// update db nasu_file
	var nasuFile = db.NasuFile{}
	nasuFile.Filename = filename
	nasuFile.Labels = strings.Join(labels, ",")
	nasuFile.Tags = strings.Join(tags, ",")
	nasuFile.Location = location
	nasuFile.Size = utils.TransformSizeToString(&size)
	nasuFile.UploadTime = _uploadTime
	nasuFile.Extension = extension
	nasuFile.Hash = hash
	if success := db.InsertNasuFile(nasuFile); !success {
		return false, "数据更新异常"
	}

	// update db nasu_meta
	for _, label := range labels {
		var nasuMeta = db.NasuMeta{MetaType: "LABEL", MetaValue: label}
		db.InsertNasuMeta(nasuMeta)
	}

	for _, tag := range tags {
		var nasuMeta = db.NasuMeta{MetaType: "TAG", MetaValue: tag}
		db.InsertNasuMeta(nasuMeta)
	}

	var nasuMetaFilename = db.NasuMeta{MetaType: "FILENAME", MetaValue: filename}
	db.InsertNasuMeta(nasuMetaFilename)

	var nasuMetaExtension = db.NasuMeta{MetaType: "EXTENSION", MetaValue: extension}
	db.InsertNasuMeta(nasuMetaExtension)
	return true, ""

}

// test script : curl localhost:8080/api/uploadFile -X POST -H "Authorization:21232f297a57a5a743894a0e4a801fc3+1672619245" -F "file=@2022-10-12.log" -F "filename=2022-10-12.log" -F "labels=foo:bar" -F "tags=0" -F "uploadTime=2022-12-30 08:00:00"

func ListFilesByCondition(filename string, extension string, labels []string, tags []string,
	startTime string, endTime string, pageSize int, pageNum int) map[string]any {
	resultMap := make(map[string]any)
	nasuFiles := db.QueryNasuFilesByCondition(filename, extension, labels, tags, startTime, endTime, pageSize, pageNum)
	resultMap["nasuFiles"] = nasuFiles
	resultMap["total"] = len(nasuFiles)
	return resultMap
}
