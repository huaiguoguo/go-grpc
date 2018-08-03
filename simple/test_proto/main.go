package main

import(
	"github.com/golang/protobuf/proto"
	"github.com/laixhe/go-grpc/simple"
	"fmt"
)

// 进行编解码 protobuf
func main(){

	reply := new(simple.UserListResponse)
	var getdata []*simple.UserModel

	for i:=0; i<=3; i++ {

		getdata = append(getdata, &simple.UserModel{Userid:int64(i),Username:"laiki"})

	}
	reply.UserList = getdata

	fmt.Println(reply)

	// 进行编码 protobuf
	data, err := proto.Marshal(reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	ry := new(simple.UserListResponse)

	// 进行解码 protobuf
	proto.Unmarshal(data, ry)

	fmt.Println(ry)
}
