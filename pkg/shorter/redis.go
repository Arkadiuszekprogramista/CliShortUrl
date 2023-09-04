package shorter

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Pool *redis.Pool
}

var myRedis *Redis

func NewRedis(r *Redis) Redis {
	r = myRedis
	return *r
}

func NewPool(host, port string) (*Redis, error) {
	pool := &redis.Pool{
		MaxIdle: 10,
		IdleTimeout: 3 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",fmt.Sprintf("%s:%s",host, port))
			},
	}

	temp := Redis{
		Pool: pool,
	}

	myRedis = &temp

	return myRedis, nil
}

func(r *Redis) AddShortUrlToRedis(url link) error {
	conn := r.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", url.addr, url.Encode())
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	
	return nil

}

func(r *Redis) LoadDataFromRedis(key string) (string, error){
	conn := r.Pool.Get()
	defer conn.Close()

	get, err := conn.Do("GET", key)
	if err != nil {
		fmt.Println("Error", err)
		return "", err
	}
	return fmt.Sprintln(get), nil
}


func(r *Redis) Close() {
	defer r.Close()
}


func(r *Redis) PrintAll() error {


	keys, err := r.Pool.Get().Do("KEYS","*")
	if err != nil {
		return err
	}



	log.Println(keys)

	return nil
}