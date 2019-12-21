package api

import "github.com/go-redis/redis"

// RedisInstance :  RedisInstance
type RedisInstance struct {
	RInstance *redis.Client
}

// User is a simple user struct for this example
type User struct {
	Username string `json:"username"`
	Item     string `json:"item"`
	Amount   int    `json:"amount"`
}
