package shorter

import "net/url"

type Service interface {
	AddShortUrlToRedis(u *url.URL) error
	LoadDataFromRedis(key string) (string, error)
	PrintAll() error
	Close() error
}
