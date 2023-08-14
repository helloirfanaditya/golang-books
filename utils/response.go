package utils

type ResponseSuccess struct {
	Code    int         `json:"code,omitempty" binding:"omitempty"`
	Message string      `json:"message,omitempty" binding:"omitempty"`
	Data    interface{} `json:"data"`
}

type MessageError struct {
	Message string `json:"message,omitempty" binding:"omitempty"`
}

type ResponseError struct {
	Code    int          `json:"code,omitempty" binding:"omitempty"`
	Message string       `json:"message,omitempty" binding:"omitempty"`
	Data    MessageError `json:"data,omitempty" binding:"omitempty"`
}

func ResSuccess(data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Code:    200,
		Message: "true",
		Data:    data,
	}
}

func ResError(data string) ResponseError {
	return ResponseError{
		Code:    400,
		Message: "false",
		Data: MessageError{
			Message: data,
		},
	}
}
