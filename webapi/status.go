package webapi

import (
	"time"

	"github.com/gin-gonic/gin"
)

// pubKey is the handler for "GET /pubkey"
func pubKey(c *gin.Context) {
	sendJSONResponse(pubKeyResponse{
		Timestamp: time.Now().Unix(),
		PubKey:    cfg.PubKey,
	}, c)
}

// fee is the handler for "GET /fee"
func fee(c *gin.Context) {
	sendJSONResponse(feeResponse{
		Timestamp: time.Now().Unix(),
		Fee:       cfg.VSPFee,
	}, c)
}
