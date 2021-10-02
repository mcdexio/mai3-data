package api

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/ethereum"
	"githhub.com/mcdexio/mai3-data/mai3"
	"githhub.com/mcdexio/mai3-data/model"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func GetOrderBook(client *ethereum.Client, poolAddr string, perpIndx int64) ([]*model.AMMDepthData, []*model.AMMDepthData, error) {
	liquidityPoolStorageMap := client.GetLiquidityPoolStorage()
	if _, ok := liquidityPoolStorageMap[poolAddr]; !ok {
		return nil, nil, errors.New("pool not exist")
	}
	var isInverse bool
	liquidityPoolStorage := liquidityPoolStorageMap[poolAddr]
	if perp, ok := liquidityPoolStorage.Perpetuals[perpIndx]; ok {
		isInverse = perp.IsInversePerpetual
	} else {
		return nil, nil, errors.New("perpetualIndex is invalid")
	}

	bids := mai3.GetAMMDepth(liquidityPoolStorage, perpIndx, isInverse, false,
		decimal.NewFromFloat(0.1), 20)
	asks := mai3.GetAMMDepth(liquidityPoolStorage, perpIndx, isInverse, true,
		decimal.NewFromFloat(0.1), 20)
	return bids, asks, nil
}

func OrderBook(c *gin.Context) {
	chainType := c.Param("chain")
	poolAddr := c.Param("pool_addr")
	perpIndex, err := strconv.ParseInt(c.Param("perp_index"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}

	if chainType != BSC && chainType != ARB {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}

	var bids, asks []*model.AMMDepthData
	var client *ethereum.Client
	switch chainType {
	case ARB:
		client = ethereum.NewClient(conf.Conf.ProviderArb1, conf.Conf.ReaderAddrArb1, conf.Conf.PoolAddrArb1)
	case BSC:
		client = ethereum.NewClient(conf.Conf.ProviderBsc, conf.Conf.ReaderAddrBsc, conf.Conf.PoolAddrBsc)
	default:
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	bids, asks, err = GetOrderBook(client, poolAddr, perpIndex)
	if err != nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: model.OrderBook{
			Timestamp: time.Now().Unix(),
			Bids:      bids,
			Asks:      asks,
		},
	})
}
