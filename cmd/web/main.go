package main

import (
	"context"
	"fmt"
	"github.com/lucysuddenly/rube/internal/rube"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := rube.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config")
	}

	// todo: add tracing

	srv := http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           rube.Handler(cfg),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	c := make(chan os.Signal, 10)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-c
		log.Printf("shutting down on interrupt, info: %+v", ctx)
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("shutting down on interrupt, err: %s, info: %+v", err.Error(), ctx)
		}
		cancel()
	}()

	// start server
	err = srv.ListenAndServe()
	fmt.Print("starting server on port "+cfg.Port)
	if err != nil && err != http.ErrServerClosed {
		log.Printf("shutting down on interrupt, err: %s, info: %+v", err.Error(), ctx)
		return
	}

	// wait for shutdown to finish
	<-ctx.Done()
}