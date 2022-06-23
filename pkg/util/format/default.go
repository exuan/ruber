package format

import (
	"time"
)

func UnixTimestamp() int64 {
	return time.Now().Unix()
}

type (
	Output struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	DefaultData struct {
	}
)

func HeaderFilterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}
