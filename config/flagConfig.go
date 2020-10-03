package config

import (
	"github.com/namsral/flag"
)

type FlagConfig struct{
	DbUrl string
	DbDatabase string
	AppSecret string
}

var Fconfig FlagConfig


func InitFlagConfig() {

	flag.StringVar(&Fconfig.DbUrl, "dburl", "", "Database URL")
	flag.StringVar(&Fconfig.DbDatabase, "dbdatabase", "", "Database Database")
	flag.StringVar(&Fconfig.AppSecret, "appsecret", "", "App secret")
	flag.Parse()
}

