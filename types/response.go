package types

const (
	SUCCESS int = 0
	FAIL    int = 1
)

type Response struct {
	Code         int         `json:"code"`
	ErrorMessage string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Code: SUCCESS,
		Data: data,
	}
}

func NewErrResponse(err error) Response {
	return Response{
		Code:         FAIL,
		ErrorMessage: err.Error(),
	}
}
