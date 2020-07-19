package models

//ErrorCode 错误码配置
type ErrorCode int

const (
	RECODE_OK         ErrorCode = 0 // 成功
	RECODE_DBERR      ErrorCode = 4001
	RECODE_NODATA     ErrorCode = 4002
	RECODE_DATAEXIST  ErrorCode = 4003
	RECODE_DATAERR    ErrorCode = 4004
	RECODE_SESSIONERR ErrorCode = 4101
	RECODE_LOGINERR   ErrorCode = 4102
	RECODE_PARAMERR   ErrorCode = 4103
	RECODE_USERERR    ErrorCode = 4104
	RECODE_ROLEERR    ErrorCode = 4105
	RECODE_PWDERR     ErrorCode = 4106
	RECODE_REQERR     ErrorCode = 4201
	RECODE_IPERR      ErrorCode = 4202
	RECODE_THIRDERR   ErrorCode = 4301
	RECODE_IOERR      ErrorCode = 4302
	RECODE_SERVERERR  ErrorCode = 4500
	RECODE_UNKNOWERR  ErrorCode = 4501
)

var recodeText = map[ErrorCode]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
}

// RecodeText 获取错误信息
func RecodeText(code ErrorCode) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
