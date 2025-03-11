package configs

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"log"
)

func LoadConfig() *AppConfig {
	var appConfig AppConfig
	ctx := context.Background()
	if err := envconfig.Process(ctx, &appConfig); err != nil {
		log.Fatal(err)
	}
	return &appConfig
}
