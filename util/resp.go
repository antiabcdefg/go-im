package util

import (
	"encoding/json"
	"go-im/args"
	"log"
	"net/http"
)

type ReturnType struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Resp(writer http.ResponseWriter, status string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	r := ReturnType{
		Status: status,
		Data:   data,
	}
	ret, err := json.Marshal(r)
	if err != nil {
		log.Println(err.Error())
	}
	writer.Write(ret)
}

func RespOk(writer http.ResponseWriter, data interface{}) {
	Resp(writer, "sucess", data)
}

func RespFail(writer http.ResponseWriter, code int, msg string) {
	r := args.ErrorArg{
		Code: code,
		Msg:  msg,
	}

	Resp(writer, "fail", r)
}

func RespList(writer http.ResponseWriter, data interface{}, total interface{}) {
	r := struct {
		Rows  interface{} `json:"rows"`
		Total interface{} `json:"total"`
	}{Rows: data, Total: total}

	Resp(writer, "sucess", r)
}
