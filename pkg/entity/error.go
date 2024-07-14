package entity

import "encoding/json"

type HandleError struct {
	Msg    string `json:"msg"`
	Status uint   `json:"status"`
}

func NewHandleError(msg string, status uint) HandleError {
	return HandleError{
		Msg:    msg,
		Status: status,
	}
}

func (h HandleError) ConvertToBytes() []byte {
	data, _ := json.Marshal(h)
	return data
}
