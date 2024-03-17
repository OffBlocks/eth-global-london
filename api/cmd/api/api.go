package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/offblocks/eth-global-london/api/pkg/api"
	"github.com/offblocks/eth-global-london/api/pkg/blockchain"
	"github.com/offblocks/eth-global-london/api/pkg/config"
	"github.com/offblocks/eth-global-london/api/pkg/db"
	"github.com/offblocks/eth-global-london/api/pkg/repository"
	"github.com/offblocks/eth-global-london/api/pkg/service"
	"github.com/offblocks/eth-global-london/api/pkg/temporal"
	"github.com/offblocks/eth-global-london/api/pkg/worker"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	h := db.Init(c.DBUrl)
	repository := repository.NewRepository(&h)

	ethereumBlockchain, err := blockchain.NewEthereumBlockchain(&c)
	if err != nil {
		log.Fatalln("Failed to create Ethereum blockchain:", err)
	}

	baseBlockchain, err := blockchain.NewBaseBlockchain(&c)
	if err != nil {
		log.Fatalln("Failed to create Base blockchain:", err)
	}

	temporalClient, err := temporal.NewClient(&c)
	if err != nil {
		log.Fatalln("Failed to connect to Temporal:", err)
	}
	defer temporalClient.Close()

	worker := worker.NewWorker(temporalClient, &repository, baseBlockchain, ethereumBlockchain)
	if err = worker.Start(); err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
	defer worker.Stop()

	apiService := service.NewApiService(repository, baseBlockchain, ethereumBlockchain, temporalClient)
	strictApiService := api.NewStrictHandler(apiService, nil)

	r := gin.Default()
	r.ContextWithFallback = true

	errorHandler := func(c *gin.Context, err error, statusCode int) {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
	}

	api.RegisterHandlersWithOptions(r, strictApiService, api.GinServerOptions{
		BaseURL:      c.ApiBaseUrl,
		ErrorHandler: errorHandler,
	})

	addr := net.JoinHostPort("0.0.0.0", c.Port)

	s := &http.Server{
		Handler: r,
		Addr:    addr,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Println("API serving on", addr)
		errCh <- s.ListenAndServe()
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	select {
	case <-sigCh:
		s.Shutdown(ctx)
	case err := <-errCh:
		log.Fatalln("Failed to serve:", err)
	}
}
