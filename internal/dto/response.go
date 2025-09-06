package dto

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

func Success(data interface{}, msg string) Response {
	return Response{
		Status:  "success",
		Message: msg,
		Data:    data,
	}
}

func SuccessWithMeta(data interface{}, msg string, meta *Meta) Response {
	return Response{
		Status:  "success",
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func Fail(err string) Response {
	return Response{
		Status: "error",
		Error:  err,
	}
}
