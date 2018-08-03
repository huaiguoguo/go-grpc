package main

import (
	"log"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/simple"

	"google.golang.org/grpc"
	"net/http"
)

const (
	// Address gRPC 服务地址
	Address = "127.0.0.1:50051"
)

var UClient pb.UserClient

var UIClient pb.UserInfoClient

func init() {

	// 连接服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭连接
	defer conn.Close()

	// 初始化 User 客户端
	UClient = pb.NewUserClient(conn)

	// 初始化 UserInfo 客户端
	UIClient = pb.NewUserInfoClient(conn)

}

// 启动 http 服务
func main(){

	http.HandleFunc("/user/get", GetUser)
	http.HandleFunc("/user/get/list", GetUserList)
	http.HandleFunc("/user/set", SetUser)
	http.HandleFunc("/userinfo/get", GetUserInfo)
	http.HandleFunc("/userinfo/set", SetUserInfo)

	http.ListenAndServe(":8080", nil)
}
