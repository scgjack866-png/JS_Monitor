package utils

type Respon struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func SuccessRespon(data interface{}) Respon {
	var respon Respon
	respon.Code = "00000"
	respon.Data = data
	respon.Msg = "一切OK!"
	return respon
}

func FailedRespon(msg string) Respon {
	var respon Respon
	respon.Code = "00001"
	respon.Data = ""
	respon.Msg = msg
	return respon
}

func FailedTokenRespon(msg string) Respon {
	var respon Respon
	respon.Code = "A0230"
	respon.Data = ""
	respon.Msg = msg
	return respon
}
