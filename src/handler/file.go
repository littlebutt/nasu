package handler

import (
	"github.com/gin-gonic/gin"
	"nasu/src/db"
	"nasu/src/service"
	"net/http"
	"strconv"
	"strings"
)

func HandleOverallFileInfo(c *gin.Context) {
	resultSlice := service.OverallFileInfo()
	c.JSON(http.StatusOK, gin.H{
		"filename": resultSlice,
	})
}

func HandleOverallLabelInfo(c *gin.Context) {
	resultMap := service.OverallLabelInfo()
	c.JSON(http.StatusOK, resultMap)
}

func HandleOverallTagInfo(c *gin.Context) {
	resultMap := service.OverallTagInfo()
	c.JSON(http.StatusOK, resultMap)
}

func HandleOverallExtensionInfo(c *gin.Context) {
	resultMap := service.OverallExtensionInfo()
	c.JSON(http.StatusOK, resultMap)
}

func HandleUploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"reason":  "Cannot find file",
		})
		return
	}

	filename := c.PostForm("filename")
	labels := strings.Split(c.PostForm("labels"), ":")
	tags := strings.Split(c.PostForm("tags"), ":")
	uploadTime := c.PostForm("uploadTime")
	extension := ""
	if len(strings.Split(filename, ".")) == 2 {
		extension = strings.Split(filename, ".")[1]
	}
	success, reason := service.UploadFile(file, filename, labels, tags, uploadTime, extension)
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"reason":  reason,
	})
}

func HandleListFilesByCondition(c *gin.Context) {
	filename := c.Query("filename")
	extension := c.Query("extension")
	labels := make([]string, 0)
	if _labels := c.Query("labels"); _labels != "" {
		labels = strings.Split(_labels, ":")
	}
	tags := make([]string, 0)
	if _tags := c.Query("tags"); _tags != "" {
		tags = strings.Split(_tags, ":")
	}
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))

	resultMap := service.ListFilesByCondition(filename, extension, labels, tags, startTime, endTime, pageSize, pageNum)
	c.JSON(http.StatusOK, resultMap)
}

func HandleModifyFile(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "id字段不正确",
		})
	}
	filename := c.PostForm("filename")
	labels := strings.Replace(c.PostForm("labels"), ":", ",", -1)
	tags := strings.Replace(c.PostForm("tags"), ":", ",", -1)
	res := service.ModifyFile(db.NasuFile{
		Id:       int64(id),
		Filename: filename,
		Labels:   labels,
		Tags:     tags,
	})
	c.JSON(http.StatusOK, gin.H{
		"success": res,
	})
}

func HandleDeleteFile(c *gin.Context) {
	filename := c.PostForm("filename")
	res := service.DeleteFile(filename)
	c.JSON(http.StatusOK, gin.H{
		"success": res,
	})
}
