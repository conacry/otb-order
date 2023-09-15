package initApp

import (
	httpServer "github.com/conacry/go-platform/pkg/http/server"
	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	httpErrorResolver "github.com/conacry/go-platform/pkg/http/server/resolver"
	httpResponse "github.com/conacry/go-platform/pkg/http/server/response"
)

type HttpServerInit struct {
	dc         *DependencyContainer
	httpConfig *httpServerModel.Config
	server     *httpServer.HttpServer
}

func NewHttpServerInit(dc *DependencyContainer) *HttpServerInit {
	return &HttpServerInit{
		dc: dc,
	}
}

func (i *HttpServerInit) Init() error {
	err := i.initHttpConfig()
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	defaultErrResolver := httpErrorResolver.NewDefaultErrorResolver()
	errResWriter, err := httpResponse.NewErrorWriter(i.dc.Logger, defaultErrResolver)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	resWriter, err := httpResponse.NewWriter(i.dc.Logger)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	server, err := httpServer.NewBuilder().
		Logger(i.dc.Logger).
		Config(i.httpConfig).
		ResponseWriter(resWriter).
		ErrorResponseWriter(errResWriter).
		Build()
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	i.dc.HttpServer = server

	return nil
}

func (i *HttpServerInit) initHttpConfig() error {
	port := i.dc.AppConfig.HttpConfig.Port
	if port == 0 {
		return ErrIllegalHttpPort
	}

	httpConfig := httpServerModel.NewDefaultConfig(port)
	httpConfig.ReadTimeout = i.dc.AppConfig.HttpConfig.HttpTimeout.Read
	httpConfig.WriteTimeout = i.dc.AppConfig.HttpConfig.HttpTimeout.Write
	httpConfig.IdleTimeOut = i.dc.AppConfig.HttpConfig.HttpTimeout.Idle

	i.httpConfig = httpConfig
	return nil
}

func (i *HttpServerInit) Start() error {
	err := i.dc.HttpServer.Start(i.dc.Ctx)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	return nil
}

func (i *HttpServerInit) Stop() error {
	err := i.dc.HttpServer.Stop(i.dc.Ctx)
	if err != nil {
		i.dc.LogError(err)
		return err
	}

	return nil
}
