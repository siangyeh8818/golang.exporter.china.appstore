package config

import "os"

type BaseConfig struct {
	CRAWLER_INTERNAL_TIME string
	APPLE_ACCOUNT         string
	APPLE_PASSWORD        string
}

func (baseCfg *BaseConfig) InitConfig(configPath string) {

	baseCfg.CRAWLER_INTERNAL_TIME, err = os.Getenv("CRAWLER_INTERNAL_TIME")

	baseCfg.APPLE_ACCOUNT, err = os.Getenv("APPLE_ACCOUNT")

	baseCfg.APPLE_PASSWORD, err = os.Getenv("APPLE_PASSWORD")

}
