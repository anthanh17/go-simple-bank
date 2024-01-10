package api

import (
	"net/http"

	db "github.com/anthanh17/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner 	 string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required, oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	// validate input data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// create account
	arg := db.CreateAccountParams {
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// retunr account
	ctx.JSON(http.StatusOK, account)
}