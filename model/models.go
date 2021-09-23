package model

import "github.com/shopspring/decimal"

type HttpResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Contract struct {
	Index                    int64           `json:"index"`
	TickerId                 string          `json:"ticker_id"`
	BaseCurrency             string          `json:"base_currency"`
	TargetCurrency           string          `json:"target_currency"`
	LastPrice                decimal.Decimal `json:"last_price"`
	BaseVolume               decimal.Decimal `json:"base_volume"`
	TargetVolume             decimal.Decimal `json:"target_volume"`
	Bid                      decimal.Decimal `json:"bid"`
	Ask                      decimal.Decimal `json:"ask"`
	High                     decimal.Decimal `json:"high"`
	Low                      decimal.Decimal `json:"low"`
	ProductType              string          `json:"product_type"`
	OpenInterest             decimal.Decimal `json:"open_interest"`
	IndexPrice               decimal.Decimal `json:"index_price"`
	IndexName                string          `json:"index_name"`
	IndexCurrency            string          `json:"index_currency"`
	FundingRate              decimal.Decimal `json:"funding_rate"`
	NextFundingRate          decimal.Decimal `json:"next_funding_rate"`
	NextFundingRateTimestamp int64           `json:"next_funding_rate_timestamp"`
	ContractType             string          `json:"contract_type"`
	ContractPrice            decimal.Decimal `json:"contract_price"`
	ContractPriceCurrency    string          `json:"contract_price_currency"`
}

type DbTradeData struct {
	BaseVolume   decimal.Decimal
	TargetVolume decimal.Decimal
	High         decimal.Decimal
	Low          decimal.Decimal
}
