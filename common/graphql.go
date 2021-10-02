package common

import "github.com/shopspring/decimal"

type GraphQuery struct {
	Query string `json:"query"`
}

var QueryPerpetuals = `{
  perpetuals(where: {liquidityPool_in: ["%v"]}) {
    id
	index
    underlying
    collateralName
    openInterest
    lastPrice
  }
}`

type GraphPerpetual struct {
	Id             string          `json:"id"`
	Index          string          `json:"index"`
	Underlying     string          `json:"underlying"`
	CollateralName string          `json:"collateralName"`
	LastPrice      decimal.Decimal `json:"lastPrice"`
	OpenInterest   decimal.Decimal `json:"openInterest"`
}

type ResponsePerpetuals struct {
	Data struct {
		Perpetuals []GraphPerpetual `json:"perpetuals"`
	} `json:"data"`
}
