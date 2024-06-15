package cache

import (
	"check/maindb"
	"github.com/gomodule/redigo/redis"
)

func Chack() {
	var tyg redis.Conn
	redis.NewPool(func() (redis.Conn, error) {
		return tyg, nil
	}, 10)
	maindb.Check()
}
