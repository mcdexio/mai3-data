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

func queryGraph(graphURL, poolAddr string) *common.ResponsePerpetuals {
	var responsePerpetuals *common.ResponsePerpetuals
	params := common.GraphQuery{
		Query: fmt.Sprintf(common.QueryPerpetuals, poolAddr),
	}
	err, code, res := common.HttpCli.Post(graphURL, nil, params, nil)
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
	var perpetualsArb, perpetualsBsd *common.ResponsePerpetuals
	var liquidityPoolArb, liquidityPoolBsd *model.LiquidityPoolStorage

	var wg sync.WaitGroup
	wg.Add(4)
	// arb1 network
	go func() {
		perpetualsArb = queryGraph(conf.Conf.SubGraphUrlArb1, conf.Conf.PoolAddrArb1)
		wg.Done()
	}()
	go func() {
		client := ethereum.NewClient(conf.Conf.ProviderArb1, conf.Conf.ReaderAddrArb1, conf.Conf.PoolAddrArb1)
		liquidityPoolArb = client.GetLiquidityPoolStorage()
		wg.Done()
	}()
	// bsc network
	go func() {
		perpetualsBsd = queryGraph(conf.Conf.SubGraphUrlBsc, conf.Conf.PoolAddrBsc)
		wg.Done()
	}()
	go func() {
		client := ethereum.NewClient(conf.Conf.ProviderBsc, conf.Conf.ReaderAddrBsc, conf.Conf.PoolAddrBsc)
		liquidityPoolBsd = client.GetLiquidityPoolStorage()
		wg.Done()
	}()
	wg.Wait()

	if perpetualsArb == nil || perpetualsBsd == nil || liquidityPoolArb == nil || liquidityPoolBsd == nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}

	var wg2 sync.WaitGroup
	wg2.Add(2)
	var resultArb, resultBsc []*model.Contract
	go func() {
		resultArb = buildContractList(perpetualsArb, liquidityPoolArb)
		wg2.Done()
	}()
	go func() {
		resultBsc = buildContractList(perpetualsBsd, liquidityPoolBsd)
		wg2.Done()
	}()
	wg2.Wait()

	result := append(resultArb, resultBsc...)
	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: result,
	})
}

func buildContractList(responsePerpetuals *common.ResponsePerpetuals, liquidityPoolStorage *model.LiquidityPoolStorage) []*model.Contract {
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

	var wg sync.WaitGroup
	for _, contract := range result {
		wg.Add(1)
		go func(contract *model.Contract, liquidityPoolStorage *model.LiquidityPoolStorage) {
			fillContract(contract, liquidityPoolStorage)
			wg.Done()
		}(contract, liquidityPoolStorage)
	}
	wg.Wait()
	return result
}

func fillContract(contract *model.Contract, liquidityPoolStorage *model.LiquidityPoolStorage) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(contract *model.Contract) {
		getTradeDataFromDb(contract)
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
}

func getTradeDataFromDb(contract *model.Contract) {
	db := common.DbInstance()
	var tradeData model.DbTradeData
	var tableName, poolAddr string
	switch contract.TargetCurrency {
	case common.CollateralUSDC:
		tableName = "arb_trade"
		poolAddr = conf.Conf.PoolAddrArb1
	case common.CollateralBUSD:
		tableName = "bsc_trade"
		poolAddr = conf.Conf.PoolAddrBsc
	default:
		return
	}
	lastDay := time.Now().Add(-24 * time.Hour).Format("2006-01-02 15:04:05")
	err := db.Table(tableName).
		Select("sum(abs(position)) as base_volume,sum(abs(position)*price) as target_volume,"+
			"max(price) as high,min(price) as low").
		Where("timestamp>? and pool_address=? and perpetual_index=?", lastDay,
			poolAddr, contract.Index).Scan(&tradeData).Error
	if err != nil {
		logger.Warn("query pgsql error", err)
		return
	}
	contract.BaseVolume = tradeData.BaseVolume
	contract.TargetVolume = tradeData.TargetVolume.Round(18)
	contract.High = tradeData.High
	contract.Low = tradeData.Low
}
