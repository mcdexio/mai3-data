package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"githhub.com/mcdexio/mai3-data/common"
	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/ethereum"
	"githhub.com/mcdexio/mai3-data/mai3"
	"githhub.com/mcdexio/mai3-data/model"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

const BSC = "bsc"
const ARB = "arbitrum"

const EthUsdcPerpetual = "0xdb282bbace4e375ff2901b84aceb33016d0d663d-1"
const BtcEthPerpetual = "0xf6b2d76c248af20009188139660a516e5c4e0532-1"

func queryGraph(graphURL string, poolAddres []string) *common.ResponsePerpetuals {
	var responsePerpetuals *common.ResponsePerpetuals
	params := common.GraphQuery{
		Query: fmt.Sprintf(common.QueryPerpetuals, strings.Join(poolAddres, "\",\"")),
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
	var perpetualsArb, perpetualsBsc *common.ResponsePerpetuals
	var liquidityPoolArb, liquidityPoolBsc map[string]*model.LiquidityPoolStorage

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
		perpetualsBsc = queryGraph(conf.Conf.SubGraphUrlBsc, conf.Conf.PoolAddrBsc)
		wg.Done()
	}()
	go func() {
		client := ethereum.NewClient(conf.Conf.ProviderBsc, conf.Conf.ReaderAddrBsc, conf.Conf.PoolAddrBsc)
		liquidityPoolBsc = client.GetLiquidityPoolStorage()
		wg.Done()
	}()
	wg.Wait()

	if perpetualsArb == nil || perpetualsBsc == nil || liquidityPoolArb == nil || liquidityPoolBsc == nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}

	var wg2 sync.WaitGroup
	wg2.Add(2)
	var resultArb, resultBsc []*model.Contract
	go func() {
		resultArb = buildContractList(ARB, perpetualsArb, liquidityPoolArb)
		wg2.Done()
	}()
	go func() {
		resultBsc = buildContractList(BSC, perpetualsBsc, liquidityPoolBsc)
		wg2.Done()
	}()
	wg2.Wait()

	result := append(resultArb, resultBsc...)
	newResult := make([]*model.Contract, 0)
	ethPrice := decimal.Zero
	// modify contract which is inverse
	for _, contract := range result {
		if contract.ContractType == "Inverse" {
			contract.BaseCurrency, contract.TargetCurrency = contract.TargetCurrency, contract.BaseCurrency
			contract.BaseVolume, contract.TargetVolume = contract.TargetVolume, contract.BaseVolume
			contract.IndexCurrency = contract.TargetCurrency
			contract.ContractPriceCurrency = contract.IndexCurrency
			if !contract.LastPrice.IsZero() {
				contract.LastPrice = decimal.NewFromInt(1).Div(contract.LastPrice)
			}
			if !contract.Bid.IsZero() {
				contract.Bid = decimal.NewFromInt(1).Div(contract.Bid)
			}
			if !contract.Ask.IsZero() {
				contract.Ask = decimal.NewFromInt(1).Div(contract.Ask)
			}
			if !contract.High.IsZero() {
				contract.High = decimal.NewFromInt(1).Div(contract.High)
			}
			if !contract.Low.IsZero() {
				contract.Low = decimal.NewFromInt(1).Div(contract.Low)
			}
			if !contract.IndexPrice.IsZero() {
				contract.IndexPrice = decimal.NewFromInt(1).Div(contract.IndexPrice)
			}
			if !contract.ContractPrice.IsZero() {
				contract.ContractPrice = decimal.NewFromInt(1).Div(contract.ContractPrice)
			}
			contract.OpenInterest = contract.OpenInterest.Div(contract.LastPrice)
		}

		// for BTC-ETH perpetual
		perpetual := fmt.Sprintf("%s-%d", contract.PoolAddr, contract.Index)
		if perpetual == EthUsdcPerpetual {
			ethPrice = contract.LastPrice
		} else if perpetual == BtcEthPerpetual {
			contract.LastPrice = contract.LastPrice.Mul(ethPrice)
			contract.Bid = contract.Bid.Mul(ethPrice)
			contract.Ask = contract.Ask.Mul(ethPrice)
			contract.High = contract.High.Mul(ethPrice)
			contract.Low = contract.Low.Mul(ethPrice)
			contract.IndexPrice = contract.IndexPrice.Mul(ethPrice)
			contract.ContractPrice = contract.ContractPrice.Mul(ethPrice)
			contract.TargetVolume = contract.TargetVolume.Mul(ethPrice)
			contract.TargetCurrency = "USD"
			contract.IndexCurrency = "USD"
			contract.ContractPriceCurrency = "USD"
		}

		newResult = append(newResult, contract)
	}
	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: newResult,
	})
}

func ContractsByChainType(c *gin.Context) {
	chainType := c.Param("chain_type")
	var perpetuals *common.ResponsePerpetuals
	var liquidityPool map[string]*model.LiquidityPoolStorage
	subgraphURL := conf.Conf.SubGraphUrlArb1
	poolAddr := conf.Conf.PoolAddrArb1
	readerAddr := conf.Conf.ReaderAddrArb1
	provider := conf.Conf.ProviderArb1
	var wg sync.WaitGroup
	wg.Add(2)
	if chainType == BSC {
		subgraphURL = conf.Conf.SubGraphUrlBsc
		poolAddr = conf.Conf.PoolAddrBsc
		readerAddr = conf.Conf.ReaderAddrBsc
		provider = conf.Conf.ProviderBsc
	} else {
		chainType = ARB
	}
	// arb1 network
	go func() {
		perpetuals = queryGraph(subgraphURL, poolAddr)
		wg.Done()
	}()
	go func() {
		client := ethereum.NewClient(provider, readerAddr, poolAddr)
		liquidityPool = client.GetLiquidityPoolStorage()
		wg.Done()
	}()
	wg.Wait()
	if perpetuals == nil || liquidityPool == nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}

	result := buildContractList(chainType, perpetuals, liquidityPool)
	newResult := make([]*model.Contract, 0)
	ethPrice := decimal.Zero
	// modify contract which is inverse
	for _, contract := range result {
		if contract.ContractType == "Inverse" {
			contract.BaseCurrency, contract.TargetCurrency = contract.TargetCurrency, contract.BaseCurrency
			contract.BaseVolume, contract.TargetVolume = contract.TargetVolume, contract.BaseVolume
			contract.IndexCurrency = contract.TargetCurrency
			contract.ContractPriceCurrency = contract.IndexCurrency
			if !contract.LastPrice.IsZero() {
				contract.LastPrice = decimal.NewFromInt(1).Div(contract.LastPrice)
			}
			if !contract.Bid.IsZero() {
				contract.Bid = decimal.NewFromInt(1).Div(contract.Bid)
			}
			if !contract.Ask.IsZero() {
				contract.Ask = decimal.NewFromInt(1).Div(contract.Ask)
			}
			if !contract.High.IsZero() {
				contract.High = decimal.NewFromInt(1).Div(contract.High)
			}
			if !contract.Low.IsZero() {
				contract.Low = decimal.NewFromInt(1).Div(contract.Low)
			}
			if !contract.IndexPrice.IsZero() {
				contract.IndexPrice = decimal.NewFromInt(1).Div(contract.IndexPrice)
			}
			if !contract.ContractPrice.IsZero() {
				contract.ContractPrice = decimal.NewFromInt(1).Div(contract.ContractPrice)
			}
			contract.OpenInterest = contract.OpenInterest.Div(contract.LastPrice)
		}

		// for BTC-ETH perpetual
		perpetual := fmt.Sprintf("%s-%d", contract.PoolAddr, contract.Index)
		if perpetual == EthUsdcPerpetual {
			ethPrice = contract.LastPrice
		} else if perpetual == BtcEthPerpetual {
			contract.LastPrice = contract.LastPrice.Mul(ethPrice)
			contract.Bid = contract.Bid.Mul(ethPrice)
			contract.Ask = contract.Ask.Mul(ethPrice)
			contract.High = contract.High.Mul(ethPrice)
			contract.Low = contract.Low.Mul(ethPrice)
			contract.IndexPrice = contract.IndexPrice.Mul(ethPrice)
			contract.ContractPrice = contract.ContractPrice.Mul(ethPrice)
			contract.TargetVolume = contract.TargetVolume.Mul(ethPrice)
			contract.TargetCurrency = "USD"
			contract.IndexCurrency = "USD"
			contract.ContractPriceCurrency = "USD"
		}

		newResult = append(newResult, contract)
	}
	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: newResult,
	})
}

func buildContractList(chainType string, responsePerpetuals *common.ResponsePerpetuals, liquidityPoolStorage map[string]*model.LiquidityPoolStorage) []*model.Contract {
	var result []*model.Contract
	for _, perp := range responsePerpetuals.Data.Perpetuals {
		index, _ := strconv.ParseInt(perp.Index, 10, 64)
		contract := &model.Contract{
			PoolAddr:       strings.Split(perp.Id, "-")[0],
			Index:          index,
			Symbol:         perp.Symbol,
			ChainType:      chainType,
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
		go func(chainType string, contract *model.Contract, liquidityPoolStorage *model.LiquidityPoolStorage) {
			fillContract(chainType, contract, liquidityPoolStorage)
			wg.Done()
		}(chainType, contract, liquidityPoolStorage[contract.PoolAddr])
	}
	wg.Wait()
	return result
}

func fillContract(chainType string, contract *model.Contract, liquidityPoolStorage *model.LiquidityPoolStorage) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func(chainType string, contract *model.Contract) {
		getTradeDataFromDb(chainType, contract)
		wg.Done()
	}(chainType, contract)
	go func() {
		perpetual := liquidityPoolStorage.Perpetuals[contract.Index]
		contract.IndexPrice = perpetual.IndexPrice
		contract.IndexName = perpetual.UnderlyingAsset
		contract.IndexCurrency = contract.TargetCurrency
		contract.FundingRate = perpetual.FundingRate
		contract.NextFundingRateTimestamp = time.Now().Add(time.Second).Unix()
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

func getTradeDataFromDb(chainType string, contract *model.Contract) {
	db := common.DbInstance()
	var tradeData model.DbTradeData
	var tableName string
	switch chainType {
	case ARB:
		tableName = "arb_trade"
	case BSC:
		tableName = "bsc_trade"
	default:
		return
	}
	lastDay := time.Now().Add(-24 * time.Hour).Format("2006-01-02 15:04:05")
	err := db.Table(tableName).
		Select("sum(abs(position)) as base_volume,sum(abs(position)*price) as target_volume,"+
			"max(price) as high,min(price) as low").
		Where("timestamp>? and pool_address=? and perpetual_index=?", lastDay,
			contract.PoolAddr, contract.Index).Scan(&tradeData).Error
	if err != nil {
		logger.Warn("query pgsql error", err)
		return
	}
	contract.BaseVolume = tradeData.BaseVolume
	contract.TargetVolume = tradeData.TargetVolume.Round(18)
	contract.High = tradeData.High
	contract.Low = tradeData.Low
}
