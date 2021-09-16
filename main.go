package main

import (
	"githhub.com/mcdexio/mai3-data/erc20"
	"githhub.com/mcdexio/mai3-data/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
)

const SERVERBUSY = "server is busy"

func cmc(c *gin.Context) {
	questParam := c.DefaultQuery("q", utils.TOTAL_SUPPLY)

	switch questParam {
	case utils.TOTAL_SUPPLY:
		client, err := ethclient.Dial(utils.PROVIDER)
		if err != nil {
			c.String(http.StatusInternalServerError, SERVERBUSY)
		}

		tokenAddress := common.HexToAddress(utils.MCB_ADDRESS)
		instance, err := erc20.NewToken(tokenAddress, client)
		if err != nil {
			c.String(http.StatusInternalServerError, SERVERBUSY)
		}
		res, err := instance.TotalSupply(nil)
		if err != nil {
			c.String(http.StatusInternalServerError, SERVERBUSY)
		}
		totalSupply := decimal.NewFromBigInt(res, -18)
		c.String(http.StatusOK, totalSupply.String())
	default:
		c.String(http.StatusBadRequest, "request is invalid")
	}
}

func main() {
	router := gin.Default()
	data := router.Group("/data")
	data.GET("/cmc", cmc)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":5001")

}
