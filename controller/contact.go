package controller

import (
	"go-im/args"
	"go-im/model"
	"go-im/service"
	"go-im/util"
	"net/http"
)

var contactService service.ContactService

func LoadFriend(w http.ResponseWriter, req *http.Request) {
	//定义一个参数结构体
	/*request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	*/
	var arg args.ContactArg
	util.Bind(req, &arg)

	users := contactService.SearchFriend(arg.UserId)
	util.RespList(w, users, len(users))
}

func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)

	comunitys := contactService.SearchComunity(arg.UserId)
	util.RespList(w, comunitys, len(comunitys))
}

func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)

	err := contactService.JoinCommunity(arg.UserId, arg.DstId)

	AddGroupId(arg.UserId, arg.DstId)
	if err != nil {
		util.RespFail(w, 30002, err.Error())
	} else {
		util.RespOk(w, nil)
	}
}

func CreateCommunity(w http.ResponseWriter, req *http.Request) {
	var arg model.Community
	util.Bind(req, &arg)

	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(w, 30001, err.Error())
	} else {
		util.RespOk(w, com)
	}
}

func Addfriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)

	err := contactService.AddFriend(arg.UserId, arg.DstId)
	if err != nil {
		util.RespFail(w, 30003, err.Error())
	} else {
		util.RespOk(w, nil)
	}
}
