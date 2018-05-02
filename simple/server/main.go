package main

import (
	"log"
	"net"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go_grpc/simple"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//由 protoc 生成 simple.pb.go 接口
// Server API for Simple service
//type SimpleServer interface {
//	// 定义 SayTest 方法
//	SayTest(context.Context, *SimpleRequest) (*SimpleReply, error)
//}

// 定义 Server 并实现约定的接口
type Server struct{}

func (this *Server) SayTest(ctx context.Context, rt *pb.SimpleRequest) (*pb.SimpleReply, error) {

	// 待返回数据结构
	resp := new(pb.SimpleReply)
	resp.UserList = make([]*pb.GetData, 0, rt.Typeid)

	// 准备数据
	var i int64 = 0
	for ; i < rt.Typeid; i++ {
		// 追加到 待返回数据
		resp.UserList = append(resp.UserList, &pb.GetData{Id: i + 1, Name: "laiki"})
	}

	return resp, nil
}

func main() {

	// 监听地址和端口
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	// 实例化 grpc Server
	serverGrpc := grpc.NewServer()

	// 注册 Simple service
	pb.RegisterSimpleServer(serverGrpc, &Server{})

	log.Println("开始监听端口 0.0.0.0:50051")

	// 启动服务
	err = serverGrpc.Serve(listen)
	if err != nil {
		log.Println("启动 Grpc 服务失败")
	}
}
