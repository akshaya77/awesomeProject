package redis

import (
	"errors"
	"fmt"
	"gopkg.in/redis.v5"
	"time"
)

type IRedisConn interface {
	GetConnection(label string) IRedisWrapper
}

// RedisConn stores map of RedisWrapper.
type RedisConn struct {
	Connections map[string]*RedisWrapper
}

// InitRedis from configuration map.
func InitRedis(address string) *RedisConn {
	var Conn RedisConn
	Conn.Connections = make(map[string]*RedisWrapper, 1)
	Conn.Connections["my-redis"] = NewRedisConn(address)
	return &Conn
}


func NewRedisConn(connection string) *RedisWrapper {
	rds := redis.NewClient(&redis.Options{
		Addr:         connection,
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     500,
		MaxRetries:   2,
		DialTimeout:  time.Millisecond * 500,
		ReadTimeout:  time.Millisecond * 500,
		WriteTimeout: time.Millisecond * 500,
	})

	return &RedisWrapper{
		addr: connection,
		conn: rds,
	}
}

func (redisConn *RedisConn) PingRedis() error {
	for _, val := range redisConn.Connections {
		_, err := val.Connection().Ping().Result()
		if err != nil {
			go func(err error, val *RedisWrapper) {
				for {
					if err != nil {
						fmt.Printf("[Ping Redis] Reconnecting to redis for address [%s] err [%v]\n", val.Address(), err)
						val.Reconnect()
						_, err = val.Connection().Ping().Result()
						time.Sleep(1 * time.Second)
						continue
					}
					break
				}
			}(err, val)
		}
	}
	return nil
}


// GetConnection seek whether this Redis object holds connection with label `connection_name`.
// Using this function is encouraged because it would print warning if such connection doesn't exist.
// Hint: connection label is the string after 'redis-' in configuration files.
func (redisConn *RedisConn) GetConnection(label string) *RedisWrapper {
	conn := redisConn.Connections[label]
	if conn == nil {
		fmt.Printf("Unable to find connection with label %s\n", label)
		return nil
	}
	return conn
}

// Convert standard Redis connection into RedisConn.
func Convert(connections map[string]*RedisWrapper) (*RedisConn, error) {
	var res RedisConn
	res.Connections = connections
	names := []string{
		"tp-object", "user-session", "recently-viewed",
		"product-stats", "stats", "12-3", "22-6", "installment", "redis-merchant",
		"sitemap",
	}
	for _, name := range names {
		con := res.GetConnection(name).Connection()
		if con == nil {
			return nil, fmt.Errorf("Unable to find connection path to Redis %s", name)
		}
	}
	return &res, nil
}

// InitMock for mocking redis in unit test
func InitMock(address string, list []string) (*RedisConn, error) {
	if len(list) == 0 {
		err := errors.New("no connection list given")
		return nil, err
	}

	var Conn RedisConn
	Conn.Connections = make(map[string]*RedisWrapper)

	for _, v := range list {
		Conn.Connections[v] = NewRedisConn(address)
	}

	return &Conn, nil
}
