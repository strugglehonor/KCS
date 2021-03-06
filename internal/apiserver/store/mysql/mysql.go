package mysql

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/strugglehonor/KCS/internal/apiserver/store"
	"github.com/strugglehonor/KCS/internal/config/db"
	"github.com/strugglehonor/KCS/internal/model"
	"github.com/strugglehonor/KCS/pkg/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormLogger "gorm.io/gorm/logger"
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

func (ms *mysqlstore) Deployment() store.DeploymentStore {
	return newDeployment(ms)
}
 
var (
	once sync.Once
	mysqlFactory *mysqlstore
)

// mysql log
type databaseLogger struct {
	logger logrus.Logger
}

const logComponent = "logComponent"

func  (l *databaseLogger) doLog() *logrus.Entry {
	return l.logger.WithField(logComponent, "Database")
}

func (l *databaseLogger) LogMode(level ormLogger.LogLevel) ormLogger.Interface {
	levelMap := map[ormLogger.LogLevel]logrus.Level{
		ormLogger.Info: logrus.InfoLevel,
		ormLogger.Warn: logrus.WarnLevel,
		ormLogger.Error: logrus.ErrorLevel,
	}

	transLevel, ok := levelMap[level]
	if ok {
		l.logger.SetLevel(transLevel)
	}else {
		l.logger.SetOutput(ioutil.Discard)
	}

	return l 
}

func (l *databaseLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.doLog().Info(msg)
}

func (l *databaseLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.doLog().Warn(msg)
}

func (l *databaseLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.doLog().Error(msg)
}

func (l *databaseLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	l.doLog().WithField("MatchRows", rows).WithField("MilliSecondTakes", elapsed.Milliseconds()).Info(sql)
}

// new db logger
func NewDBlogger(Rawlogger logrus.Logger, level logrus.Level) *databaseLogger {
	l := &databaseLogger{
		logger: Rawlogger,
	}
	l.logger.SetLevel(level)
	return l
}

func GetMySQLFactoryOr() (*mysqlstore, error) {
	var dbConn *gorm.DB
	var err error
	log.INFO.Println("begin to connext mysql")
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s", db.User, db.Passwd, db.Host, db.Port, db.DBName, db.Charset)
		dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "t_", // ????????????
				SingularTable: true, // ??????????????????
			},
			Logger: NewDBlogger(*logrus.New(), logrus.DebugLevel),
		})
		mysqlFactory = &mysqlstore{db: dbConn}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to init db, failed to get mysqlStore, error message: %v", err)
	}

	log.INFO.Println("connect to mysql successfully")
	return mysqlFactory, nil
} 

func AutoMigrate(mysqlstore *mysqlstore) error {
	// err := mysqlstore.db.AutoMigrate(&model.Deployment{}, &model.Node{}, &model.Volume{}, &model.Cluster{}, &model.Pod{})
	// err := mysqlstore.db.Migrator().DropTable(&model.Deployment{}, &model.Node{}, &model.Volume{}, &model.Cluster{}, &model.Pod{})
	// if err != nil {
	// 	return err
	// }

	err := mysqlstore.db.AutoMigrate(&model.Deployment{}, &model.Node{}, &model.Volume{}, &model.Cluster{}, &model.Pod{})
	return err
}