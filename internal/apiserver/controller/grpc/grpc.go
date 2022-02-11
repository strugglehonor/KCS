package grpc

import (
	"net"

	"github.com/strugglehonor/KCS/internal/apiserver/controller/grpc/cache"
	pb "github.com/strugglehonor/KCS/internal/apiserver/controller/grpc/proto/v1"
	"github.com/strugglehonor/KCS/internal/apiserver/store/mysql"
	grpcConf "github.com/strugglehonor/KCS/internal/config/grpc"
	"github.com/strugglehonor/KCS/pkg/log"

	"google.golang.org/grpc"
)


func InitGrpc() error {
	log.INFO.Println("init grpc server")
	lis, err := net.Listen("tcp", grpcConf.Port)
	if err != nil {
		log.FATAL.Printf("failed to listen:%v\n", err)
		return err
	}
	
	s := grpc.NewServer()

	mysqlstore, err := mysql.GetMySQLFactoryOr()
	if err != nil {
		log.FATAL.Printf("called GetMySQLFactoryOr() failed: %v\n", err)
		return err
	}

	cacheIns, err := cache.GetCacheOr(mysqlstore)
	if err != nil {
		log.FATAL.Printf("occur error when call GetCacheOr(): %v\n", err)
		return err
	}

	pb.RegisterCacheServer(s, cacheIns)
	if err := s.Serve(lis); err != nil {
		log.FATAL.Printf("failed to serve: %v", err)
		return err
	}
	log.INFO.Println("init grpc successfully")
	return nil
}