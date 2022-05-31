package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-demo/users"
	"log"
	"net"
)


//UserServer  实现User服务的业务对象
type UserServer struct {

}

//UserIndex 实现了User 服务接口的所有方法
func (u *UserServer) UserIndex(ctx context.Context, in *users.UserIndexRequest) (*users.UserIndexResponse, error) {
	log.Printf("receive users index request:page %d page_size %d", in.Page, in.PageSize)
	return &users.UserIndexResponse{
		Err: 0,
		Msg: "success",
		Data: []*users.UserEntity{
			{Name: "aaaa", Age: 28},
			{Name: "bbbb", Age: 1},
		},
	}, nil
}

//UserView 获取详情
func (u *UserServer) UserView(ctx context.Context, in *users.UserViewRequest) (*users.UserViewResponse, error) {
	log.Printf("receive users uid request:uid %d", in.Uid)
	return &users.UserViewResponse{
		Err: 0,
		Msg: "success",
		Data: &users.UserEntity{
			Name: "aaaa", Age: 28,
		},
	}, nil
}

//UserPost 提交数据
func (u *UserServer) UserPost(ctx context.Context, in *users.UserPostRequest) (*users.UserPostResponse, error) {
	log.Printf("receive users uid request:name %s password:%s,age:%d", in.Name, in.Password, in.Age)
	return &users.UserPostResponse{
		Err: 0,
		Msg: "success",
	}, nil
}

//UserDelete 删除数据
func (u *UserServer) UserDelete(ctx context.Context, in *users.UserDeleteRequest) (*users.UserDeleteResponse, error) {
	log.Printf("receive users uid request:uid %d", in.Uid)
	return &users.UserDeleteResponse{
		Err: 0,
		Msg: "success",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("failed to listen", err)
	}

	//创建rpc服务
	grpcServer := grpc.NewServer()

	//为User服务注册业务实现 将User服务绑定到RPC服务器上
	users.RegisterUserServer(grpcServer, &UserServer{})

	//注册反射服务， 这个服务是CLI使用的， 跟服务本身没有关系
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("faild to server,", err)
	}
}

