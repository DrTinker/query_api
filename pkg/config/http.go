package config

import (
	
	"query_api/pkg/helper"

	"gopkg.in/ini.v1"
)

var HttpConfig *httpConfig

type httpConfig struct {
	Address string
	Port int
	source *ini.File
}

func (s *httpConfig) Load(path string) *httpConfig {
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

func (s *httpConfig)InitHttp() *httpConfig {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}
	s.Address = s.source.Section("HttpServer").Key("address").MustString("127.0.0.1")
	s.Port = s.source.Section("HttpServer").Key("port").MustInt(8080)
	return s
}