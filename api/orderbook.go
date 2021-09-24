package api

import (
	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/ethereum"
	"githhub.com/mcdexio/mai3-data/mai3"
	"githhub.com/mcdexio/mai3-data/model"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strings"
)

func OrderBook(c *gin.Context) {
	tickerId := c.Param("ticker_id")
	temp := strings.Split(tickerId, "-")
	if len(temp) != 2 {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	UnderlyingAsset := temp[0]
	collateralName := temp[1]
	if collateralName != conf.Conf.PoolCollateral {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	liquidityPoolStorage := ethereum.Client.GetLiquidityPoolStorage()
	perpetualIndex := int64(-1)
	var isInverse bool
	for index, perp := range liquidityPoolStorage.Perpetuals {
		if perp.UnderlyingAsset == UnderlyingAsset {
			perpetualIndex = index
			isInverse = perp.IsInversePerpetual
			break
		}
	}
	if perpetualIndex == -1 {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	bids := mai3.GetAMMDepth(liquidityPoolStorage, perpetualIndex, isInverse, false,
		decimal.NewFromFloat(0.1), 20)
	asks := mai3.GetAMMDepth(liquidityPoolStorage, perpetualIndex, isInverse, true,
		decimal.NewFromFloat(0.1), 20)
	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: model.OrderBook{
			Bids: bids,
			Asks: asks,
		},
	})
}
