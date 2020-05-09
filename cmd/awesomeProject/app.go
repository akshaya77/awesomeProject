package main

import (
	"awesomeProject/internal/delivery/api"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/repository/cache/redis"
	v1 "awesomeProject/internal/usecase/v1"
	"github.com/tokopedia/grace"
	"log"
	"time"
)

func main() {

	dbLayer := repository.New()

	redis.RedisInit()

	uApi := v1.GetApi(dbLayer)

	api.New(uApi)

	g := GraceCfg{}

	log.Fatal(grace.ServeWithConfig(":"+"9000", g.ToGraceConfig(), nil))

}

type GraceCfg struct {
	Timeout          string
	HTTPReadTimeout  string
	HTTPWriteTimeout string
}

func (g GraceCfg) ToGraceConfig() grace.Config {
	timeout, err := time.ParseDuration(g.Timeout)
	if err != nil {
		timeout = time.Second * 5
	}

	readTimeout, err := time.ParseDuration(g.HTTPReadTimeout)
	if err != nil {
		readTimeout = time.Second * 10
	}

	writeTimeout, err := time.ParseDuration(g.HTTPWriteTimeout)
	if err != nil {
		writeTimeout = time.Second * 10
	}

	return grace.Config{
		Timeout:          timeout,
		HTTPReadTimeout:  readTimeout,
		HTTPWriteTimeout: writeTimeout,
	}
}

type Redis map[string]struct {
	Address string
}

