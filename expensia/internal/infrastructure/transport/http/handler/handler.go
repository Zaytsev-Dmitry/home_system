package handler

import (
	"expensia/pkg/config_loader"
)

type ExpensiaApi struct {
}

func NewExpensiaApi(config *config_loader.AppConfig) *ExpensiaApi {
	return &ExpensiaApi{}
}
