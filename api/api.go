package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//success code, usually 200
	Code int
	//Account balance
	Balance int64
}

type Error struct {
	Code    int
	Message string
}
