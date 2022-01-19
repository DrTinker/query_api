package conf

import "errors"

// RPC连接相关错误
var ErrNotFoundClient = errors.New("not found grpc conn")
var ErrConnShutdown = errors.New("grpc conn shutdown")