package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/simple"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	// 获取 GET 的参数
	userid := r.FormValue("userid")
	id, err := strconv.ParseInt(userid, 10, 0)
	if err != nil {
		w.Write([]byte("userid The parameters must be integers"))
		return
	}

	// 调用 Grpc 的远程接口
	data, err := UClient.GetUser(context.Background(), &pb.UserRequest{Userid: id})
	if err != nil {
		w.Write([]byte("Grpc: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data)
	w.Write(js)

}

func GetUserList(w http.ResponseWriter, r *http.Request) {

	// 调用 Grpc 的远程接口
	data, err := UClient.GetUserList(context.Background(), &pb.UserRequest{})
	if err != nil {
		w.Write([]byte("Grpc: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data.UserList)
	w.Write(js)

}

func SetUser(w http.ResponseWriter, r *http.Request) {

	// 获取 GET 的参数
	userid := r.FormValue("userid")
	id, err := strconv.ParseInt(userid, 10, 0)
	if err != nil {
		w.Write([]byte("userid The parameters must be integers"))
		return
	}

	name := r.FormValue("username")
	if name == "" {
		w.Write([]byte("username Not null"))
		return
	}

	// 调用 Grpc 的远程接口
	data, err := UClient.SetUser(context.Background(), &pb.UserModel{Userid:id, Username:name})
	if err != nil {
		w.Write([]byte("Grpc: " + err.Error()))
		return
	}


	// json 格式化
	js, _ := json.Marshal(data)
	w.Write(js)

}
