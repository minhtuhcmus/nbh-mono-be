package main

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/config"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/registry"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const DefaultConfigPath = "config_dev.yaml"

func main() {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = DefaultConfigPath
	}

	err := config.Setup(fmt.Sprintf("config/%s", DefaultConfigPath))
	if err != nil {
		panic(fmt.Errorf("error getting config %v", err))
	}

	err = datastore.SetupDB()
	if err != nil {
		panic(fmt.Errorf("error connecting to db %v", err))
	}

	datastore.SetupCache()

	listener, errChan := runHTTPServer()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func(errChan chan error) {
		errCh := errChan
		select {
		case <-sigCh:
			_ = listener.Close()
		case _ = <-errCh:
			_ = listener.Close()
		}
		cancel()
	}(errChan)
	<-ctx.Done()
}

func runHTTPServer() (listener net.Listener, ch chan error) {
	router, err := registry.InitHTTPServer(context.Background())
	if err != nil {
		panic(err)
	}

	conf := config.GetConfig()

	addr := fmt.Sprintf(":%s", conf.Server.ServerPort)
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	ch = make(chan error)
	go func() {
		ch <- http.Serve(listener, router)
	}()

	return listener, ch
}
