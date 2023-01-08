package service

import (
	"io"
	"mime/multipart"
	"nasu/src/context"
	"nasu/src/db"
	"nasu/src/log"
	"nasu/src/utils"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func OverallFileInfo() (resultSlice []string) {
	resultSlice = make([]string, 0)
	nasuMetas := db.NasuMetaRepo.QueryNasuMetasByType("FILENAME")
	sort.SliceStable(nasuMetas, func(i, j int) bool {
		return (nasuMetas)[i].GmtModified.Unix() > (nasuMetas)[j].GmtModified.Unix()
	})
	for _, nasuMeta := range nasuMetas {
		resultSlice = append(resultSlice, nasuMeta.MetaValue)
	}
	return
}

func OverallLabelInfo() (resultMap map[string]int) {
	resultMap = make(map[string]int)
	nasuFiles := db.NasuFileRepo.QueryNasuFiles()
	for _, nasuFile := range nasuFiles {
		labels := nasuFile.GetLabels()
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
	nasuFiles := db.NasuFileRepo.QueryNasuFiles()
	for _, nasuFile := range nasuFiles {
		tags := nasuFile.GetTags()
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

func OverallExtensionInfo() (resultMap map[string]int) {
	resultMap = make(map[string]int)
	nasuFiles := db.NasuFileRepo.QueryNasuFiles()
	for _, nasuFile := range nasuFiles {
		if _, ok := resultMap[nasuFile.Extension]; ok {
			resultMap[nasuFile.Extension] += 1
		} else {
			resultMap[nasuFile.Extension] = 1
		}
	}
	return
}

// UploadFile 上传一个文件到某个路径（路径由文件哈希值决定）。该方法包括预检验重名文件，拷贝文件，更新数据库
func UploadFile(file *multipart.FileHeader, filename string,
	labels []string, tags []string, uploadTime string, extension string) (success bool, reason string) {
	// pre-check filename duplication
	precheckeds := db.NasuMetaRepo.QueryNasuMetasByType("FILENAME")
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
	res := db.NasuFileRepo.QueryNasuFileByHash(hash)
	if res != nil {
		return false, "重复上传相同的文件"
	}
	hashPrefix, _ := strconv.Atoi(db.NasuMetaRepo.QueryNasuMetaByType("HASH_PREFIX").MetaValue)
	targetPath := filepath.Join(context.NasuContext.ResourcesDir, hash[:hashPrefix])
	if existed := utils.IsPathOrFileExisted(targetPath); !existed {
		_ = os.Mkdir(targetPath, os.ModePerm)
	}
	location := filepath.Join(targetPath, filename)
	dst, _ := os.Create(location)
	defer dst.Close()
	src.Seek(0, 0)
	_, err = io.Copy(dst, src)
	if err != nil {
		return false, "上传过程中出现未知错误: " + err.Error()
	}

	// update db nasu_file
	var nasuFile = db.NasuFile{}
	nasuFile.Filename = filename
	nasuFile.Labels = strings.Join(labels, ",")
	nasuFile.Tags = strings.Join(tags, ",")
	nasuFile.Location = utils.TransformFromPathToLocation(location)
	nasuFile.Size = utils.TransformSizeToString(&size)
	nasuFile.UploadTime = _uploadTime
	nasuFile.Extension = extension
	nasuFile.Hash = hash
	if success := db.NasuFileRepo.InsertNasuFile(nasuFile); !success {
		return false, "数据更新异常"
	}

	// update db nasu_meta
	var nasuMetaFilename = db.NasuMeta{MetaType: "FILENAME", MetaValue: filename}
	db.NasuMetaRepo.InsertNasuMeta(&nasuMetaFilename)

	return true, ""
}

// test script : curl localhost:8080/api/uploadFile -X POST -H "Authorization:21232f297a57a5a743894a0e4a801fc3+1672619245" -F "file=@2022-10-12.log" -F "filename=2022-10-12.log" -F "labels=foo:bar" -F "tags=0" -F "uploadTime=2022-12-30 08:00:00"

func ListFilesByCondition(filename string, extension string, labels []string, tags []string,
	startTime string, endTime string, pageSize int, pageNum int) map[string]any {
	resultMap := make(map[string]any)
	nasuFiles := db.NasuFileRepo.QueryNasuFilesByCondition(filename, extension, labels, tags, startTime, endTime, pageSize, pageNum)
	resultMap["nasuFiles"] = nasuFiles
	resultMap["total"] = len(nasuFiles)
	return resultMap
}

func ModifyFile(nasuFile db.NasuFile) bool {
	// pre-check
	if nasuFile.Id == 0 {
		return false
	}
	oldNasuFile := db.NasuFileRepo.QueryNasuFileById(nasuFile.Id)
	if oldNasuFile == nil {
		return false
	}
	// update nasu_file
	extension := ""
	if strs := strings.Split(nasuFile.Filename, "."); len(strs) == 2 {
		extension = strs[1]
	}
	updatedNasuFile := db.NasuFile{
		Id:        nasuFile.Id,
		Filename:  nasuFile.Filename,
		Labels:    nasuFile.Labels,
		Tags:      nasuFile.Tags,
		Extension: extension,
	}
	res := db.NasuFileRepo.UpdateNasuFile(&updatedNasuFile)
	if !res {
		return false
	}

	// update nasu_meta
	res = true
	if nasuFile.Filename != "" {
		res = res && db.NasuMetaRepo.DeleteNasuMetaByMetaTypeAndMetaValue("FILENAME", oldNasuFile.Filename)
		res = res && db.NasuMetaRepo.InsertNasuMeta(&db.NasuMeta{
			MetaType:  "FILENAME",
			MetaValue: nasuFile.Filename,
		})
	}
	return res
}

func DeleteFile(filename string) bool {
	// pre-check
	if filename == "" {
		return false
	}
	// delete file
	nasuFiles := db.NasuFileRepo.QueryNasuFilesByCondition(filename, "", []string{}, []string{},
		"", "", 10, 1)
	if len(nasuFiles) == 0 {
		return false
	}
	location := nasuFiles[0].Location
	err := os.Remove(location)
	if err != nil {
		log.Log.Error("[Nasu-service] Fail to remove file, filename: %s, err: %s", filename, err.Error())
		return false
	}

	// delete nasu_file
	res := db.NasuFileRepo.DeleteNasuFileByFilename(filename)
	if !res {
		return false
	}

	//delete nasu_file
	res = db.NasuMetaRepo.DeleteNasuMetaByMetaTypeAndMetaValue("FILENAME", filename)
	return res
}
