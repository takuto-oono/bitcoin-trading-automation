package mysql

import (
	"fmt"

	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bitcoin-trading-automation/internal/config"
)

// gormを使ってmysqlに接続する

type MYSQL struct {
	Config config.Config
	DB     *gorm.DB
}

func NewMYSQL(cfg config.Config) (*MYSQL, error) {
	db, err := connectMYSQL(cfg)
	if err != nil {
		return nil, err
	}

	// マイグレーションの管理がまだでkていないので、ここで自動マイグレーションする
	if err := db.AutoMigrate(&Ticker{}); err != nil {
		return nil, err
	}

	return &MYSQL{
		Config: cfg,
		DB:     db,
	}, nil
}

func connectMYSQL(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.MYSQL.User, cfg.MYSQL.Password, cfg.MYSQL.Host, cfg.MYSQL.Port, cfg.MYSQL.DbName)

	return gorm.Open(gormMysql.Open(dsn), &gorm.Config{})
}
