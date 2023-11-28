package http

// Response ...
type Response struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
	Error       interface{} `json:"error"`
	RequestId   string      `json:"requestId"`
	Message     string      `json:"message"`
}
