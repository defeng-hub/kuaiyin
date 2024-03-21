package service

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "kuaiyin-user/api/user/v1"
	"kuaiyin-user/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServiceServer

	uc *biz.GreeterUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.GreeterUsecase) *UserService {
	return &UserService{uc: uc}
}

func (u *UserService) Login(c context.Context, in *v1.KuaiyinUserLoginRequest) (*v1.KuaiyinUserLoginResponse, error) {
	u.uc.CreateGreeter(c, &biz.Greeter{Hello: in.Username})
	fmt.Println("aaaa")
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (u *UserService) Register(context.Context, *v1.KuaiyinUserRegisterRequest) (*v1.KuaiyinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (u *UserService) UserInfo(context.Context, *v1.KuaiyinUserRequest) (*v1.KuaiyinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
