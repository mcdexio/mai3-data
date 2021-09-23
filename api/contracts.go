package api

import (
	"encoding/json"
	"fmt"
	"githhub.com/mcdexio/mai3-data/common"
	"githhub.com/mcdexio/mai3-data/conf"
	"githhub.com/mcdexio/mai3-data/model"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

func Contracts(c *gin.Context) {
	var responsePerpetuals common.ResponsePerpetuals
	params := common.GraphQuery{
		Query: fmt.Sprintf(common.QueryPerpetuals, conf.Conf.PoolAddr),
	}
	err, code, res := common.HttpCli.Post(conf.Conf.SubGraphURL, nil, params, nil)
	if err != nil || code != 200 {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}
	err = json.Unmarshal(res, &responsePerpetuals)
	if err != nil {
		c.JSON(http.StatusOK, model.HttpResponse{
			Code: -1,
		})
		return
	}

	var result []*model.Contract
	for _, perp := range responsePerpetuals.Data.Perpetuals {
		contract := &model.Contract{
			Index:          perp.Index,
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
		go func(contract *model.Contract) {
			FillContract(contract)
			wg.Done()
		}(contract)
	}
	wg.Wait()

	c.JSON(http.StatusOK, model.HttpResponse{
		Code: 0,
		Data: result,
	})
}

func FillContract(contract *model.Contract) *model.Contract {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(contract *model.Contract) {
		getDataFromDb(contract)
		wg.Done()
	}(contract)
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
