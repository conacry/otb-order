package initApp

import (
	log "github.com/conacry/go-platform/pkg/logger"
)

type LoggerInit struct {
	dc *DependencyContainer
}

func NewLoggerInit(dc *DependencyContainer) *LoggerInit {
	return &LoggerInit{
		dc: dc,
	}
}

func (i *LoggerInit) Init() error {
	zapLogger := log.NewZapLogger(i.dc.AppConfig.Name)
	i.dc.Logger = zapLogger

	return nil
}

func (i *LoggerInit) Start() error {
	return nil
}

func (i *LoggerInit) Stop() error {
	return nil
}
