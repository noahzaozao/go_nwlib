package utils

import (
	goconfig "github.com/micro/go-config"
	"log"

	setting_conf "github.com/noahzaozao/go_nwlib/conf/setting"
)


func LoadConfig() *setting_conf.SettingsConfig {
	// Load json config file
	if err := goconfig.LoadFile("./config.yaml"); err != nil {
		log.Println(err.Error())
		return nil
	}
	return GetConfig()
}

func GetConfig() *setting_conf.SettingsConfig {
	var settingsConfig setting_conf.SettingsConfig
	if err := goconfig.Get("config").Scan(&settingsConfig); err != nil {
		log.Println(err.Error())
		return nil
	}
	return &settingsConfig
}

func CheckSettingConfig(settingsConfig *setting_conf.SettingsConfig) {
	log.Println("DEBUG: " + settingsConfig.DEBUG)
	log.Println("CHARSET: " + settingsConfig.DefaultCharset)

	if len(settingsConfig.DATABASES) < 1 {
		log.Println("DATABASE config not exist")
		return
	}

	if len(settingsConfig.CACHES) < 1 {
		log.Println("CACHES config not exist")
		return
	}
}