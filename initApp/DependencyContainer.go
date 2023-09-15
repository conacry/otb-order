package initApp

import (
	"context"
	"fmt"
	httpServer "github.com/conacry/go-platform/pkg/http/server"
	initapp "github.com/conacry/go-platform/pkg/init"
	log "github.com/conacry/go-platform/pkg/logger"
	"github.com/conacry/go-platform/pkg/postgresql"
)

type DependencyContainer struct {
	Ctx        context.Context
	AppConfig  *AppConfig
	Logger     log.Logger
	PostgreSQL *postgresql.PostgreSQL
	HttpServer *httpServer.HttpServer
}

func NewDependencyContainer(ctx context.Context, appConfig *AppConfig) (*DependencyContainer, error) {
	if ctx == nil {
		return nil, ErrAppContextIsRequired
	}
	if appConfig == nil {
		return nil, ErrAppConfigIsRequired
	}

	dc := &DependencyContainer{
		Ctx:       ctx,
		AppConfig: appConfig,
	}

	return dc, nil
}

func (dc *DependencyContainer) LogError(err error) {
	if dc.Logger == nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		dc.Logger.LogError(dc.Ctx, err)
	}
}

func (dc *DependencyContainer) LogInfo(msg string) {
	if dc.Logger == nil {
		fmt.Println(msg)
	} else {
		dc.Logger.LogInfo(context.Background(), msg)
	}
}

func GetDependencyInitChains(dc *DependencyContainer) []initapp.DependencyService {
	return []initapp.DependencyService{
		NewLoggerInit(dc),
		NewPostgresqlInit(dc),
		NewHttpServerInit(dc),
	}
}
