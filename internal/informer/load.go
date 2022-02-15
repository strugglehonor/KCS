package informer

import (
	"encoding/json"

	"github.com/AlekSi/pointer"
	pb "github.com/strugglehonor/KCS/internal/apiserver/controller/grpc/proto/v1"
	"github.com/strugglehonor/KCS/pkg/log"
)

func (load Load) PodReload() error {
	log.INFO.Println("Pod Reload begin!")

	c, err := GetGrpcConn()
	if err != nil {
		log.FATAL.Printf("Error occur when called GetGrpcConn(): %s\n", err.Error())
		
	}

	req := &pb.ListPodsRequest{
		Offset: pointer.ToInt64(0),
		Limit: pointer.ToInt64(-1),
	}

	cacheClient := pb.NewCacheClient(c.client)
	resp, err := cacheClient.ListPods(load.ctx, req)
	if err != nil {
		return err
	}

	log.INFO.Printf("Pod Found %d total\n", len(resp.Items))

	for _, podinfo := range resp.Items {
		podinfoStr, err := json.Marshal(podinfo)
		if err != nil {
			log.FATAL.Printf("podinfo marshal failed which name is %s, error messages: %s", podinfo.Name, err.Error())
			return err
		}
		
		load.redisCli.Set(podinfo.Name, string(podinfoStr))
	}

	return nil
}