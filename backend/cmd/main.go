package main

import (
	"fmt"

	"github.com/AmadoMuerte/sysStats/internal/config"
	"github.com/AmadoMuerte/sysStats/internal/db"
	"github.com/AmadoMuerte/sysStats/internal/http-server/server"
)

// @title sysStats API
// @version 1.0
// @description API для sysStats
// @contact.name GitHub
// @contact.url https://github.com/AmadoMuerte
// @basePath /api/v1
// @schemes http https
func main() {
	cfg, err := config.NewConfig(nil)
	if err != nil {
		panic(err)
	}

	storage := db.New(cfg)
	fmt.Print("Connected to DB\n")
	server := server.New(cfg, storage)
	server.Start()
}
