package shorter

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Pool *redis.Pool
}

func NewPool(host, port string) (Service, error) {
	pool := &redis.Pool{
		MaxIdle: 10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			 return redis.Dial("tcp",fmt.Sprintf("%s:%s",host, port))
			},
	}
	
	return &Redis{pool}, nil
}

func(r *Redis) AddShortUrlToRedis(url link) (string, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	add, err := conn.Do("SET", url.addr, url.Encode())
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(add)
	return "Added", nil

}

func(r *Redis) LoadDataFromRedis(key string) (string, error){
	conn := r.Pool.Get()
	defer conn.Close()

	get, err := conn.Do("GET", key)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println(get)
	return fmt.Sprintln(get), nil
}


func(r *Redis) Close() error {
	return r.Pool.Close()
}