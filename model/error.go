package model

type MsgCode int

const (
	Ok MsgCode = iota
	Unknown
	ErrorMailInvalid
	ErrorUrlInvalid
	ErrorOversize
	ErrorEmpty
	ErrorParamParse
	ErrorUsernameExist
	ErrorMailExist
	ErrorUserExist
	ErrorLoginFailed
	ErrorLoginExpire
	ErrorRowsNotExist
	ErrorExpired
	ErrorForeverNotSupport
	ErrorInternal
	ErrorJWTExpired
	ErrorRegisterDisabled
)

var ( // todo 使用配置文件
	MsgTable = map[MsgCode]string{
		Ok:                     "Ok",
		Unknown:                "未知",
		ErrorMailInvalid:       "邮箱格式不规范",
		ErrorUrlInvalid:        "链接格式不规范",
		ErrorOversize:          "长度超过限制",
		ErrorEmpty:             "参数不能为空",
		ErrorParamParse:        "参数解析出错",
		ErrorUsernameExist:     "用户名已存在",
		ErrorMailExist:         "邮箱已存在",
		ErrorUserExist:         "用户已存在",
		ErrorLoginFailed:       "登录失败",
		ErrorLoginExpire:       "登录态失效",
		ErrorRowsNotExist:      "记录不存在",
		ErrorExpired:           "已过期",
		ErrorForeverNotSupport: "暂不支持永久有效",
		ErrorInternal:          "服务器内部错误",
		ErrorJWTExpired:        "登录态失效",
		ErrorRegisterDisabled:  "注册已关闭",
	}
)
