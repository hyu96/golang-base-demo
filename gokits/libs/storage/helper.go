package storage

import (
	"fmt"
	"sync"
)

func LoadClient[T any](dbName string, mapDao, sourceClient *sync.Map, sfunc func(interface{}) T) T {
	db, ok := mapDao.Load(dbName)
	if ok {
		if db == nil {
			panic(fmt.Errorf("LoadDAO: DB %s returned nil", dbName))
		}
		return db.(T)
	}

	sourceDB, ok := sourceClient.Load(dbName)
	if !ok || sourceDB == nil {
		panic(fmt.Errorf("LoadDAO: DB %s does not exist, ok:%t", dbName, ok))
	}

	db = sfunc(sourceDB)
	mapDao.Store(dbName, db)

	return db.(T)
}
