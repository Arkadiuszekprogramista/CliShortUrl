package shorter

import (
	"app/pkg/models"
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

func NewPool(host, port string) (*Redis, error) {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 3 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		},
	}

	temp := Redis{
		Pool: pool,
	}

	myRedis = &temp

	return myRedis, nil
}

func (r *Redis) AddShortUrlToRedis(addr *url.URL) error {
	conn := r.Pool.Get()
	defer conn.Close()

	u, err := Encode(addr)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", addr.String(), u)
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	return nil
}

func (r *Redis) LoadDataFromRedis(key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	get, err := redis.Values(conn.Do("HGETALL", key))

	if err == redis.ErrNil {
		return get, err
	}

	if err != nil {
		return get, err

	} else {
		return get, nil
	}

}

func (r *Redis) Close() {
	defer r.Close()
}

func (r *Redis) PrintAllKeys() ([]string, error) {

	keys, err := redis.Strings(r.Pool.Get().Do("KEYS", "*"))

	if len(keys) >= 100 {
		log.Println("Loaded more then 100 records, Please use different command")

	} else {

		for _, key := range keys {
			log.Println(key)
		}
	}

	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (r *Redis) AddUserToRedis(u *models.Result) error {
	conn := r.Pool.Get()

	defer conn.Close()

	_, err := r.Pool.Get().Do("HMSET", u.Email,"firstname", u.Name.First, "lastname", u.Name.Last, "email", u.Email)
	log.Printf("adding %s %s %s user to redis", u.Name.First, u.Name.Last, u.Email)
	if err != nil {
		return err
	}



	return nil

}