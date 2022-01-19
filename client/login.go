package rpc_client

import (
	"context"
	"query_api/conf"
	"query_api/infrastructure/rpc/user"
	"query_api/models"

	"google.golang.org/grpc"
)

// query_rpc连接
var RpcConn *grpc.ClientConn

type UseruserLogin interface {
	UserLogin(ctx context.Context, params models.Login) (bool, user.User, error)
}

type userLoginClient struct{
	
}

var UserLoginClient userLoginClient

func (u userLoginClient) UserLogin(ctx context.Context, params models.Login) (bool, *user.User, error) {
	id := params.User_id
	pwd := params.User_pwd
	req := user.UserLoginReq{
		UserId: id,
		UserPwd: pwd,
	}
	c := user.NewUserServiceClient(RpcConn)
	resp, err := c.UserLogin(ctx, &req)
	if err != nil {
		return false, nil, err
	}
	if (resp.Resp.Code == conf.RPC_SUCCESS_CODE) {
		return true, resp.UserInfo, nil
	}
	return false, nil, nil
}