package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

const targetFilePath = "../utils/fileutils_test.go"

func TestIsPathOrFileExisted(t *testing.T) {
	type test struct {
		desc   string
		input  string
		expect bool
	}
	tests := [...]test{
		{desc: "test for existing", input: targetFilePath, expect: true},
		{desc: "test for not existing", input: "./nothing", expect: false},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.expect, IsPathOrFileExisted(test.input))
		})
	}
}

func TestGetFileMd5(t *testing.T) {
	file, _ := os.Open(targetFilePath)
	res := GetFileMd5(file)
	assert.NotEmpty(t, res)
}

func TestTransformSizeToString(t *testing.T) {
	file, _ := os.Open(targetFilePath)
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	res := TransformSizeToString(&size)
	assert.NotEmpty(t, res)
}

func TestTransformFromPathToLocation(t *testing.T) {
	filePath, _ := filepath.Abs(targetFilePath)
	res := TransformFromPathToLocation(filePath)
	assert.NotEmpty(t, res)

}
