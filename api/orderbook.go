package api

import (
	"errors"
	"githhub.com/mcdexio/mai3-data/common"
	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/ethereum"
	"githhub.com/mcdexio/mai3-data/mai3"
	"githhub.com/mcdexio/mai3-data/model"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strings"
)

func GetOrderBook(client *ethereum.Client, UnderlyingAsset string) ([]*model.AMMDepthData, []*model.AMMDepthData, error) {
	liquidityPoolStorage := client.GetLiquidityPoolStorage()
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
		return nil, nil, errors.New("perpetualIndex is invalid")
	}
	bids := mai3.GetAMMDepth(liquidityPoolStorage, perpetualIndex, isInverse, false,
		decimal.NewFromFloat(0.1), 20)
	asks := mai3.GetAMMDepth(liquidityPoolStorage, perpetualIndex, isInverse, true,
		decimal.NewFromFloat(0.1), 20)
	return bids, asks, nil
}

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

	var bids, asks []*model.AMMDepthData
	var err error
	var client *ethereum.Client
	switch collateralName {
	case common.CollateralUSDC:
		client = ethereum.NewClient(conf.Conf.ProviderArb1, conf.Conf.ReaderAddrArb1, conf.Conf.PoolAddrArb1)
	case common.CollateralBUSD:
		client = ethereum.NewClient(conf.Conf.ProviderBsc, conf.Conf.ReaderAddrBsc, conf.Conf.PoolAddrBsc)
	default:
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	bids, asks, err = GetOrderBook(client, UnderlyingAsset)
	if err != nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: model.OrderBook{
			Bids: bids,
			Asks: asks,
		},
	})
}
