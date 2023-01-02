package main

import (
	"encoding/json"
	"fmt"
)

type ApiResult struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"string"`
	Data map[string]interface{} `json:"data"`
}

func main() {
	m := make(map[string]interface{})
	m["age"] = 200
	m["isMan"] = true
	m["name"] = "韩玉坤"
	apiResult := ApiResult{Code: 200, Msg: "talk is cheap,show me the code", Data: m}
	s := marshal2JsonString(apiResult)
	fmt.Println(s)
	ar := unmarshal2Struct(s)
	fmt.Println(ar)
}

/*
从struct转换成string字符串， 从string字符串转换成struct对象
*/
func marshal2JsonString(apiResult ApiResult) string {
	apiResultStr, err := json.Marshal(&apiResult)
	if err != nil {
		println(err)
	}
	return string(apiResultStr)
}

/*
从string字符串转换为struct对象
*/
func unmarshal2Struct(resultString string) ApiResult {
	var apiResult ApiResult
	err := json.Unmarshal([]byte(resultString), &apiResult)
	if err != nil {
		println(err)
	}
	return apiResult
}
