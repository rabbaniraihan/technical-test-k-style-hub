package database

import (
	"k-style-test/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func (runDB *DB) StartDB() {
	var err error
	runDB.db, err = gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang-test"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDb, err := runDB.db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	runDB.db.AutoMigrate(model.Member{}, model.Product{}, model.Review_product{}, model.Like_review{})
}

func (getDB *DB) GetDB() *gorm.DB {
	return getDB.db
}
