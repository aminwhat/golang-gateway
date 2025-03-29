package db_package

import "gorm.io/gorm"

var dbSingletons map[string]*gorm.DB = map[string]*gorm.DB{}

func (dbConfig *DbConfig) OpenSingleton(dbKey string) (*gorm.DB, error) {

	if dbConfig.Db == nil && dbSingletons[dbKey] != nil {
		dbConfig.Db = dbSingletons[dbKey]
	} else if dbConfig.Db != nil && dbSingletons[dbKey] == nil {
		dbSingletons[dbKey] = dbConfig.Db
	} else if dbConfig.Db != nil && dbSingletons[dbKey] != nil {
		dbSingletons[dbKey] = dbConfig.Db
	}

	db, err := dbConfig.Open()

	if err == nil {
		dbSingletons[dbKey] = db
	}

	return db, err
}

// func GetSingletonDB(dbKey string)
