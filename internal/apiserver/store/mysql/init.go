package mysql

import (
	"github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/pkg/log"
)

func init() {
	mysqlstore, err := GetMySQLFactoryOr()
	if err != nil {
		log.FATAL.Panic("get mysql factory failed: %v", err)
	}
	store.SetClient(mysqlstore)
	AutoMigrate(mysqlstore)
}