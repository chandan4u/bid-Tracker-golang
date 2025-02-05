package main

import (
	api "berlin/internal/api"
	lib "berlin/lib"
	u "berlin/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/fatih/color"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// RedisInstance :  RedisInstance for redis client
type RedisInstance struct {
	RInstance *redis.Client
}

// PORT : Declare global variable
var PORT string

func init() {
	// [ LoadConfig : Load configuration from config files ]
	configuration, err := lib.LoadConfig()
	if err != nil {
		PORT = "8000"
	} else {
		PORT = configuration.Keys["project"]["HTTP_PORT"]
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
	}
	// [ createRouter : creating route and injecting redis client in routes. ]
	router := createRouter()
	log.Blue("Start serve on PORT :: %s", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	start(server)
}

func createRouter() *mux.Router {

	// [ InitRedisClient : Initilizing redis server ]
	client := u.InitRedisClient()
	redisHandlerClient := &api.RedisInstance{RInstance: &client}
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", redisHandlerClient.HealthCheck).Methods("GET")

	// [ PathPrefix : Added path prefix for routes ]
	subrouter := router.PathPrefix("/api/berlin/internal/").Subrouter()
	subrouter.HandleFunc("/add-biding", redisHandlerClient.AddBiding).Methods("POST")
	subrouter.HandleFunc("/all-bids-by-user", redisHandlerClient.GetAllBidsByUser).Methods("POST")
	subrouter.HandleFunc("/all-bids-by-item", redisHandlerClient.GetAllBidsByItem).Methods("POST")
	subrouter.HandleFunc("/winner-by-item", redisHandlerClient.GetWinnerByItem).Methods("POST")

	// [ NotFoundHandler : If request route not found throw custom message ]
	router.NotFoundHandler = http.HandlerFunc(u.NotFoundHandler)
	return router
}

func start(server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()
	gracefulshutdown(server)
}

// gracefulshutdown : grace fully shutdown for exception case.
func gracefulshutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-stop

	log.Yellow("Shutting the server down.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := server.Shutdown(ctx); err != nil {
		log.Red("Something went's wrong in shutdown server :: &s", err)
	} else {
		log.Blue("Server Stop successfully")
	}
}
