package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	cache_conf "github.com/noahzaozao/go_nwlib/conf/cache"
	"log"
	"sync"
)

type CacheManager struct {
	config cache_conf.RedisConfig
}

var instance *CacheManager
var once sync.Once

func CacheMgr() *CacheManager {
	once.Do(func () {
		instance = &CacheManager{}
	})
	return instance
}

//
// 初始化缓存配置文件
//
func (cacheMgr *CacheManager) Init(cacheConfig cache_conf.RedisConfig) error {
	cacheMgr.config = cacheConfig
	if cacheMgr.config.Type == "redis" {
		dbConn, err := cacheMgr.Conn()
		if err != nil {
			return err
		}
		defer dbConn.Close()
		log.Println("Cache connected")
	} else {
		log.Println("Cache Type is incorrect")
	}
	return nil
}

//
// 获取缓存连接
//
func (cacheMgr *CacheManager) Conn() (*redis.Client, error) {
	connStr := fmt.Sprintf(
		"%s:%s",
		cacheMgr.config.Host,
		cacheMgr.config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     connStr,
		Password: cacheMgr.config.Password, // no password set
		DB:       cacheMgr.config.DB,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
