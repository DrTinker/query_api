package conf

import "errors"

// RPC
// RPC连接相关错误
var ErrNotFoundClient = errors.New("not found grpc conn")
var ErrConnShutdown = errors.New("grpc conn shutdown")
// RPC空返回值错误
var ErrRpcEmptyResp = errors.New("rpc empty resp")

// 业务逻辑
// 登录
var ErrPwdError = errors.New("password error")
// jwt
var ErrJWTTokenEmpty = errors.New("token could not be empty")
var ErrJWTTokenInvaild = errors.New("token is invaild")
var ErrJWTTokenTime = errors.New("token is either expired or not active yet")
var ErrJWTTokenHandle = errors.New("Couldn't handle this token:")