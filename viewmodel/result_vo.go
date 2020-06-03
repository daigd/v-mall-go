package viewmodel

// ResultCode 响应结果码
type ResultCode int

const (
	// SUCCESS 处理成功
	SUCCESS ResultCode = 200
	// FAIL 处理失败
	FAIL ResultCode = -1000
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
		Code:    SUCCESS,
		Message: "处理成功",
	}
}

func ResultSuccessData(data interface{}) *ResultVO {
	return &ResultVO{
		Code:    SUCCESS,
		Message: "处理成功",
		Data:    data,
	}
}

// ResultErrorCode 错误码响应
func ResultErrorCode(code ResultCode, message string) *ResultVO {
	return &ResultVO{
		Code:    code,
		Message: message,
	}
}

func ResultErrorMsg(message string) *ResultVO {
	return &ResultVO{
		Code:    FAIL,
		Message: message,
	}
}
