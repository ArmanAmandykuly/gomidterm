package main

import (
	"github.com/ArmanAmandykuly/gomidterm/pkg/database/postgres"
	"github.com/ArmanAmandykuly/gomidterm/internal/handler"
)

func main() {
	postgres.DBInit()
	handler.SetupRouter();
}