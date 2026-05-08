package utils

type ErrorCode struct {
	Code       int    // 业务错误码
	Message    string // 给用户看的提示信息
	HttpStatus int    // 对应的HTTP状态码
}

// 实现 error 接口（可选，方便作为普通 error 传递）
func (e *ErrorCode) Error() string {
	return e.Message
}

// ====================== 错误定义（全局变量） ======================
// 通用错误（1xxxx）:
var (
	Success = &ErrorCode{
		Code:       0,
		Message:    "Success",
		HttpStatus: 200,
	}
	ErrServer = &ErrorCode{
		Code:       0,
		Message:    "服务器开小差了，请稍后重试",
		HttpStatus: 500,
	}
	ErrParam = &ErrorCode{
		Code:       10001,
		Message:    "参数错误，请检查输入",
		HttpStatus: 200,
	}
	ErrUnauthorized = &ErrorCode{
		Code:       10002,
		Message:    "未登录或登录已过期，请重新登录",
		HttpStatus: 401,
	}
	ErrForbidden = &ErrorCode{
		Code:       10003,
		Message:    "权限不足，无法执行此操作",
		HttpStatus: 403,
	}
	ErrUsernameOrPasswordInvalid = &ErrorCode{
		Code:       10004,
		Message:    "用户名或密码错误",
		HttpStatus: 200,
	}
)
