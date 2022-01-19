package config

import (
	"query_api/conf"
	"query_api/pkg/helper"

	"gopkg.in/ini.v1"
)

var RpcConfig *rpcConfig

type rpcConfig struct {
	Address                string
	Port                   int
	ClientPoolConnsSizeCap int
	DialTimeout            int
	KeepAlive              int
	KeepAliveTimeout       int
	source                 *ini.File
}

func (s *rpcConfig) Load(path string) *rpcConfig {
	//判断配置文件是否存在
	exists, err := helper.PathExists(path)
	if !exists {
		return s
	}
	s.source, err = ini.Load(path)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *rpcConfig) InitRpc() *rpcConfig {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}
	section := s.source.Section("RpcServer")
	s.Address = section.Key("address").MustString("127.0.0.1")
	s.Port = section.Key("port").MustInt(8080)
	s.ClientPoolConnsSizeCap = section.Key("clientPoolConnsSizeCap").MustInt(conf.DefaultClientPoolConnsSizeCap)
	s.DialTimeout = section.Key("dialTimeout").MustInt(int(conf.DefaultDialTimeout))
	s.KeepAlive = section.Key("keepAlive").MustInt(int(conf.DefaultKeepAlive))
	s.KeepAliveTimeout = section.Key("keepAliveTimeout").MustInt(int(conf.DefaultKeepAliveTimeout))
	return s
}
