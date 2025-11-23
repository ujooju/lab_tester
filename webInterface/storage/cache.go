package storage

import (
	"log"
	"sync"
)

var TokenCache *memoryCahce

type memoryCahce struct {
	data  map[string]string
	mutex sync.RWMutex
}

func InitCache() {
	log.Println("initialazing cache...")
	TokenCache = &memoryCahce{
		data:  make(map[string]string),
		mutex: sync.RWMutex{},
	}
	log.Println("cache initialized")
}

func (mc *memoryCahce) Set(key string, value string) {
	mc.Maintain()
	mc.mutex.Lock()
	mc.data[key] = value
	mc.mutex.Unlock()
}

func (mc *memoryCahce) Get(key string) (string, bool) {
	mc.mutex.RLock()
	res, ok := mc.data[key]
	mc.mutex.RUnlock()
	return res, ok
}

func (mc *memoryCahce) Has(key string) bool {
	mc.mutex.RLock()
	_, ok := mc.data[key]
	mc.mutex.RUnlock()
	return ok
}

func (mc *memoryCahce) Delete(key string) {
	mc.mutex.Lock()
	delete(mc.data, key)
	mc.mutex.Unlock()
}

func (mc *memoryCahce) Maintain() {
	mc.mutex.Lock()
	if len(mc.data) > 300 {
		for k := range mc.data {
			delete(mc.data, k)
		}
	}
	mc.mutex.Unlock()
	log.Println("cache cleaned")
}
