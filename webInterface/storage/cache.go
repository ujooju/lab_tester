package storage

import "log"

var TokenCache *memoryCahce

type memoryCahce struct {
	data map[string]string
}

func InitCache() {
	log.Println("initialazing cache...")
	TokenCache = &memoryCahce{
		data: make(map[string]string),
	}
	log.Println("cache initialized")
}

func (mc *memoryCahce) Set(key string, value string) {
	mc.Maintain()
	mc.data[key] = value
}

func (mc *memoryCahce) Get(key string) (string, bool) {
	res, ok := mc.data[key]
	return res, ok
}

func (mc *memoryCahce) Has(key string) bool {
	_, ok := mc.data[key]
	return ok
}

func (mc *memoryCahce) Delete(key string) {
	delete(mc.data, key)
}

func (mc *memoryCahce) Maintain() {
	if len(mc.data) > 300 {
		for k := range mc.data {
			delete(mc.data, k)
		}
	}
	log.Println("cache cleaned")
}
