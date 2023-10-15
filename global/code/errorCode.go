package code

type Code int

const (
	// 统一错误

	Success        = 0    // succeed
	SYSInsideError = -1   // 系统内部错误
	BindError      = 9001 // bind绑定错误

	// user相关错误

	EmailExistError       = 9100 // 邮箱已被注册. 邮箱冲突
	EmailPatternError     = 9101 // 非法邮箱格式
	PasswordPatternError  = 9102 // 非法密码格式
	InvalidUserOrPassword = 9103 // 用户名或密码输入错误
)
