package errors

const (
	OK                  = 0
	CodeParameterErrors = iota + 10000
	CodeProcessingFailure
	CodeUserNotLoggedIn
	CodeInvalidToken
	CodeServerError
	CodeUserRegistering
	CodeUserRegistered
	CodeUnFound
	CodeServiceUnavailable
	CodeUnauthorized
	CodeIncorrectPassword
)

var errorText = map[int]string{
	OK:                     "",
	CodeParameterErrors:    "参数错误",
	CodeUserNotLoggedIn:    "用户未登录",
	CodeInvalidToken:       "错误token",
	CodeProcessingFailure:  "处理失败",
	CodeServerError:        "server error",
	CodeUserRegistered:     "该用户已注册",
	CodeUserRegistering:    "用户注册中",
	CodeUnFound:            "未找到数据",
	CodeServiceUnavailable: "服务不可用",
	CodeUnauthorized:       "unauthorized",
	CodeIncorrectPassword:  "错误的帐号或密码",
}

var (
	ParameterErrors     = New(CodeParameterErrors, ErrorText(CodeParameterErrors))
	ProcessingFailure   = New(CodeProcessingFailure, ErrorText(CodeProcessingFailure))
	InternalServerError = New(CodeServerError, ErrorText(CodeServerError))
	UnFound             = New(CodeUnFound, ErrorText(CodeUnFound))
	ServiceUnavailable  = New(CodeServiceUnavailable, ErrorText(CodeServiceUnavailable))
	UserRegistered      = New(CodeUserRegistered, ErrorText(CodeUserRegistered))
	Unauthorized        = New(CodeUnauthorized, ErrorText(CodeUnauthorized))
	IncorrectPassword   = New(CodeIncorrectPassword, ErrorText(CodeIncorrectPassword))
)

func ErrorText(code int) string {
	return errorText[code]
}
