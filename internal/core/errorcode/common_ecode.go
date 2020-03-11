package errorcode

var (
	OK = 0 // 正确

	NoModified         = -304 // 木有改动
	RequestErr         = -400 // 请求错误
	Unauthorized       = -401 // 未认证
	AccessDenied       = -403 // 访问权限不足
	NothingFound       = -404 // 啥都木有
	MethodNotAllowed   = -405 // 不支持该方法
	Conflict           = -409 // 冲突
	Canceled           = -498 // 客户端取消请求
	ServerErr          = -500 // 服务器错误
	ServiceUnavailable = -503 // 过载保护,服务暂不可用
	Deadline           = -504 // 服务调用超时
	LimitExceed        = -509 // 超出限制
)
