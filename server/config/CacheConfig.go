package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/spf13/viper"
)

var cacheOnce sync.Once

var cacheChain *cache.ChainCache

func GetCache() *cache.ChainCache {

	// only run init of dbConnection once. this should be thred safe
	cacheOnce.Do(func() {
		var err error
		var caches []cache.SetterCacheInterface

		// ristretto is always used as first level of cache
		ristrettoCache, err := ristretto.NewCache(&ristretto.Config{NumCounters: 1000, MaxCost: 100, BufferItems: 64})
		if err != nil {
			panic(err)
		}
		ristrettoStore := store.NewRistretto(ristrettoCache, &store.Options{Expiration: 30 * time.Second})
		caches = append(caches, cache.New(ristrettoStore))

		// add memcache if configuared
		if viper.GetViper().IsSet("cache.memcache") {

			fmt.Println("Memcache cache is set")
			memcacheStore := store.NewMemcache(
				memcache.New(fmt.Sprintf("%s:%s", viper.GetString("cache.memcache.host"), viper.GetString("cache.memcache.port"))),
				&store.Options{
					Expiration: time.Duration(viper.GetInt("cache.memcache.expiration_time")) * time.Second,
				},
			)

			caches = append(caches, cache.New(memcacheStore))
		}

		// add redis if configuared
		// TODO: redisStore seems not to have the correct signitaure of Clear method
		// if viper.GetViper().IsSet("cache.redis") {
		// 	fmt.Println("Redis cache is set")
		// 	redisClient := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", viper.GetString("cache.redis.host"), viper.GetString("cache.redis.port"))})
		// 	redisStore := store.NewRedis(redisClient, &store.Options{Expiration: 5 * time.Second})
		// 	caches = append(caches, cache.New(redisStore))
		// }

		// add all configuared caches to chain
		cacheChain = cache.NewChain(
			caches...,
		)
	})

	return cacheChain
}
