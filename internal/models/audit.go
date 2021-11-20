package models

type Audit struct {
	BaseModel
	UserId     uint64 `json:"user_id"`
	User       User   `json:"user"`
	HttpMethod string `json:"http_method"`
	Path       string `json:"path"`
	Data       string `json:"data"`
}