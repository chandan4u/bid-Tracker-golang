package api

import (
	u "berlin/utils"
	"net/http"
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

	// [Make Map : User struct data structure for storing all User information]
	allBidsByUsers := make(map[string][]u.User)
	winner := u.User{}
	bidAmount := 0

	// [redClient SCAN : scan redis on the basis of item and return redis key]
	recordIteration := redClient.RInstance.Scan(0, "*"+requestItem, 0).Iterator()
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

		if bidAmount == 0 || bidAmount > userRecords.Amount {
			bidAmount = userRecords.Amount
			winner = userRecords
		}
	}
	if err := recordIteration.Err(); err != nil {
		u.Respond(resW, u.Message(false, "Oops something went's wrong!"))
		return
	}
	// [Append : append current record into allBidsByUsers map User struct]
	allBidsByUsers["winner"] = append(allBidsByUsers["winner"], winner)

	u.RespondWithData(resW, allBidsByUsers)
	return
}
