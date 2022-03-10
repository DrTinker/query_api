package config

import (
	"errors"
	"query_api/conf"
	"query_api/models"
	"query_api/pkg/helper"

	"gopkg.in/ini.v1"
)

type ConfigClientImpl struct {
	Http   *models.HttpConfig
	Rpc    *models.RpcConfig
	source *ini.File
}

func NewConfigClientImpl() *ConfigClientImpl {
	return &ConfigClientImpl{}
}

func (c *ConfigClientImpl) Load(path string) error {
	var err error
	//判断配置文件是否存在
	exists, _ := helper.PathExists(path)
	if !exists {
		return errors.New("config path not exists")
	}
	c.source, err = ini.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigClientImpl) GetHttpConfig() (*models.HttpConfig, error) {
	//判断配置是否加载成功
	if c.source == nil {
		return nil, errors.New("empty http config")
	}
	c.Http = &models.HttpConfig{}
	c.Http.Address = c.source.Section("HttpServer").Key("address").MustString("127.0.0.1")
	c.Http.Port = c.source.Section("HttpServer").Key("port").MustInt(8080)
	return c.Http, nil
}

func (c *ConfigClientImpl) GetRPCConfig() (*models.RpcConfig, error) {
	//判断配置是否加载成功
	if c.source == nil {
		return nil, errors.New("empty rpc config")
	}
	section := c.source.Section("RpcServer")
	c.Rpc = &models.RpcConfig{}
	c.Rpc.Address = section.Key("address").MustString("127.0.0.1")
	c.Rpc.Port = section.Key("port").MustInt(50052)
	c.Rpc.ClientPoolConnsSizeCap = section.Key("clientPoolConnsSizeCap").MustInt(conf.DefaultClientPoolConnsSizeCap)
	c.Rpc.DialTimeout = section.Key("dialTimeout").MustInt(int(conf.DefaultDialTimeout))
	c.Rpc.KeepAlive = section.Key("keepAlive").MustInt(int(conf.DefaultKeepAlive))
	c.Rpc.KeepAliveTimeout = section.Key("keepAliveTimeout").MustInt(int(conf.DefaultKeepAliveTimeout))
	return c.Rpc, nil
}
