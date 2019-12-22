package api

import (
	u "berlin/utils"
	"net/http"
)

/*
	[http call for get all items that user bid]
	[ POST ] http://localhost:8000/api/berlin/internal/all-bids-by-user?username=daniel
	params {
		username: daniel
	}
*/

// GetAllBidByUser : Get all items on the basis of username.
func (redClient *RedisInstance) GetAllBidByUser(resW http.ResponseWriter, reqR *http.Request) {

	// [ UserRequestProcessing ] Request parse and params validation
	requestProcessingStatus := u.UserRequestProcessing(reqR)
	if requestProcessingStatus["status"] != true {
		u.Respond(resW, u.Message(false, requestProcessingStatus["message"].(string)))
		return
	}
	filterRequestParams := requestProcessingStatus["data"].(map[string]interface{})
	requestUsername := filterRequestParams["requestUsername"].(string)

	// [ Make Map : User struct data structure for storing all User information ]
	var allBidsByUsers = make(map[string]interface{})
	var bidbyuser []interface{}

	// [ redClient SCAN : scan redis on the basis of username and return redis key ]
	recordIteration := redClient.RInstance.Scan(0, requestUsername+"*", 0).Iterator()
	for recordIteration.Next() {

		// [ Add temp Map interface to hold current redis user information ]
		var temp = make(map[string]interface{})

		// [ redClient HGetAll : get all data on the basis of key value ]
		m, err := redClient.RInstance.HGetAll(recordIteration.Val()).Result()
		if err != nil {
			u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
			return
		}

		// [ Append : Take user information form temp memory and append in bidbyuser []interface ]
		temp["item"] = m["Item"]
		temp["amount"] = m["Amount"]
		bidbyuser = append(bidbyuser, temp)
	}

	// [ Assign all user information inside allBidsByUsers map[string]interface{} ]
	allBidsByUsers[requestUsername] = bidbyuser

	if err := recordIteration.Err(); err != nil {
		u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
		return
	}

	u.Respond(resW, allBidsByUsers)
	return
}
