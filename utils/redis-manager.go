package utils

import (
	"berlin/lib"
	"fmt"

	log "github.com/fatih/color"
	"github.com/go-redis/redis"
)

// RedisInstance :  RedisInstance for redis client
type RedisInstance struct {
	RInstance *redis.Client
}

// RedisPort : Declare global variable
var RedisPort string

func init() {
	configuration, err := lib.LoadConfig()
	if err != nil {
		RedisPort = "localhost:6379"
	} else {
		RedisPort = configuration.Keys["project"]["REDIS_PORT"]
	}
}

// InitRedisClient : InitRedisClient
func InitRedisClient() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisPort,
		Password: "",
		DB:       0,
	})
	pingResponse := Ping(client)
	if pingResponse["status"] != true {
		fmt.Println(pingResponse["message"])
	}
	return *client
}

// Ping : Ping
func Ping(redClient *redis.Client) map[string]interface{} {
	pong, err := redClient.Ping().Result()
	if err != nil {
		log.Blue("Cannot Initialize Redis Client  %s", err)
		return Message(false, "Cannot Initialize Redis Client")
	}
	log.Blue("Redis Client Successfully Initialized . . . %s", pong)
	return Message(true, "Redis Client Successfully Initialized . . ."+pong)
}
