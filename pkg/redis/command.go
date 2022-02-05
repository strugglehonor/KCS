package redis

import (
	"strconv"

	"github.com/gomodule/redigo/redis"
)

//Append 实现 APPEND 命令
func (rc *RedisCli) Append(key, value string) error {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	_, err := conn.Do("APPEND", key, value)
	return err
}

//Decr 实现DECR key
func (rc *RedisCli) Decr(key string) (int64, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	result, err := redis.Int64(conn.Do("DECR", key))
	return result, err
}

//Decrby 实现DECRBY
func (rc *RedisCli) DecrBy(key string, value int) (int64, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	result, err := redis.Int64(conn.Do("DECRBY", key, strconv.Itoa(value)))
	return result, err
}

//Get 实现GET
func (rc *RedisCli) Get(key string) (string, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	return redis.String(conn.Do("GET", key))
}

func (rc *RedisCli) Incr(key string) (int64, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	result, err := redis.Int64(conn.Do("INCR", key))
	return result, err
}

func (rc *RedisCli) IncrBy(key string, value int) (int64, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	result, err := redis.Int64(conn.Do("INCRBY", key, strconv.Itoa(value)))
	return result, err
}

func (rc *RedisCli) IncrbyFloat(key string, value float64) (string, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	floatStr := strconv.FormatFloat(value, 'f', 6, 32)
	return redis.String(conn.Do("INCRBYFLOAT", key, floatStr))
}

func (rc *RedisCli) Set(key, value string) (int64, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	result, err := conn.Do("SET", key, value)
	if result == "OK" {
		return 1, nil
	}
	return 0, err
}

func (rc *RedisCli) SetEx(key, value string, second int32) (int64, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	result, err := conn.Do("SETEX", key, second, value)
	if result == "OK" {
		return 1, nil
	}
	return 0, err
}

func (rc *RedisCli) SetNX(key, value string) (int64, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	return redis.Int64(conn.Do("SETNX", key, value))
}

/// key 操作.
func (rc *RedisCli) Del(key string) error {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	_, err := conn.Do("DEL", key)
	return err
}

func (rc *RedisCli) Exists(key string) (bool, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	result, err := redis.Int64(conn.Do("EXISTS", key))
	return result == 1, err
}

func (rc *RedisCli) Expire(key string, second int32) error {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	_, err := conn.Do("EXPIRE", key, second)
	return err
}

func (rc *RedisCli) ExpireAt(key string, timestamp int32) error {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	_, err := conn.Do("EXPIREAT", key, timestamp)
	return err
}

func (rc *RedisCli) Multi() error {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	_, err := conn.Do("MULTI")
	return err
}

func (rc *RedisCli) Exec() error {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	_, err := conn.Do("EXEC")
	return err
}

// 执行 lua 脚本.
func (rc *RedisCli) Eval(scripts string, keys []string, argvs []string) (string, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	return redis.String(conn.Do("EVAL", redis.Args{}.Add(scripts).Add(len(keys)).AddFlat(keys).AddFlat(len(argvs)).AddFlat(argvs)...))
}

// keys
func (rc *RedisCli) Keys(pattern string) ([]string, error) {
	conn, pool := rc.GetConn()
	defer func() {
		if pool {
			conn.Close()
		}
	}()
	return redis.Strings(conn.Do("keys", pattern))
}
