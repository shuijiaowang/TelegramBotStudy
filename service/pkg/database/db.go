package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func InitDB() error {
	// 简化配置（实际项目中应从环境变量读取）
	cfg := Config{
		Host:     "mysql2.sqlpub.com",
		Port:     "3307",
		User:     "shuijiaowang",
		Password: "uOsGjEWCXerZhWVC",
		DBName:   "shuijiaowang",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return err
	}

	return nil
}
