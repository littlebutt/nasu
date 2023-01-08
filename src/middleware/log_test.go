package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestLogRequired(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	req := http.Request{
		Method: "GET",
		URL:    nil,
	}
	LogRequired()(&gin.Context{Request: &req, Writer: nil})
}
