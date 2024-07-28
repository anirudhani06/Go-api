package helpers

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}
