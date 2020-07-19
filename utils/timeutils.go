package utils

import (
	"time"
)

// 获取当前时间 （毫秒）
func NowMillis() int64 {
	nanos := time.Now().UnixNano()
	millis := nanos / int64(time.Millisecond)
	return millis
}

