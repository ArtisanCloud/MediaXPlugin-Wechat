package contract

type PublishRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type PublishResponse struct {
	BaseResponse
}
