package contract

type BaseResponse struct {
	Code    int64  `json:"code,omitempty"`
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"msg,omitempty"`
}
