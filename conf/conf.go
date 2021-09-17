package conf

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port     string `default:":5012"`
	Provider string
}

var Conf Config

// Init is parse config file
func Init() (err error) {
	return envconfig.Process("mcdex", &Conf)
}
