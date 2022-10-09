package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type result struct { //定义返回数据格式
	Code int
	Msg  string
	Data []string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/homepage", homepage)
	err := http.ListenAndServe("127.0.0.1:9300", nil)
	if err != nil {
		return
	} //设置监听的端口
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/login.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homepage")
	if r.Method == "GET" {
		write, _ := template.ParseFiles("view/homepage.html")
		err := write.Execute(w, nil)
		if err != nil {
			return
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	if r.Method == "POST" || r.Method == "GET" {
		fmt.Println(r.Form)
		email, _ := r.Form["emailLogin"]
		emailLogin := email[0]
		password, _ := r.Form["passwordLogin"]
		passwordLogin := password[0]
		if emailLogin == "12345678@qq.com" && passwordLogin == "123" {
			fmt.Println("go")
			http.Redirect(w, r, "/homepage", http.StatusTemporaryRedirect)
			arr := &result{
				200,
				"登陆成功",
				[]string{},
			}
			b, jsonErr := json.Marshal(arr) //json化结果集
			if jsonErr != nil {
				fmt.Println("encoding faild")
			} else {
				io.WriteString(w, string(b)) //返回结果
				fmt.Println(string(b))
			}
		} else {
			_, err := w.Write([]byte("<script>alert('wrong')</script>"))
			if err != nil {
				return
			}
		}
	}
}
