package helper

import (
	"time"
	configs "template/Configs"
)

// CacheGetString ...
func CacheGetString(key string) (string, error) {
	val, err := configs.Cache.Get(configs.Ctx, key).Result()
	return val, err
}

// CacheSetString ...
func CacheSetString(key string, value string, ttl time.Duration) error {
	ttl = ttl * time.Second
	err := configs.Cache.Set(configs.Ctx, key, value, ttl).Err()
	return err
}

// CacheDelele ...
// Removes the specified keys. A key is ignored if it does not exist
// Return the number of keys that were removed.
func CacheDelele(key string) int64 {
	val := configs.Cache.Del(configs.Ctx, key).Val()
	return val
}

// CacheExists ...
// Returns if key exists
func CacheExists(key string) bool {
	val := configs.Cache.Exists(configs.Ctx, key).Val()
	return val > 0
}
