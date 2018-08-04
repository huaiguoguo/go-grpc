package main

import (
	"net/http"
	"context"
	"encoding/json"

	pb "github.com/laixhe/go-grpc/simple"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request){

	// 调用 Grpc 的远程接口
	data, err := UIClient.GetUserInfo(context.Background(), &pb.UserInfoRequest{Userid: 10000})
	if err != nil {
		w.Write([]byte("Grpc: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data)
	w.Write(js)

}
func SetUserInfo(w http.ResponseWriter, r *http.Request){

	// 调用 Grpc 的远程接口
	data, err := UIClient.SetUserInfo(context.Background(), &pb.UserInfoModel{Userid: 20000, Username:"SetUserInfo"})
	if err != nil {
		w.Write([]byte("Grpc: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data)
	w.Write(js)
}
