package api

import (
    "encoding/json"
    "net/http"
)

// Coin balance parameters
type CoinBalanceParams struct {
    Username string
}

// Coin balance response
type CoinBalanceResponse struct {
    // Success Code
    Code int
    
    // Account Balance
    Balance int64
}

type Error struct {
    // Error Code
    Code int
    
    // Error Message
    Message string
}
