package initApp

import (
	"errors"
	"fmt"
	"github.com/conacry/go-platform/pkg/mongo"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
)

type MongoDBInit struct {
	dc      *DependencyContainer
	mongoDB *mongo.MongoDB
}

func NewMongoDBInit(dc *DependencyContainer) *MongoDBInit {
	return &MongoDBInit{
		dc: dc,
	}
}

func (i *MongoDBInit) Init() error {
	mongoConfig := &mongoModel.Config{
		URL:      i.dc.AppConfig.MongoConfig.URL,
		Database: i.dc.AppConfig.MongoConfig.Database,
	}

	mongoDB, err := mongo.NewBuilder().
		Logger(i.dc.Logger).
		Config(mongoConfig).
		Build()
	if err != nil {
		errMsg := fmt.Sprintf("MongoDB initialization failed. Cause: %q", err.Error())
		wrappedErr := errors.New(errMsg)
		i.dc.LogError(wrappedErr)
		return err
	}

	i.mongoDB = mongoDB
	i.dc.MongoDB = mongoDB

	return nil
}

func (i *MongoDBInit) Start() error {
	err := i.mongoDB.Start(i.dc.Ctx)
	if err != nil {
		errMsg := fmt.Sprintf("MongoDB starting failed. Cause: %q", err.Error())
		wrappedErr := errors.New(errMsg)
		i.dc.LogError(wrappedErr)
		return err
	}

	return nil
}

func (i *MongoDBInit) Stop() error {
	err := i.mongoDB.Stop(i.dc.Ctx)
	if err != nil {
		errMsg := fmt.Sprintf("MongoDB stopping failed. Cause: %q", err.Error())
		wrappedErr := errors.New(errMsg)
		i.dc.LogError(wrappedErr)
		return err
	}

	return nil
}
