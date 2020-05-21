package config

import "os"

type BaseConfig struct {
	CRAWLER_INTERNAL_TIME string
	APPLE_ACCOUNT         string
	APPLE_PASSWORD        string
}

func (baseCfg *BaseConfig) InitConfig() {

	baseCfg.CRAWLER_INTERNAL_TIME = os.Getenv("CRAWLER_INTERNAL_TIME")

	baseCfg.APPLE_ACCOUNT = os.Getenv("APPLE_ACCOUNT")

	baseCfg.APPLE_PASSWORD = os.Getenv("APPLE_PASSWORD")

}
