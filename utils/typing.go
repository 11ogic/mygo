package utils

type RequestType struct {
	Data string `json:"data"`
	Code uint8  `json:"code"`
	Type string `json:"type"`
}

type ResponseType struct {
	Msg  string `json:"msg"`
	Code uint32 `json:"code"`
	Data string `json:"data"`
}
