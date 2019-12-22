package api

import (
	u "berlin/utils"
	"net/http"
	"strconv"
)

/*
	[http call for get winner on the basis of item]
	[ POST ] http://localhost:8080/api/berlin/internal/winner-by-item?item=t2
	params {
		item: t3
	}
*/

// GetWinnerByItem : Get winner on the basis of item.
func (redClient *RedisInstance) GetWinnerByItem(resW http.ResponseWriter, reqR *http.Request) {

	// [ UserRequestProcessing ] Request parse and params validation
	requestProcessingStatus := u.ItemRequestProcessing(reqR)
	if requestProcessingStatus["status"] != true {
		u.Respond(resW, u.Message(false, requestProcessingStatus["message"].(string)))
		return
	}
	filterRequestParams := requestProcessingStatus["data"].(map[string]interface{})
	requestItem := filterRequestParams["requestItem"].(string)

	// [ Make Map : Storing winner information inside Map interface ]
	var allBidsByUsers = make(map[string]interface{})
	var winner = make(map[string]interface{})
	bidAmount := 0
	var amount int

	// [ redClient SCAN : scan redis on the basis of item and return redis key ]
	recordIteration := redClient.RInstance.Scan(0, "*"+requestItem, 0).Iterator()
	for recordIteration.Next() {

		// [ redClient HGetAll : get all data on the basis of key value ]
		m, err := redClient.RInstance.HGetAll(recordIteration.Val()).Result()
		if err != nil {
			u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
			return
		}

		// [ Checking bidAmount and comparing with min amount ]
		amount, _ = strconv.Atoi(m["Amount"])
		if bidAmount == 0 || bidAmount > amount {
			bidAmount = amount
			winner["username"] = m["Username"]
			winner["amount"] = m["Amount"]
		}
	}
	if err := recordIteration.Err(); err != nil {
		u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
		return
	}

	// [ Checking is there anyone bid for the given item ]
	if len(winner) == 0 {
		u.Respond(resW, u.Message(true, "There is no any bid for the given Item :: "+requestItem))
		return
	}

	// [ Append : append current record into allBidsByUsers map[string]interface{} ]
	allBidsByUsers[requestItem] = winner

	u.Respond(resW, allBidsByUsers)
	return
}
