package main

import (
	"context"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/simple"
)

// 定义 User 并实现约定的接口
type User struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

func (this *User) GetUser(ctx context.Context, ut *pb.UserRequest) (*pb.UserModel, error) {

	// 待返回数据结构
	resp := new(pb.UserModel)

	data := userList.Get(ut.Userid)
	if data != nil {
		resp.Userid = data.UserId
		resp.Username = data.UserName
	}

	return resp, nil
}

func (this *User) GetUserList(ctx context.Context, ut *pb.UserRequest) (*pb.UserListResponse, error) {

	// 获取 user 所有数据
	data := userList.List()
	list := make([]*pb.UserModel, 0, len(data))

	for _, v := range data {

		list = append(list, &pb.UserModel{
			Userid:   v.UserId,
			Username: v.UserName,
		})

	}

	// 待返回数据结构
	resp := new(pb.UserListResponse)
	resp.UserList = list

	return resp, nil
}

func (this *User) SetUser(ctx context.Context, ut *pb.UserModel) (*pb.UserResponse, error) {

	userList.Set(ut.Userid, &User{UserId: ut.Userid, UserName: ut.Username})

	// 待返回数据结构
	resp := new(pb.UserResponse)
	resp.Code = 1
	resp.Msg = "成功"

	return resp, nil
}
