package shorter


type Service interface{
	AddShortUrlToRedis(url link) error
	LoadDataFromRedis(key string) (string, error)
	PrintAll() error
	Close() error
}
