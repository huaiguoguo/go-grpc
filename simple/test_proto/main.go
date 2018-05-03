package main

import(
	"github.com/golang/protobuf/proto"
	"github.com/laixhe/go_grpc/simple"
	"fmt"
)

// 进行编解码 protobuf
func main(){

	reply := new(simple.SimpleReply)
	var getdata []*simple.GetData

	for i:=0; i<=3; i++ {

		getdata = append(getdata, &simple.GetData{Id:int64(i),Name:"laiki"})

	}
	reply.UserList = getdata

	fmt.Println(reply)

	// 进行编码 protobuf
	data, err := proto.Marshal(reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	ry := new(simple.SimpleReply)

	// 进行解码 protobuf
	proto.Unmarshal(data, ry)

	fmt.Println(ry)
}
