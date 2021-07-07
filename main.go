package main

import (
    "io"
    "log"
    "net/http"
	"encoding/json"
)

type JsonResult struct{
	Code int `json:"code"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request)  {
	ret:=new(JsonResult)
	ret.Code = 0
	ret.Message = "success"
	msg, _ := json.Marshal(ret)
	io.WriteString(w, string(msg))
}

func main() {
    http.HandleFunc("/api/account/register", handler)    
    http.HandleFunc("/api/account/login", handler)    
    err := http.ListenAndServe(":4000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
