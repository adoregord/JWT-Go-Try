package domain

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
