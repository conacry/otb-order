package main

import (
	"context"
	"fmt"
	initapp "github.com/conacry/go-platform/pkg/init"
	"gopkg.in/yaml.v3"
	"log"
	initApp "online-shop-order/initApp"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	ctx, _ := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	configPath, err := filepath.Abs(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalln("illegal config path: ", err)
	}

	f, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln("can't find config file: ", err)
	}

	config := &initApp.AppConfig{}
	err = yaml.Unmarshal(f, config)
	if err != nil {
		log.Fatalln("can't start application: ", err)
	}

	dc, err := initApp.NewDependencyContainer(ctx, config)
	if err != nil {
		log.Fatalln("can't create dependency container: ", err)
	}

	depInitializer, _ := initapp.NewDependenciesInitializer().
		DependencyInitChain(initApp.GetDependencyInitChains(dc)).
		Build()

	err = depInitializer.InitDependencies()
	if err != nil {
		log.Fatalln("can't start application: ", err)
	}

	err = depInitializer.StartDependencies()
	if err != nil {
		log.Fatalln("Can't start application: ", err)
	}

	fmt.Println("Application is started successfully")
	<-ctx.Done()

	err = depInitializer.StopDependencies()
	if err != nil {
		log.Fatalln("All systems closed with errors. LastError: ", err)
	}

	fmt.Println("Application is shouted down successfully")
}
