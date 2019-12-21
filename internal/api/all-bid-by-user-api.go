package api

import (
	u "berlin/utils"
	"encoding/json"
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

	// [Make Map : User struct data structure for storing all User information]
	allBidsByUsers := make(map[string][]u.User)

	// [redClient SCAN : scan redis on the basis of username and return redis key]
	recordIteration := redClient.RInstance.Scan(0, requestUsername+"*", 0).Iterator()
	for recordIteration.Next() {

		// [redClient HGetAll : get all data on the basis of key value]
		m, err := redClient.RInstance.HGetAll(recordIteration.Val()).Result()
		if err != nil {
			u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
			return
		}

		// [MapRequestedDataIntoUser : It map all the redis key value on User struct]
		mappedDataInUserStruct := u.MapRequestedDataIntoUser(m)
		convertedInterface := mappedDataInUserStruct["data"].(map[string]interface{})
		userRecords := convertedInterface["requestUserInfo"].(u.User)

		// [Append : append current record into allBidsByUsers map User struct]
		allBidsByUsers[recordIteration.Val()] = append(allBidsByUsers[recordIteration.Val()], userRecords)
	}
	if err := recordIteration.Err(); err != nil {
		u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
		return
	}

	json.NewEncoder(resW).Encode(allBidsByUsers)
	return
}
