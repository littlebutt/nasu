package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestCookieRequired(t *testing.T) {
	type test struct {
		desc  string
		input *gin.Context
	}
	cookies := []http.Cookie{
		{Name: "token", Value: "test"},
		{Name: "token", Value: "test 1"},
		{Name: "token", Value: "test bar"},
	}
	var reqs []http.Request
	for _, cookie := range cookies {
		req := http.Request{Header: map[string][]string{}}
		req.AddCookie(&cookie)
		reqs = append(reqs, req)
	}
	tests := [...]test{
		{desc: "test for bad format", input: &gin.Context{Request: &reqs[0]}},
		{desc: "test for bad value", input: &gin.Context{Request: &reqs[1]}},
		{desc: "test for bad value format", input: &gin.Context{Request: &reqs[2]}},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			func() {
				defer func() {
					if err := recover(); err != nil {
						t.Log(err)
					}
				}()
				CookieRequired()(test.input)
			}()
		})
	}
}
