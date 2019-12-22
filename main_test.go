package main

import (
	api "berlin/internal/api"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

// newTestRedis returns a redis.Cmdable.
func newTestRedis() redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return *client
}

func TestServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}
	client := newTestRedis()
	redisHandlerClient := &api.RedisInstance{RInstance: &client}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redisHandlerClient.HealthCheck)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(rr.Body.String()), &result); err != nil {
		panic(err)
	}
	if result["status"] != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), result["status"])
	}
}
