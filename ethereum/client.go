package ethereum

import (
	"context"
	"githhub.com/mcdexio/mai3-data/ethereum/reader"
	"githhub.com/mcdexio/mai3-data/mai3"
	"githhub.com/mcdexio/mai3-data/model"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
)

type Client struct {
	ethCli     *ethclient.Client
	readerAddr string
	poolAddr   string
}

func NewClient(provider, readerAddr, poolAddr string) *Client {
	var err error
	ethCli, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("eth client dial error", err)
	}
	client := &Client{
		ethCli:     ethCli,
		readerAddr: readerAddr,
		poolAddr:   poolAddr,
	}
	return client
}

func (client *Client) GetLiquidityPoolStorage() *model.LiquidityPoolStorage {
	address := ethCommon.HexToAddress(client.readerAddr)
	contract, err := reader.NewReader(address, client.ethCli)
	if err != nil {
		return nil
	}

	liquidityPool := ethCommon.HexToAddress(client.poolAddr)
	opts := &ethBind.CallOpts{
		Context: context.Background(),
	}
	res, err := contract.GetLiquidityPoolStorage(opts, liquidityPool)
	if err != nil {
		return nil
	}
	if !res.IsSynced {
		return nil
	}
	rsp := &model.LiquidityPoolStorage{}
	rsp.VaultFeeRate = decimal.NewFromBigInt(res.Pool.IntNums[0], -mai3.DECIMALS)
	rsp.PoolCashBalance = decimal.NewFromBigInt(res.Pool.IntNums[1], -mai3.DECIMALS)
	rsp.Perpetuals = make(map[int64]*model.PerpetualStorage)

	for i, perpetual := range res.Pool.Perpetuals {
		storage := &model.PerpetualStorage{
			MarkPrice:               decimal.NewFromBigInt(perpetual.Nums[1], -mai3.DECIMALS),
			IndexPrice:              decimal.NewFromBigInt(perpetual.Nums[2], -mai3.DECIMALS),
			FundingRate:             decimal.NewFromBigInt(perpetual.Nums[3], -mai3.DECIMALS),
			UnitAccumulativeFunding: decimal.NewFromBigInt(perpetual.Nums[4], -mai3.DECIMALS),
			InitialMarginRate:       decimal.NewFromBigInt(perpetual.Nums[5], -mai3.DECIMALS),
			MaintenanceMarginRate:   decimal.NewFromBigInt(perpetual.Nums[6], -mai3.DECIMALS),
			OperatorFeeRate:         decimal.NewFromBigInt(perpetual.Nums[7], -mai3.DECIMALS),
			LpFeeRate:               decimal.NewFromBigInt(perpetual.Nums[8], -mai3.DECIMALS),
			ReferrerRebateRate:      decimal.NewFromBigInt(perpetual.Nums[9], -mai3.DECIMALS),
			LiquidationPenaltyRate:  decimal.NewFromBigInt(perpetual.Nums[10], -mai3.DECIMALS),
			KeeperGasReward:         decimal.NewFromBigInt(perpetual.Nums[11], -mai3.DECIMALS),
			InsuranceFundRate:       decimal.NewFromBigInt(perpetual.Nums[12], -mai3.DECIMALS),
			HalfSpread:              decimal.NewFromBigInt(perpetual.Nums[13], -mai3.DECIMALS),
			OpenSlippageFactor:      decimal.NewFromBigInt(perpetual.Nums[16], -mai3.DECIMALS),
			CloseSlippageFactor:     decimal.NewFromBigInt(perpetual.Nums[19], -mai3.DECIMALS),
			FundingRateLimit:        decimal.NewFromBigInt(perpetual.Nums[22], -mai3.DECIMALS),
			MaxLeverage:             decimal.NewFromBigInt(perpetual.Nums[25], -mai3.DECIMALS),
			MaxClosePriceDiscount:   decimal.NewFromBigInt(perpetual.Nums[28], -mai3.DECIMALS),
			OpenInterest:            decimal.NewFromBigInt(perpetual.Nums[31], -mai3.DECIMALS),
			MaxOpenInterestRate:     decimal.NewFromBigInt(perpetual.Nums[32], -mai3.DECIMALS),
			FundingRateFactor:       decimal.NewFromBigInt(perpetual.Nums[33], -mai3.DECIMALS),
			AmmCashBalance:          decimal.NewFromBigInt(perpetual.AmmCashBalance, -mai3.DECIMALS),
			AmmPositionAmount:       decimal.NewFromBigInt(perpetual.AmmPositionAmount, -mai3.DECIMALS),
			UnderlyingAsset:         perpetual.UnderlyingAsset,
			IsInversePerpetual:      perpetual.IsInversePerpetual,
		}
		if perpetual.State == model.PerpetualNormal {
			storage.IsNormal = true
		}
		rsp.Perpetuals[int64(i)] = storage
	}
	return rsp
}
