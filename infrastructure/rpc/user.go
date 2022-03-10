package rpc

import (
	"context"
	"fmt"
	"query_api/conf"
	"query_api/grpc_gen/user"

	"query_api/models"
	"query_api/pkg/helper"

	"github.com/pkg/errors"
)

type UserClientImpl struct {
	server user.UserServiceClient
}

func NewUserServiceClientImpl() *UserClientImpl {
	return &UserClientImpl{
		server: user.NewUserServiceClient(models.RpcConn),
	}
}

func (u *UserClientImpl) CreateUser(ctx context.Context, data *user.User) error {
	id := helper.GenUid(data.UserName, int(data.Phone))
	data.UserId = int32(id)
	req := &user.CreateUserReq{
		UserInfo: data,
	}
	resp, err := u.server.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.Resp.Code != conf.RPC_SUCCESS_CODE {
		return errors.New(fmt.Sprintf("UserClient error, code: %+v", resp.Resp.Code))
	}
	return nil
}

func (u *UserClientImpl) GetUserByUserID(ctx context.Context, id int32) (*user.User, error) {
	req := user.GetUserByUserIDReq{
		UserId: id,
	}
	resp, err := u.server.GetUserByUserID(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.Code == conf.RPC_SUCCESS_CODE {
		return resp.GetUserInfo(), nil
	}
	return nil, conf.ErrRpcEmptyResp
}
