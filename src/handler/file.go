package handler

import (
	"github.com/gin-gonic/gin"
	"nasu/src/service"
	"net/http"
	"strconv"
	"strings"
)

func HandleOverallMetaInfo(c *gin.Context) {
	resultMap := service.OverallMetaInfo()
	c.JSON(http.StatusOK, resultMap)
}

func HandleOverallLabelInfo(c *gin.Context) {
	resultMap := service.OverallLabelInfo()
	c.JSON(http.StatusOK, resultMap)
}

func HandleOverallTagInfo(c *gin.Context) {
	resultMap := service.OverallTagInfo()
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
