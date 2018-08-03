package main

import(
	"context"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/simple"

)

// 定义 User 并实现约定的接口
type User struct{

}

func (this *User) GetUser(ctx context.Context, ut *pb.UserRequest) (*pb.UserModel, error) {

	// 待返回数据结构
	resp := new(pb.UserModel)

	return resp, nil
}

func (this *User) GetUserList(ctx context.Context, ut *pb.UserRequest) (*pb.UserListResponse, error) {

	// 待返回数据结构
	resp := new(pb.UserListResponse)

	return resp, nil
}

func (this *User) SetUser(ctx context.Context, ut *pb.UserModel) (*pb.UserResponse, error) {

	// 待返回数据结构
	resp := new(pb.UserResponse)

	return resp, nil
}
