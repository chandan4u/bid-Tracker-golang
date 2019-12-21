package api

import (
	u "berlin/utils"
	"net/http"
)

/*
	[ http call for healthcheck ]
	[ POST ] http://localhost:8080/healthcheck
*/

// HealthCheck : HealthCheck
func (c *RedisInstance) HealthCheck(w http.ResponseWriter, r *http.Request) {
	pingResponse := u.Ping(c.RInstance)
	if pingResponse["status"] != true {
		u.Respond(w, u.Message(true, pingResponse["message"].(string)))
		return
	}
	u.Respond(w, u.Message(true, "Health check OK"))
	return
}
