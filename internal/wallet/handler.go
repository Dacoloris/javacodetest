package wallet

import (
	"fmt"
	"github.com/dacoloris/javacodetest/configs"
	"github.com/dacoloris/javacodetest/pkg/req"
	"github.com/dacoloris/javacodetest/pkg/res"
	"github.com/google/uuid"
	"net/http"
)

type HandlerDeps struct {
	IRepository
	WalletService
	*configs.Config
}

type Handler struct {
	IRepository
	WalletService
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		IRepository:   deps.IRepository,
		WalletService: deps.WalletService,
	}

	router.Handle("POST /api/v1/wallet", handler.WalletUpdate())
	router.Handle("GET /api/v1/wallets/{WALLET_UUID}", handler.WalletGet())
}

func (handler *Handler) WalletGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID, err := uuid.Parse(r.PathValue("WALLET_UUID"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		balance, err := handler.WalletService.GetBalance(walletID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		res.Json(w, fmt.Sprintf("Balance: %f", balance), http.StatusCreated)
	}
}

func (handler *Handler) WalletUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[OperationRequest](&w, r)
		if err != nil {
			return
		}

		balance, _ := handler.WalletService.GetBalance(body.WalletID)
		result, err := handler.WalletService.CalcBalance(balance, body.OperationType, body.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.IRepository.Update(&Wallet{
			ID:      body.WalletID,
			Balance: result,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, nil, http.StatusCreated)
	}
}
