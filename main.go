package main

import (
	"githhub.com/mcdexio/mai3-data/api"
	"githhub.com/mcdexio/mai3-data/common"
	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/ethereum/erc20"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
)

func cmc(c *gin.Context) {
	questParam := c.DefaultQuery("q", common.TOTAL_SUPPLY)

	switch questParam {
	case common.TOTAL_SUPPLY:
		client, err := ethclient.Dial(conf.Conf.ProviderL1)
		if err != nil {
			c.String(http.StatusInternalServerError, common.SERVERBUSY)
		}

		tokenAddress := eth_common.HexToAddress(common.MCB_ADDRESS)
		instance, err := erc20.NewToken(tokenAddress, client)
		if err != nil {
			c.String(http.StatusInternalServerError, common.SERVERBUSY)
		}
		res, err := instance.TotalSupply(nil)
		if err != nil {
			c.String(http.StatusInternalServerError, common.SERVERBUSY)
		}
		totalSupply := decimal.NewFromBigInt(res, -18)
		c.String(http.StatusOK, totalSupply.String())
	default:
		c.String(http.StatusBadRequest, "request is invalid")
	}
}

func main() {
	// config decimal marshal json to number
	decimal.MarshalJSONWithoutQuotes = true
	// init config
	if err := conf.Init(); err != nil {
		log.Fatal("init config error, ", err)
	}

	router := gin.Default()
	data := router.Group("/data")
	data.GET("/cmc", cmc)
	data.GET("/contracts", api.Contracts)
	data.GET("/orderbook/:ticker_id", api.OrderBook)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(conf.Conf.Port)
}
