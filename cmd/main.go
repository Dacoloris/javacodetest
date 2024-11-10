package main

import (
	"fmt"
	"github.com/dacoloris/javacodetest/configs"
	"github.com/dacoloris/javacodetest/internal/wallet"
	"github.com/dacoloris/javacodetest/pkg/db"
	"github.com/dacoloris/javacodetest/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
	router := http.NewServeMux()

	walletRepository := wallet.NewRepository(database)

	walletService := wallet.NewService(walletRepository)

	wallet.NewHandler(router, wallet.HandlerDeps{
		Config:        conf,
		IRepository:   walletRepository,
		WalletService: walletService,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":" + conf.Port,
		Handler: stack(router),
	}

	fmt.Printf("Server is listening on port %s\n", conf.Port)
	panic(server.ListenAndServe())
}
