package controller

import (
	"fmt"
	"go-im/args"
	"go-im/model"
	"go-im/service"
	"go-im/util"
	"math/rand"
	"net/http"
)

var userService service.UserService

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	name := request.PostForm.Get("name")
	passwd := request.PostForm.Get("passwd")

	user, err := userService.Login(name, passwd)
	if err != nil {
		util.RespFail(writer, args.CODE_LOGIN_FAIL, err.Error())
	} else {
		util.RespOk(writer, user)
	}
}

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	name := request.PostForm.Get("name")
	plainpwd := request.PostForm.Get("passwd")
	telephone := fmt.Sprintf("%06d", rand.Int31())
	avatar := "/asset/images/avatar0.png"
	sex := model.SEX_UNKNOW

	user, err := userService.Register(name, plainpwd, telephone, avatar, sex)
	if err != nil {
		util.RespFail(writer, args.CODE_LOGIN_FAIL, err.Error())
	} else {
		util.RespOk(writer, user)
	}
}

func checkToken(userId int64, token string) bool {
	user := userService.FindById(userId)
	return user.Token == token
}
