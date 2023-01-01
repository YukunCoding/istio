package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"
)

type JsonResult struct {
	Code int               `json:"code"`
	Msg  string            `json:"message"`
	Data map[string]string `json:"data"`
}

// 测试HTTP服务器
func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/debug/pprof/", pprof.Index)
	sm.HandleFunc("/debug/profile", pprof.Profile)
	sm.HandleFunc("/debug/symbol", pprof.Symbol)
	sm.HandleFunc("/debug/trace", pprof.Trace)

	sm.HandleFunc("/get", Get)
	sm.HandleFunc("/post", Post)
	sm.HandleFunc("/postjson", PostJson)
	if err := http.ListenAndServe("localhost:80", sm); err != nil {
		log.Fatalf("I am sorry, Start server Failed!: %s", err.Error())
	}
}

// 读取Get请求中的参数 r.FormValue("key")
func Get(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	w.Header().Set("content-type", "text/json")
	m := make(map[string]string)
	m["username"] = username
	m["password"] = password
	b, err := json.Marshal(JsonResult{Code: 200, Msg: "kate", Data: m})
	if err == nil {
		w.Write(b)
	}
}

// 读取Post请求body中获取参数 r.PostFormValue()
func Post(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	w.Header().Set("content-type", "text/json")
	m := make(map[string]string)
	m["username"] = username
	m["password"] = password
	b, err := json.Marshal(JsonResult{Code: 200, Msg: "kate", Data: m})
	if err == nil {
		w.Write(b)
	}
}

// 读取Post请求body中获取参数 定义结构体
// json.NewDecoder(r.Body).Decode(&rb)
func PostJson(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var rb RequestBody
	json.NewDecoder(r.Body).Decode(&rb)
	w.Header().Set("Content-Type", "text/json")
	m := make(map[string]string)
	m["username"] = rb.Username
	m["password"] = rb.Password
	b, err := json.Marshal(JsonResult{Code: 200, Msg: "kate", Data: m})
	if err == nil {
		w.Write(b)
	}
}
