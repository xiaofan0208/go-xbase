package utils

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

func init() {
	beego.AddFuncMap("timeformat", timeformat)
}

func timeformat(in interface{}) (out string) {
	var tInt64 int64
	switch in.(type) {
	case string:
		tInt64, _ = strconv.ParseInt(in.(string), 10, 64)
	case int64:
		tInt64 = in.(int64)
	}
	out = time.Unix(tInt64/1000, 0).Format("2006-01-02 15:04:05")
	return
}
