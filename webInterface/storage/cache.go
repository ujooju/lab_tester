package storage

var Cache map[string]CacheEntity = make(map[string]CacheEntity)

type CacheEntity struct {
	Token string
}
