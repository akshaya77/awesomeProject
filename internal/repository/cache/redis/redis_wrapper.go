package redis

import (
	"gopkg.in/redis.v5"
	"net"
	"os"
	"sync"
	"syscall"
)


type IRedisWrapper interface {
	TryReconnect(err error) error
	Reconnect() error
	Connection() *redis.Client
	Address() string
}

// RedisWrapper implements IRedisWrapper with real Redis client.
type RedisWrapper struct {
	// We save connection parameters for use on reconnect.
	addr string
	conn *redis.Client
	// Be careful when copying RedisWrapper since it stores mutex by value.
	// Copied mutex would refer to different lock when locked.
	// Advice: pass RedisWrapper by pointer/reference
	mutex sync.Mutex
}



func NewRedisWrapper(address string, connection *redis.Client) *RedisWrapper {
	return &RedisWrapper{addr: address, conn: connection}
}


func (r *RedisWrapper) TryReconnect(err error) error {
	// The real syscall error is buried under various structs, so
	// we need to do some dance to get it out.
	if err != nil {
		switch e := err.(type) {
		case *net.OpError:
			switch n := e.Err.(type) {
			case *os.SyscallError:
				if n.Err == syscall.EPIPE {
					return r.Reconnect()
				}
			}
		}
	}
	return nil
}

// Reconnect to Redis server.
func (r *RedisWrapper) Reconnect() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.conn.Close()

	rds := redis.NewClient(&redis.Options{
		Addr:     r.addr,
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 500,
	})

	r.conn = rds
	return nil
}

// Connection to Redis client.
func (r *RedisWrapper) Connection() *redis.Client {
	return r.conn
}

// Address of this Redis.
func (r *RedisWrapper) Address() string {
	return r.addr
}