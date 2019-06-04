package setting

import "github.com/noahzaozao/go_nwlib/conf/database"
import "github.com/noahzaozao/go_nwlib/conf/cache"
import "github.com/noahzaozao/go_nwlib/conf/oss"

type SettingsConfig struct {
	DEBUG          string              `yaml:"DEBUG"`
	DefaultCharset string              `yaml:"DEFAULT_CHARSET"`
	SecretKey      string              `yaml:"SECRET_KEY"`
	OSS      	   oss.OSSConfig       `yaml:"OSS"`
	CACHES         []cache.RedisConfig `yaml:"CACHES"`
	DATABASES      []database.DBConfig `yaml:"DATABASES"`
}
