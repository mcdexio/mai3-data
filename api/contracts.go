package api

import (
	"encoding/json"
	"fmt"
	"githhub.com/mcdexio/mai3-data/common"
	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/ethereum"
	"githhub.com/mcdexio/mai3-data/mai3"
	"githhub.com/mcdexio/mai3-data/model"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func queryGraph() *common.ResponsePerpetuals {
	var responsePerpetuals *common.ResponsePerpetuals
	params := common.GraphQuery{
		Query: fmt.Sprintf(common.QueryPerpetuals, conf.Conf.PoolAddr),
	}
	err, code, res := common.HttpCli.Post(conf.Conf.SubGraphURL, nil, params, nil)
	if err != nil || code != 200 {
		return nil
	}
	err = json.Unmarshal(res, &responsePerpetuals)
	if err != nil {
		return nil
	}
	return responsePerpetuals
}

func Contracts(c *gin.Context) {
	var responsePerpetuals *common.ResponsePerpetuals
	var liquidityPoolStorage *model.LiquidityPoolStorage

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		responsePerpetuals = queryGraph()
		wg.Done()
	}()
	go func() {
		liquidityPoolStorage = ethereum.Client.GetLiquidityPoolStorage()
		wg.Done()
	}()
	wg.Wait()

	if responsePerpetuals == nil || liquidityPoolStorage == nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}

	var result []*model.Contract
	for _, perp := range responsePerpetuals.Data.Perpetuals {
		index, _ := strconv.ParseInt(perp.Index, 10, 64)
		contract := &model.Contract{
			Index:          index,
			BaseCurrency:   perp.Underlying,
			TargetCurrency: perp.CollateralName,
			TickerId:       fmt.Sprintf("%s-%s", perp.Underlying, perp.CollateralName),
			LastPrice:      perp.LastPrice,
			ProductType:    "Perpetual",
			OpenInterest:   perp.OpenInterest,
		}
		result = append(result, contract)
	}

	var wg2 sync.WaitGroup
	for _, contract := range result {
		wg2.Add(1)
		go func(contract *model.Contract, liquidityPoolStorage *model.LiquidityPoolStorage) {
			FillContract(contract, liquidityPoolStorage)
			wg2.Done()
		}(contract, liquidityPoolStorage)
	}
	wg2.Wait()

	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: result,
	})
}

func FillContract(contract *model.Contract, liquidityPoolStorage *model.LiquidityPoolStorage) *model.Contract {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(contract *model.Contract) {
		getDataFromDb(contract)
		wg.Done()
	}(contract)
	go func() {
		perpetual := liquidityPoolStorage.Perpetuals[contract.Index]
		contract.IndexPrice = perpetual.IndexPrice
		contract.IndexName = perpetual.UnderlyingAsset
		contract.IndexCurrency = contract.TargetCurrency
		contract.FundingRate = perpetual.FundingRate
		contract.Bid = mai3.ComputeBestAskBidPrice(liquidityPoolStorage, contract.Index, true)
		contract.Ask = mai3.ComputeBestAskBidPrice(liquidityPoolStorage, contract.Index, false)
		if perpetual.IsInversePerpetual {
			contract.ContractType = "Inverse"
		} else {
			contract.ContractType = "Vanilla"
		}
		contract.ContractPrice = contract.IndexPrice
		contract.ContractPriceCurrency = contract.IndexCurrency
		wg.Done()
	}()
	wg.Wait()
	return contract
}

func getDataFromDb(contract *model.Contract) {
	db := common.DbInstance()
	var tradeData model.DbTradeData
	lastDay := time.Now().Add(-24 * time.Hour).Format("2006-01-02 15:04:05")
	err := db.Table("t_trade").
		Select("sum(abs(position)) as base_volume,sum(abs(position)*price) as target_volume,"+
			"max(price) as high,min(price) as low").
		Where("timestamp>? and pool_address=? and perpetual_index=?", lastDay,
			conf.Conf.PoolAddr, contract.Index).Scan(&tradeData).Error
	if err != nil {
		logger.Warn("query pgsql error", err)
		return
	}
	contract.BaseVolume = tradeData.BaseVolume
	contract.TargetVolume = tradeData.TargetVolume.Round(18)
	contract.High = tradeData.High
	contract.Low = tradeData.Low
}
