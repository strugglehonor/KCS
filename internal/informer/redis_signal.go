package informer

import (
	"context"
	"sync"

	"github.com/strugglehonor/KCS/internal/apiserver/service/v1"
	redisConf "github.com/strugglehonor/KCS/internal/config/redis"
	"github.com/strugglehonor/KCS/pkg/log"
	"github.com/strugglehonor/KCS/pkg/redis"
)

type Loader interface {
	PodReload() error
}

type Load struct {
	lock *sync.RWMutex
	ctx context.Context
	load Loader
	redisCli *redis.RedisCli
}

func (l *Load) Start() {
	go startPubSubLoop(l)
}
 
func startPubSubLoop(load *Load) {
	for {
		data, err := load.redisCli.Subscribe(redisConf.RedisChannal, int64(redisConf.ChannalTimeOut))
		if err != nil {
			log.FATAL.Printf("redis Subscribe %s occur error: %v\n", redisConf.RedisChannal, err)
		}
		
		switch string(data) {
		case v1.DeploymentChangedNotification:

		case v1.ClusterChangedNotification:

		case v1.PodChangedNotification:
			if err := load.load.PodReload(); err != nil {
				log.FATAL.Printf("Error occur when Reload Pod data: %s", err.Error())
			}
		}
	}
}