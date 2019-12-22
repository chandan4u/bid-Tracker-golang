package api_test

import (
	"berlin/internal/api"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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
	mr.HSet("daniel_bitcoin", `[{"username":"daniel","item":"bitcoin","amount":10}]`, "test2")
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	return *client
}

func TestHealthCheckAPI(t *testing.T) {
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

func TestAddBidingAPI(t *testing.T) {
	data := url.Values{}
	data.Set("username", "daniel")
	data.Set("item", "bitcoin")
	data.Set("amount", "100")
	req, err := http.NewRequest("POST", "/api/berlin/internal/add-biding", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := newTestRedis()
	redisHandlerClient := &api.RedisInstance{RInstance: &client}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redisHandlerClient.AddBiding)
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

func TestGetAllBidsByUserAPI(t *testing.T) {
	data := url.Values{}
	data.Set("username", "d")
	req, err := http.NewRequest("POST", "/api/berlin/internal/all-bids-by-user", strings.NewReader(data.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := newTestRedis()
	redisHandlerClient := &api.RedisInstance{RInstance: &client}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redisHandlerClient.GetAllBidsByUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(rr.Body.String()), &result); err != nil {
		panic(err)
	}
	if result["status"] != false {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), result["status"])
	}
	if result["message"] != "Username should greater than 1 character." {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), result["status"])
	}
}
