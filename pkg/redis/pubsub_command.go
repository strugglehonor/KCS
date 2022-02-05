package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/strugglehonor/KCS/pkg/log"
)

func (rc *RedisCli) Public(channel, msg string) error {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	_, err := conn.Do("PUBLISH", channel, msg)
	return err
}

func (rc *RedisCli) Subscribe(channel string, timeout int64) ([]byte, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channel)

	for {
		switch n := psc.ReceiveWithTimeout(time.Duration(timeout) * time.Second).(type) {
		case error:
			return nil, fmt.Errorf("timeout")
		case redis.Message:
			return n.Data, nil
		}
	}
}

