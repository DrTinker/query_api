package client

import (
	"context"
	"query_api/conf"
	"query_api/infrastructure/rpc/user"
	"query_api/models"
)

type UserOption interface {
	GetUserByUserID(ctx context.Context, id int32) (*user.User, error)
}

type userOptionClient struct{
	
}

var UserOptionClient userOptionClient

func (u userOptionClient) GetUserByUserID(ctx context.Context, id int32) (*user.User, error) {
	req := user.GetUserByUserIDReq{
		UserId: id,
	}
	c := user.NewUserServiceClient(models.RpcConn)
	resp, err := c.GetUserByUserID(ctx, &req)
	if err != nil {
		return nil, err
	}
	if (resp.Resp.Code == conf.RPC_SUCCESS_CODE) {
		return resp.UserInfo, nil
	}
	return nil, conf.ErrRpcEmptyResp
}