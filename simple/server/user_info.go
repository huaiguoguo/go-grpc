package main

import(
	"context"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/simple"
)

// 定义 UserInfo 并实现约定的接口
type UserInfo struct {}

func (this *UserInfo) GetUserInfo (ctx context.Context, ut *pb.UserInfoRequest) (*pb.UserInfoModel, error) {

	// 待返回数据结构
	resp := new(pb.UserInfoModel)

	return resp, nil
}

func (this *UserInfo) SetUserInfo (ctx context.Context, ut *pb.UserInfoModel) (*pb.UserInfoResponse, error) {

	// 待返回数据结构
	resp := new(pb.UserInfoResponse)

	return resp, nil
}
