package main

import (
	pb "xuexi_grpc/xuexitest" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}

	//关闭连接
	defer conn.Close()

	//初始化客户端
	client := pb.NewXuexitestClient(conn)

	//初始化 TestRequest 请求结构
	reqBody := new(pb.TestRequest)
	reqBody.Typeid = 10

	//调用客户端 SayTest 方法
	r, err := client.SayTest(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Println(r.Getdataarr)
}
