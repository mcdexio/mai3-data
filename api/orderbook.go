package api

import (
	"errors"
	"net/http"
	"time"

	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/ethereum"
	"githhub.com/mcdexio/mai3-data/mai3"
	"githhub.com/mcdexio/mai3-data/model"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

var TickerPoolAddrMap = map[string]string{
	"ETH-USDC": "0xab324146c49b23658e5b3930e641bdbdf089cbac",
	"BTC-USDC": "0xab324146c49b23658e5b3930e641bdbdf089cbac",
	"BTC-BUSD": "0xdb282bbace4e375ff2901b84aceb33016d0d663d",
	"ETH-BUSD": "0xdb282bbace4e375ff2901b84aceb33016d0d663d",
	"BNB-BUSD": "0xdb282bbace4e375ff2901b84aceb33016d0d663d",
	"USD-ETH":  "0xf6b2d76c248af20009188139660a516e5c4e0532",
	"USD-BTCB": "0x2ea001032b0eb424120b4dec51bf02db0df46c78",
}

var TickerChainTypeMap = map[string]string{
	"ETH-USDC": ARB,
	"BTC-USDC": ARB,
	"BTC-BUSD": BSC,
	"ETH-BUSD": BSC,
	"BNB-BUSD": BSC,
	"USD-ETH":  BSC,
	"USD-BTCB": BSC,
}

var TickerPerpIndexMap = map[string]int64{
	"ETH-USDC": 0,
	"BTC-USDC": 1,
	"BTC-BUSD": 0,
	"ETH-BUSD": 1,
	"BNB-BUSD": 2,
	"USD-ETH":  0,
	"USD-BTCB": 1,
}

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
	tickerID := c.Param("ticker_id")
	poolAddr, ok := TickerPoolAddrMap[tickerID]
	if !ok {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	perpIndex, ok := TickerPerpIndexMap[tickerID]
	if !ok {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -2,
		})
		return
	}

	chainType, ok := TickerChainTypeMap[tickerID]
	if !ok {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -3,
		})
		return
	}

	var bids, asks []*model.AMMDepthData
	var err error
	var client *ethereum.Client
	switch chainType {
	case ARB:
		client = ethereum.NewClient(conf.Conf.ProviderArb1, conf.Conf.ReaderAddrArb1, conf.Conf.PoolAddrArb1)
	case BSC:
		client = ethereum.NewClient(conf.Conf.ProviderBsc, conf.Conf.ReaderAddrBsc, conf.Conf.PoolAddrBsc)
	default:
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -4,
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
