package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	database_conf "github.com/noahzaozao/go_nwlib/conf/database"
	"github.com/noahzaozao/go_nwlib/models"
	"log"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
	// import _ "github.com/jinzhu/gorm/dialects/sqlite"
	// import _ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/jinzhu/gorm"
	"sync"
)

type DatabaseManager struct {
	config database_conf.DBConfig
}

var instance *DatabaseManager
var once sync.Once

func DBManager() *DatabaseManager {
	once.Do(func () {
		instance = &DatabaseManager{}
	})
	return instance
}

//
// 初始化数据库配置文件
//
func (dbMgr *DatabaseManager) Init(databaseConfig database_conf.DBConfig) error {
	dbMgr.config = databaseConfig
	if dbMgr.config.Type == "mysql" {
		dbConn, err := dbMgr.Conn()
		if err != nil {
			return err
		}
		dbConn.AutoMigrate(&models.User{})
		dbConn.AutoMigrate(&models.WechatMappUser{})
		defer dbConn.Close()
		log.Println("Database connected")
	} else {
		log.Println("Database Type is incorrect")
	}
	return nil
}

//
// 获取数据库连接
//
func (dbMgr *DatabaseManager) Conn() (*gorm.DB, error) {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		dbMgr.config.User,
		dbMgr.config.Password,
		dbMgr.config.Host,
		dbMgr.config.Port,
		dbMgr.config.Database)
	dbConn, err := gorm.Open(dbMgr.config.Type, connStr)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
