package errorcode

const (
	// 成功
	Succeed = 1
	// 参数验证错误
	ParameterError = 2
	// 数据库写入错误
	InsertError = 3
	// 数据库修改错误
	UpdateError = 4
	// 重复提交
	Repeat = 5
	// 数据库查询错误
	FindError = 6
	// 登录密码错误
	PwdError = 7
	// 删除错误错误
	DeleteError = 8
	// 导出标签错误
	ExportTagError = 9
)
