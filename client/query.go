package client

import (
	"context"
	"query_api/grpc_gen/query"
)

type QueryClient interface {
	GetQueryByID(ctx context.Context, id int32) (*query.Query, error)
	CreateQuery(ctx context.Context, id int32, data *query.Query) error
}

var (
	queryClient QueryClient
)

func GetQueryClient() QueryClient {
	return queryClient
}

func InitQueryClient(client QueryClient) {
	queryClient = client
}
