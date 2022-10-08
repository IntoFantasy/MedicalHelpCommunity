package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./LoginHtml/static/"))))
	http.HandleFunc("/login", login) //设置访问的路由
	err := http.ListenAndServe("127.0.0.1:9300", nil)
	if err != nil {
		return
	} //设置监听的端口
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, err := template.ParseFiles("LoginHtml/login.html")
		if err != nil {
			return
		}
		err = write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}
