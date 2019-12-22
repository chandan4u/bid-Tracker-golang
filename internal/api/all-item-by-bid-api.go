package api

import (
	u "berlin/utils"
	"net/http"
)

/*
	[http call for get all bids on the basis of item]
	[ POST ] http://localhost:8080/api/berlin/internal/all-item-by-bid?item=t3
	params {
		item: t3
	}
*/

// GetAllItemByBid : Get all bids on the basis of Item.
func (redClient *RedisInstance) GetAllItemByBid(resW http.ResponseWriter, reqR *http.Request) {

	// [ UserRequestProcessing ] Request parse and params validation
	requestProcessingStatus := u.ItemRequestProcessing(reqR)
	if requestProcessingStatus["status"] != true {
		u.Respond(resW, u.Message(false, requestProcessingStatus["message"].(string)))
		return
	}
	filterRequestParams := requestProcessingStatus["data"].(map[string]interface{})
	requestItem := filterRequestParams["requestItem"].(string)

	// [Make Map : User struct data structure for storing all User information]
	var allBidsByUsers = make(map[string]interface{})
	var bidbyuser []interface{}

	// [redClient SCAN : scan redis on the basis of item and return redis key]
	recordIteration := redClient.RInstance.Scan(0, "*"+requestItem, 0).Iterator()
	for recordIteration.Next() {

		// [ Add temp Map interface to hold current redis user information ]
		var temp = make(map[string]interface{})

		// [redClient HGetAll : get all data on the basis of key value]
		m, err := redClient.RInstance.HGetAll(recordIteration.Val()).Result()
		if err != nil {
			u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
			return
		}

		// [ Append : Take user information form temp memory and append in bidbyuser []interface ]
		temp["username"] = m["Username"]
		temp["amount"] = m["Amount"]
		bidbyuser = append(bidbyuser, temp)
	}

	// [ Assign all user information inside allBidsByUsers map[string]interface{} ]
	allBidsByUsers[requestItem] = bidbyuser

	if err := recordIteration.Err(); err != nil {
		u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
		return
	}

	u.Respond(resW, allBidsByUsers)
	return
}
