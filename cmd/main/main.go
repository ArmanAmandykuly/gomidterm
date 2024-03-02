package main

import (
	"github.com/ArmanAmandykuly/gomidterm/pkg/database/postgres/postgres.go"
	"github.com/ArmanAmandykuly/gomidterm/internal/handler/handler.go"
)

func main() {
	DBInit()
	SetupRouter();
}