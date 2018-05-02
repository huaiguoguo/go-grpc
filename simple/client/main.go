package main

import (
	"log"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go_grpc/simple"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

)

const (
	// Address gRPC 服务地址
	Address = "127.0.0.1:50051"
)

func main() {

	// 连接服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	//关闭连接
	defer conn.Close()

	//初始化客户端
	client := pb.NewSimpleClient(conn)

	//初始化 TestRequest 请求结构
	reqBody := &pb.SimpleRequest{Typeid:10}

	//调用客户端 SayTest 方法
	r, err := client.SayTest(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range r.UserList {
		log.Println(v)
	}


}
