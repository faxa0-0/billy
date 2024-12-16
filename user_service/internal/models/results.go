package models

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}
type CreatedResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
