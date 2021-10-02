package conf

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port            string `default:":5012"`
	ProviderL1      string
	DbConnStr       string
	ProviderArb1    string
	ProviderBsc     string
	SubGraphUrlArb1 string
	SubGraphUrlBsc  string
	PoolAddrArb1    []string
	PoolAddrBsc     []string
	ReaderAddrArb1  string
	ReaderAddrBsc   string
}

var Conf Config

// Init is parse config file
func Init() (err error) {
	return envconfig.Process("mcdex", &Conf)
}
