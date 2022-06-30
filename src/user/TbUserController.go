package user

import (
	"commons"
	"encoding/json"
	"net/http"
)

// UserHandler 所有user模块的handler
func UserHandler() {
	commons.Router.HandleFunc("/login", loginController)
}

//loginController 登录
func loginController(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	er := LoginService(username, password)
	b, _ := json.Marshal(er)                                         //把结构体转换为json数据
	w.Header().Set("Content-Type", "application/json;charset=utf-8") //设置响应内容为json
	w.Write(b)
}
