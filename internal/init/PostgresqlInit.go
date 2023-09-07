package initApp

import (
	"errors"
	"fmt"
	"github.com/conacry/go-platform/pkg/postgresql"
	"time"
)

type PostgresqlInit struct {
	dc         *DependencyContainer
	postgreSQL *postgresql.PostgreSQL
}

func NewPostgresqlInit(dc *DependencyContainer) *PostgresqlInit {
	return &PostgresqlInit{
		dc: dc,
	}
}

func (i *PostgresqlInit) Init() error {
	postgresqlConfig := &postgresql.Config{
		Url:             i.dc.AppConfig.PostrgesqlConfig.Url,
		MaxConnIdleTime: time.Minute * time.Duration(i.dc.AppConfig.PostrgesqlConfig.MaxConnIdleTime),
		MaxConns:        int32(i.dc.AppConfig.PostrgesqlConfig.MaxConns),
		MinConns:        int32(i.dc.AppConfig.PostrgesqlConfig.MinConns),
	}

	postgreSQL, err := postgresql.NewBuilder().
		Config(postgresqlConfig).
		Logger(i.dc.Logger).
		Build()
	if err != nil {
		errMsg := fmt.Sprintf("PostgreSQL initialization failed. Cause: %q", err.Error())
		wrappedErr := errors.New(errMsg)
		i.dc.LogError(wrappedErr)
		return err
	}

	i.postgreSQL = postgreSQL
	i.dc.PostgreSQL = postgreSQL

	return nil
}

func (i *PostgresqlInit) Start() error {
	err := i.postgreSQL.Start(i.dc.Ctx)
	if err != nil {
		errMsg := fmt.Sprintf("PostgreSQL starting failed. Cause: %q", err.Error())
		wrappedErr := errors.New(errMsg)
		i.dc.LogError(wrappedErr)
		return err
	}

	return nil
}

func (i *PostgresqlInit) Stop() error {
	err := i.postgreSQL.Stop(i.dc.Ctx)
	if err != nil {
		errMsg := fmt.Sprintf("PostgreSQL stopping failed. Cause: %q", err.Error())
		wrappedErr := errors.New(errMsg)
		i.dc.LogError(wrappedErr)
		return err
	}

	return nil
}
