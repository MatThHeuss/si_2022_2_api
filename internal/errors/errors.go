package errors

type Errors struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
