package db_package

import (
	log_package "github.com/aminwhat/golang-gateway/packages/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbConfig struct {
	DbName string
	Db     *gorm.DB
}

func FromGromDB(db *gorm.DB) (dbConfig *DbConfig) {
	dbConfig.Db = db
	dbConfig.DbName = db.Name()

	return
}

func (dbConfig *DbConfig) FromGromDB(db *gorm.DB) {
	dbConfig.Db = db
	dbConfig.DbName = db.Name()
}

func (dbConfig *DbConfig) Open() (*gorm.DB, error) {
	if dbConfig.Db != nil {
		log_package.Warning("Try to open an already opened Database: " + dbConfig.DbName)
		return dbConfig.Db, nil
	}

	db, err := gorm.Open(sqlite.Open(dbConfig.DbName+".db"), &gorm.Config{})
	if err != nil {
		log_package.Error("Failed to Open Sqlite Gorm db")
		return nil, err
	}

	dbConfig.Db = db

	return db, nil
}

func (dbConfig *DbConfig) MigrateChanges() error {
	return dbConfig.Db.AutoMigrate()
}
