package utils

type RequestType struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type ResponseType struct {
	Msg  string `json:"msg"`
	Code uint32 `json:"code"`
	Data string `json:"data"`
}
