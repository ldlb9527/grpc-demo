package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-demo/users"
	"log"
	"time"
)

func main() {

	//建立链接
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("did not connect", err)
	}

	defer conn.Close()

	userClient := users.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3) //设定请求超时时间 3s

	defer cancel()

	//UserIndex 请求
	userIndexResponse, err := userClient.UserIndex(ctx, &users.UserIndexRequest{
		Page: 1,
		PageSize: 12,
	})

	if err != nil {
		log.Printf("user index could not greet: %v", err)
	}

	if 0 == userIndexResponse.Err {
		log.Printf("user index success: %s", userIndexResponse.Msg)
		// 包含 UserEntity 的数组列表
		userEntityList := userIndexResponse.Data
		for _, row := range userEntityList {
			fmt.Println(row.Name, row.Age)
		}
	} else {
		log.Printf("user index error: %d", userIndexResponse.Err)
	}

	// UserView 请求
	userViewResponse, err := userClient.UserView(ctx, &users.UserViewRequest{Uid: 1})

	if err != nil {
		log.Printf("user view could not greet: %v", err)
	}

	if 0 == userViewResponse.Err {
		log.Printf("user view success: %s", userViewResponse.Msg)
		userEntity := userViewResponse.Data
		fmt.Println(userEntity.Name, userEntity.Age)
	} else {
		log.Printf("user view error: %d", userViewResponse.Err)
	}

	// UserPost 请求
	userPostReponse, err := userClient.UserPost(ctx, &users.UserPostRequest{Name: "big_cat", Password: "123456", Age: 29})

	if err != nil {
		log.Printf("user post could not greet: %v", err)
	}

	if 0 == userPostReponse.Err {
		log.Printf("user post success: %s", userPostReponse.Msg)
	} else {
		log.Printf("user post error: %d", userPostReponse.Err)
	}

	// UserDelete 请求

	userDeleteReponse, err := userClient.UserDelete(ctx, &users.UserDeleteRequest{Uid: 1})

	if err != nil {
		log.Printf("user delete could not greet: %v", err)
	}

	if 0 == userDeleteReponse.Err {
		log.Printf("user delete success: %s", userDeleteReponse.Msg)
	} else {
		log.Printf("user delete error: %d", userDeleteReponse.Err)
	}

}
