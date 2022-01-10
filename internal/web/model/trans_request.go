package model

// SimpleRequestBody only save the next-route.
type SimpleRequestBody struct {
	NextBody  interface{} `json:"next_body"`
	NextRoute string      `json:"next_route"`
}

type TraceRequestBody struct {
	NextService string `json:"next_service"`
	NextMethod  string `json:"next_method"`
}
