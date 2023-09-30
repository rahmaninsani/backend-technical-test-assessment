package web

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message,omitempty" metadata:",optional"`
	Data    interface{} `json:"data,omitempty" metadata:",optional"`
}
