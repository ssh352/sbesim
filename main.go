package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"log"
	"sbe/configure"
	"sbe/handler"
	"sbe/record"
	"sbe/service"
	"time"
)

func main() {

	app := fx.New(
		fx.Provide(
			configure.New,
			handler.New,
			service.New,
		),
		fx.Invoke(func(s service.SBEServer, c *record.Config) error {
			ctx := context.Background()
			port := fmt.Sprintf(":%s", c.Port)
			if err := s.Listen(ctx, port); err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			go s.Start(ctx)
			return nil
		}),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}