package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/config"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/infra/db"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/interface/handler"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/interface/router"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/registry"
)

func main() {
	code := run()
	os.Exit(code)
}

func run() int {
	dbConn, err := db.NewDBConn(config.DataSourceURL)
	if err != nil {
		fmt.Printf("[Error] unable to new db connection: %v\n", err)
		return 1
	}

	appDB, err := db.NewDB(dbConn)
	if err != nil {
		fmt.Printf("[Error] unable to new db: %v\n", err)
		return 1
	}

	reg := registry.NewRegistry(appDB)
	hdl := handler.NewHandler(reg)

	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	router.RegisterHandler(e, hdl, config.BaseURL)

	go func() {
		addr := fmt.Sprintf("0.0.0.0:%d", config.Port)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			fmt.Printf("[Error] unable to start: %v\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	<-ctx.Done()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(timeoutCtx); err != nil {
		fmt.Printf("[Error] unable to shutdown: %v\n", err)
	}
	return 0
}
