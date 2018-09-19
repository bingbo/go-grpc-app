package service

import (
	"go-grpc-app/dao"
	"go-grpc-app/pb"
	"context"
	"log"
	"fmt"
	"strconv"
)

type UserService struct {
}

var userDao *dao.UserDao

func init() {
	userDao = dao.NewUserDao()
}
func NewUserService() *UserService {
	return &UserService{}
}
func (this *UserService) GetUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("receive user request: %v", request)
	//user := userDao.GetUser(int(request.Id))
	//fmt.Println(user)
	return &pb.UserResponse{Users: []*pb.User{}}, nil
}

func (this *UserService) GetUsers(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("receive user request: %v", request)
	users, err := userDao.GetUsers()
	if err != nil {
		log.Fatalln(err)
		return nil, nil
	}
	var result []*pb.User
	for index, item := range users {
		log.Println(index, item)
		var user pb.User
		fmt.Println(index, string(item["id"].([]byte)), string(item["name"].([]byte)), string(item["password"].([]byte)), string(item["email"].([]byte)))
		id, _ := strconv.Atoi(string(item["id"].([]byte)))
		user.Id = int32(id)
		user.Name = string(item["name"].([]byte))
		user.Password = string(item["password"].([]byte))
		user.Email = string(item["email"].([]byte))
		fmt.Println(user)
		result = append(result, &user)
	}

	return &pb.UserResponse{Users: result}, err
}
