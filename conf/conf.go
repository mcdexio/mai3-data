package conf

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port         string `default:":5012"`
	ProviderL1   string
	ProviderArb1 string
	DbConnStr    string
	SubGraphURL  string
	PoolAddr     string
	ReaderAddr   string
}

var Conf Config

// Init is parse config file
func Init() (err error) {
	return envconfig.Process("mcdex", &Conf)
}
