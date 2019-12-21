package utils

import (
	"berlin/lib"
	"fmt"

	"github.com/go-redis/redis"
)

// RedisInstance :  RedisInstance for redis client
type RedisInstance struct {
	RInstance *redis.Client
}

// RedisPort : Declare global variable
var RedisPort string

func init() {
	configuration, _ := lib.LoadConfig()
	RedisPort = configuration.Keys["project"]["REDIS_PORT"]
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
		fmt.Println("Cannot Initialize Redis Client ", err)
		return Message(false, "Cannot Initialize Redis Client")
	}
	return Message(false, "Redis Client Successfully Initialized . . ."+pong)
}
