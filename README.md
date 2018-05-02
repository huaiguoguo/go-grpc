# 主要平时学习 go 语言的 Grpc

### simple 简单例子
> client 客户端

> server 服务端

> simple.proto

## 安装步骤

### 一、安装 protobuf (protoc)
#### 下载地址：https://github.com/google/protobuf
#### 主要用于编写好的 protobuf 文件，生成于 go 文件
#### 命令(在protobuf文件的目录)：
> protoc --go_out=plugins=grpc:. test.proto

### 二、安装 protoc-gen-go
> go get github.com/golang/protobuf

> go get github.com/golang/protobuf/protoc-gen-go

### 三、安装 grpc-go (已经打包在 vendor 中了)

#### 有三个依赖和本身也被墙了，可以采用手动下载，Github 有。
#### 第一个：go get godoc.org/golang.org/x/net  被墙了,可以手动下载 https://github.com/golang/net
#### 第二个：go get godoc.org/golang.org/x/text 被墙了,可以手动下载 https://github.com/golang/text
#### 第三个：go get google.golang.org/genproto  被墙了,可以手动下载 https://github.com/google/go-genproto
#### 第四个：go get google.golang.org/grpc      被墙了,可以手动下载 https://github.com/grpc/grpc-go
#### 下载后将目录改 GO 的目录规则

![Image text](https://github.com/laixhe/go_grpc/blob/master/grpc-go.png)

#### 然后没了

