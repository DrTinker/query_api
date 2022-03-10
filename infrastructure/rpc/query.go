package rpc

import (
	"context"
	"fmt"
	"query_api/conf"
	"query_api/grpc_gen/query"
	"query_api/models"
	"query_api/pkg/helper"

	"github.com/pkg/errors"
)

type QueryClientImpl struct {
	server query.QueryServiceClient
}

func NewQueryClientImpl() *QueryClientImpl {
	return &QueryClientImpl{
		server: query.NewQueryServiceClient(models.RpcConn),
	}
}

func (q *QueryClientImpl) CreateQuery(ctx context.Context, id int32, data *query.Query) error {
	qid := helper.GenQid()
	data.QueryId = int32(qid)
	req := &query.CreateQueryReq{
		Query:  data,
		UserId: id,
	}
	resp, err := q.server.CreateQuery(ctx, req)
	if err != nil {
		return err
	}
	if resp.Resp.Code != conf.RPC_SUCCESS_CODE {
		return errors.New(fmt.Sprintf("QueryClient error, code: %+v", resp.Resp.Code))
	}
	return nil
}

func (q *QueryClientImpl) GetQueryByID(ctx context.Context, id int32) (*query.Query, error) {
	req := query.GetQueryByIDReq{
		QueryId: id,
	}
	resp, err := q.server.GetQueryByID(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.Code == conf.RPC_SUCCESS_CODE {
		return resp.GetQuery(), nil
	}
	return nil, conf.ErrRpcEmptyResp
}
