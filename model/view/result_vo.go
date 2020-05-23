package model

// ResultCode 响应结果码
type ResultCode int

const (
	// Success 处理成功
	Success ResultCode = 200
	// Fail 处理失败
	Fail ResultCode = -1000
)

// ResultVO 服务响应对象
type ResultVO struct {
	Code    ResultCode  `json:"code"`    // 响应码
	Message string      `json:"message"` // 响应描述
	Data    interface{} `json:"data"`    // 响应数据
}

// ResultSuccess 处理成功，默认响应
func ResultSuccess() *ResultVO {
	return &ResultVO{
		Code:    Success,
		Message: "处理成功",
	}
}

// ResultErrorCode 错误码响应
func ResultErrorCode(code ResultCode, message string) *ResultVO {
	return &ResultVO{
		Code:    code,
		Message: message,
	}
}
