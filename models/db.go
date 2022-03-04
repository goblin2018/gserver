package models

import (
	"database/sql"
	"fmt"
	"gserver/pkg/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	var err error
	// This is in dev mode
	// deploy should use ssl
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		conf.DB.Host,
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Name,
		conf.DB.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		panic(fmt.Errorf("can't connect to db: %v", err))
	}

	var sqlDb *sql.DB
	sqlDb, err = db.DB()
	sqlDb.SetMaxOpenConns(conf.DB.MaxOpenConns)
	sqlDb.SetMaxIdleConns(conf.DB.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(conf.DB.ConnMaxLifetime)
	migrate()
}

func migrate() {
	err := db.AutoMigrate(
		//&User{}
	)

	if err != nil {
		panic(fmt.Errorf("migrate error: %v", err))
	}
}
