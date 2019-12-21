package api

import (
	u "berlin/utils"
	"net/http"

	"github.com/fatih/structs"
)

/*
	[http call for add bid]
	[ POST ] http://localhost:8080/api/berlin/internal/add-biding?username=daniel&amount=10&item=bitcoin
	params {
		username: daniel,
		item: bitcoin,
		amount: 10
	}
*/

// AddBiding : Add bid on the basis of username, item and amount inside redis.
func (redClient *RedisInstance) AddBiding(resW http.ResponseWriter, reqR *http.Request) {
	// [ AddBidRequestProcessing ] Request parse and params validation
	requestProcessingStatus := u.AddBidRequestProcessing(reqR)
	if requestProcessingStatus["status"] != true {
		u.Respond(resW, u.Message(false, requestProcessingStatus["message"].(string)))
		return
	}

	filterRequestParams := requestProcessingStatus["data"].(map[string]interface{})
	requestUsername := filterRequestParams["requestUsername"].(string)
	requestItem := filterRequestParams["requestItem"].(string)
	requestAmount := filterRequestParams["requestAmount"].(int)

	// [ User ] Converting filter user input into User struct for inserting in redis.
	userStruct := User{
		Username: requestUsername,
		Item:     requestItem,
		Amount:   requestAmount,
	}
	// [structs] Using struct library convert the request into Map then insert into redis.
	userStructMap := structs.Map(userStruct)
	redisSetErr := redClient.RInstance.HMSet(requestUsername+"_"+requestItem, userStructMap).Err()
	if redisSetErr != nil {
		u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
		return
	}
	u.Respond(resW, u.Message(true, "Bid successfully added for username :: "+requestUsername))
	return
}
