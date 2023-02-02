package Result

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccesResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}