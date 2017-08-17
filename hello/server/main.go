package main

import (
	"net"

	pb "xuexi_grpc/xuexitest" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"strconv"
)

const (
	// Address gRPC服务地址
	Address = "0.0.0.0:50021"
)

// 定义helloService并实现约定的接口
type testService struct{}

func (t *testService) SayTest(ctx context.Context, in *pb.TestRequest) (*pb.TestReply, error) {
	//待返回数据
	resp := new(pb.TestReply)

	//准备数据
	var i int64 = 0
	for ; i < in.Typeid; i++ {
		//追加到 待返回数据
		resp.Getdataarr = append(resp.Getdataarr, &pb.TestReply_GetData{Id: i + 1, Name: "laixhe" + strconv.FormatInt(i+1, 10)})
	}

	grpclog.Println(in.Typeid)

	return resp, nil
}

func main() {
	//监听地址和端口
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("监听端口失败: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterXuexitestServer(s, &testService{})

	grpclog.Println("服务：" + Address)

	//启动服务
	s.Serve(listen)
}
