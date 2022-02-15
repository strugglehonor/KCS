package main

import (
	"github.com/gin-gonic/gin"
	"github.com/strugglehonor/KCS/internal/apiserver/controller/grpc"
	"github.com/strugglehonor/KCS/internal/apiserver/controller/restful"
	r "github.com/strugglehonor/KCS/internal/apiserver/store/redis"
	"github.com/strugglehonor/KCS/pkg/log"
	"github.com/strugglehonor/KCS/pkg/redis"
)

var (
	redisCli  *redis.RedisCli
)

func main() {
	log.INFO.Println("starting apiserver")

	redisCli = r.InitRedis()

	restful.InitRouter(gin.Default())  // gin  这里不能传入一个空的 gin.Engine{}

	if err := grpc.InitGrpc(); err != nil {  // grpc server
		panic(err)
	}

	log.INFO.Println("apiserver start successfully")
}