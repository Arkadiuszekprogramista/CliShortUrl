package shorter

import (
	"fmt"
	"log"
	"net/url"
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

func AddrValidation(addr string) (*url.URL, error) {

	url, err := url.ParseRequestURI(addr)
	if err != nil {
		return url, err
	} else {
		return url, nil
	}
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


func(r *Redis) AddShortUrlToRedis(addr string) error {
	conn := r.Pool.Get()
	defer conn.Close()

	url, err := AddrValidation(addr)
	if err != nil {
		log.Printf("%s is not a URL", addr)
		return err

	} else {

		u, err := Encode(url.Path)
		if err != nil {
			return err
		}
	
		_, err = conn.Do("SET", url.Path, u)
		if err != nil {
			fmt.Println("Error", err)
			return err
		}
		
		return nil
	}

}


func(r *Redis) LoadDataFromRedis(key string) (string, error){
	conn := r.Pool.Get()
	defer conn.Close()

	get, err := redis.String(conn.Do("GET", key))

	if err == redis.ErrNil {
		return get, err
	}

	if err != nil {
		return get, err

	} else {
		return get, nil
	}

}


func(r *Redis) Close() {
	defer r.Close()
}


func(r *Redis) PrintAll() (interface{}, error) {

	var myKeys interface{}

	keys, err := redis.Strings(r.Pool.Get().Do("KEYS", "*"))

	if len(keys) >= 100 {
		log.Println("Loaded more then 100 records, Please use different command")
	} else {
		for _, key := range keys{
			log.Println(key)
		}
	}

	if err != nil {
		return nil ,err
	}
	
	myKeys = keys

	return myKeys, err
}
