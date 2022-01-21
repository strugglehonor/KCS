package mysql

import (
	"fmt"
	"sync"

	"github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/internal/config/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type mysqlstore struct {
	db *gorm.DB
}

func (ms *mysqlstore) Cluster() store.ClusterStore {
	return newCluster(ms)
}

func (ms *mysqlstore) Pod() store.PodStore {
	return newPod(ms)
}

func (ms *mysqlstore) Volume() store.VolumeStore {
	return newVolume(ms)
}

var (
	once sync.Once
	mysqlFactory store.Factory
)

func GetMySQLFactoryOr() (store.Factory, error) {
	var dbConn *gorm.DB
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("gorm:gorm@tcp(%s:%v)/gorm?charset=%s", db.Host, db.Port, db.Charset)
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "t_", // 表名前缀
				SingularTable: true, // 使用单数表名
			},
		})
		mysqlFactory = &mysqlstore{db: dbConn}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to init db, failed to get mysqlStore, error message: %v", err)
	}

	return mysqlFactory, nil
} 