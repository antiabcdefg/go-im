package main

import (
	"go-im/controller"
	"html/template"
	"io"
	"log"
	"net/http"
)

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tpname := v.Name()
		http.HandleFunc(tpname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tpname, nil)
		})
	}
}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"Hello, World!")
	})
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)
	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)
	http.HandleFunc("/contact/createcommunity", controller.CreateCommunity)
	http.HandleFunc("/contact/addfriend", controller.Addfriend)
	http.HandleFunc("/chat", controller.Chat)
	http.HandleFunc("/attach/upload", controller.Upload)
	//提供静态资源
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))

	RegisterView()
	//启动服务器
	http.ListenAndServe(":8080", nil)
}
