package utils

import (
	jsoniter "github.com/json-iterator/go"
)

// JSONF 格式化任意结构至字符串，方便打印日志
func JSONF(v interface{}) string {
	res, err := jsoniter.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(res)
}
