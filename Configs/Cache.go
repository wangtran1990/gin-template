package configs

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// Cache ... Global cache variable
var Cache *redis.Client

// Ctx ... Global background context variable
var Ctx = context.Background()

// Config variables from .env
var status, addr, password string
var defaultDB int

// InitCache ... Caching initialization
func InitCache() error {
	status = os.Getenv("CACHE_ON")
	addr = os.Getenv("CACHE_ADDR")
	password = os.Getenv("CACHE_PASSWORD")
	defaultDB, _ = strconv.Atoi(os.Getenv("CACHE_DEFAULT_DB"))

	if status != "1" { // ON
		return errors.New("Cache off by setting")
	}

	Cache = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,  // no password set
		DB:       defaultDB, // use default DB
	})

	_, err := Cache.Ping(Ctx).Result()
	return err
}
