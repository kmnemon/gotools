package service

import (
	"simpleserver/simpledb"
)

func FindByPrimaryKeyService(key string) ([]byte, bool) {
	return simpledb.FindByPrimaryKey("./simpledb/db2", key)
}
