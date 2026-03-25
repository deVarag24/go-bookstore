package apiResponse

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Response
	Data interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Response
	Error interface{} `json:"error,omitempty"`
}

func NewSuccessResponse(msg string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Response: Response{
			Status:  true,
			Message: msg,
		},
		Data: data,
	}
}

func NewErrorResponse(msg string, err interface{}) ErrorResponse {
	return ErrorResponse{
		Response: Response{
			Status:  false,
			Message: msg,
		},
		Error: err,
	}
}
