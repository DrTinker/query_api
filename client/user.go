package client

import (
	"context"
	"query_api/grpc_gen/user"
)

type UserClient interface {
	GetUserByUserID(ctx context.Context, id int32) (*user.User, error)
	CreateUser(ctx context.Context, user *user.User) error
}

var (
	userClient UserClient
)

func GetUserClient() UserClient {
	return userClient
}

func InitUserClient(client UserClient) {
	userClient = client
}
