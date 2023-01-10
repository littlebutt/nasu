package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/littlebutt/nasu/src/context"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func setup() {
	context.NasuContext.Password = "test"
}

func TestAuthRequired(t *testing.T) {
	type test struct {
		desc  string
		input *gin.Context
	}
	req1 := http.Request{
		Header: map[string][]string{
			"Authorization": {"test"},
		},
	}
	req2 := http.Request{
		Header: map[string][]string{
			"Authorization": {"test+0"},
		},
	}
	req3 := http.Request{
		Header: map[string][]string{
			"Authorization": {"testt+" + strconv.Itoa(int(time.Now().Unix()))},
		},
	}
	req4 := http.Request{
		Header: map[string][]string{
			"Authorization": {"test+" + strconv.Itoa(int(time.Now().Unix()))},
		},
	}

	tests := [...]test{
		{desc: "bad authorization format", input: &gin.Context{Request: &req1}},
		{desc: "bad authorization with poor timestamp", input: &gin.Context{Request: &req2}},
		{desc: "bad authorization with poor password", input: &gin.Context{Request: &req3}},
		{desc: "good authorization", input: &gin.Context{Request: &req4}},
	}

	// 出现了，恶心逻辑
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			func() {
				defer func() {
					if err := recover(); err != nil {
						t.Log(err)
					}
				}()
				AuthRequired()(test.input)
			}()
		})
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
