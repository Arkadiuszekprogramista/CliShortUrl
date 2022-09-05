package shorter


type Service interface{
	AddShortUrlToRedis(url link)(string, error)
	LoadDataFromRedis(key string) (string, error)
	Close() error
}
