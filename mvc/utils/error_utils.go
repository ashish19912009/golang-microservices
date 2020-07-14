package utils

//ApplicationError Package
type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
