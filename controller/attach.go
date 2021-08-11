package controller

import (
	"fmt"
	"go-im/util"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	os.MkdirAll("./mnt", os.ModePerm)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	UploadLocal(writer, request)
}

func UploadLocal(writer http.ResponseWriter, request *http.Request) {
	//获得上传的源文件s
	srcfile, head, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, 40001, err.Error())
	}

	//创建一个新文件d
	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}

	filename := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int31(), suffix)
	dstfile, err := os.Create("./mnt/" + filename)
	if err != nil {
		util.RespFail(writer, 40002, err.Error())
		return
	}

	//将源文件内容copy到新文件
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		util.RespFail(writer, 40003, err.Error())
		return
	}

	url := "/mnt/" + filename
	util.RespOk(writer, url)
}
