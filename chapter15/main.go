package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/pprof"
)

type ApiResult struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"message"`
	Data map[string]interface{} `json:"data"`
}

// 测试HTTP服务器
func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/debug/pprof/", pprof.Index)
	sm.HandleFunc("/debug/profile", pprof.Profile)
	sm.HandleFunc("/debug/symbol", pprof.Symbol)
	sm.HandleFunc("/debug/trace", pprof.Trace)
	sm.HandleFunc("/login", Login)
	if err := http.ListenAndServe("localhost:80", sm); err != nil {
		log.Fatalf("I am sorry, Start server Failed!: %s", err.Error())
	}
}

/*
登录接口，传入相应的用户名和密码，登录成功返回对应的成功信息，并返回accesstoken， 登录失败，提示失败信息
*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/json")
	type RequestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var rb RequestBody
	json.NewDecoder(r.Body).Decode(&rb)
	var apiResult ApiResult
	if rb.Username == "kate" && rb.Password == "123456" {
		apiResult = ApiResult{
			Code: 200,
			Msg:  "登录成功",
			Data: nil,
		}
	} else {
		apiResult = ApiResult{
			Code: 200,
			Msg:  "登录失败",
			Data: nil,
		}
	}
	b, err := json.Marshal(apiResult)
	if err == nil {
		w.Write(b)
	}
}
