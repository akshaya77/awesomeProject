package repository

import "awesomeProject/internal/repository/cache/redis"

type DBConnection struct {
	redis redis.RedisDb
}

func New() DBRepository {
	dbConnection := &DBConnection{
		redis: redis.RedisDb{},
	}
	return dbConnection
}