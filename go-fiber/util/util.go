package util

import "strings"

type JResponse struct {
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
	Success bool   `json:"success"`

	Body interface{} `json:"body,omitempty"`
}

func NewJResponse(err error, data interface{}) JResponse {
	jres := JResponse{"", 0, true, data}
	if err != nil {
		jres.Success = false
		jres.Error = err.Error()
		jres.Code = GetErrorCode(err)
	}
	return jres
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
