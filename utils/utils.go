package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// User is a simple user struct for this example
type User struct {
	Username string `json:"username"`
	Item     string `json:"item"`
	Amount   int    `json:"amount"`
}

// Message : [ It help to handle custom messages in map interface structure ]
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// RequestDataWithMessage : [ It help to handle cutom message with extra value in interface structure ]
func RequestDataWithMessage(status bool, message string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message, "data": data}
}

// Respond : [ It help to handle http respone in json format]
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// AddBidRequestValidation : [ User request validation for fields username, item and amount ]
func AddBidRequestValidation(reqUsername string, reqItem string, reqAmount int) map[string]interface{} {
	if len(reqUsername) < 2 {
		return Message(false, "Username should greater than 1 character.")
	}
	if len(reqItem) < 2 {
		return Message(false, "Item should greater than 1 character.")
	}
	if reqAmount < 1 {
		return Message(false, "Amount should greater than 0.")
	}
	return Message(true, "User input successfully validate.")
}

// AddBidRequestProcessing : [ It process the add biding request and validate values are correct or not]
func AddBidRequestProcessing(reqR *http.Request) map[string]interface{} {
	requestFormError := reqR.ParseForm()
	if requestFormError != nil {
		return Message(false, "Oops something wents wrong!")
	}
	requestUsername := reqR.Form.Get("username")
	requestItem := reqR.Form.Get("item")
	requestAmount := reqR.Form.Get("amount")
	if requestAmount == "" || requestItem == "" || requestUsername == "" {
		return Message(false, "Required all the params for process bid.")
	}
	amount, cnvtError := strconv.Atoi(requestAmount)
	if cnvtError != nil {
		return Message(false, "Request bit amount conversation error")
	}
	// [ RequestValidation ] : Validating request value
	validationStatus := AddBidRequestValidation(requestUsername, requestItem, amount)
	if validationStatus["status"] != true {
		return Message(false, validationStatus["message"].(string))
	}
	// [ filteredUserInfo ] : Adding filter values in function response
	filteredUserInfo := map[string]interface{}{
		"requestUsername": requestUsername,
		"requestItem":     requestItem,
		"requestAmount":   amount,
	}
	return RequestDataWithMessage(true, "None", filteredUserInfo)
}

// MapRequestedDataIntoUser :  MapRequestedDataIntoUser
func MapRequestedDataIntoUser(redisData map[string]string) map[string]interface{} {
	usr := User{}
	for key, value := range redisData {
		switch key {
		case "Username":
			usr.Username = value
		case "Item":
			usr.Item = value
		case "Amount":
			usr.Amount, _ = strconv.Atoi(value)
		}
	}
	filteredUserInfo := map[string]interface{}{
		"requestUserInfo": usr,
	}
	return RequestDataWithMessage(true, "None", filteredUserInfo)
}

// UserRequestValidation : UserRequestValidation
func UserRequestValidation(reqUsername string) map[string]interface{} {
	if len(reqUsername) < 2 {
		return Message(false, "Username should greater than 1 character.")
	}
	return Message(true, "User input successfully validate.")
}

// UserRequestProcessing : [ It process the add biding request and validate values are correct or not]
func UserRequestProcessing(reqR *http.Request) map[string]interface{} {
	requestFormError := reqR.ParseForm()
	if requestFormError != nil {
		return Message(false, "Oops something wents wrong!")
	}
	requestUsername := reqR.Form.Get("username")
	// [ RequestValidation ] : Validating request value
	validationStatus := UserRequestValidation(requestUsername)
	if validationStatus["status"] != true {
		return Message(false, validationStatus["message"].(string))
	}
	filteredUsername := map[string]interface{}{
		"requestUsername": requestUsername,
	}
	return RequestDataWithMessage(true, "None", filteredUsername)
}

// ItemRequestValidation : UserRequestValidation
func ItemRequestValidation(reqItem string) map[string]interface{} {
	if len(reqItem) < 2 {
		return Message(false, "Item should greater than 1 character.")
	}
	return Message(true, "User input successfully validate.")
}

// ItemRequestProcessing : [ It process the add biding request and validate values are correct or not]
func ItemRequestProcessing(reqR *http.Request) map[string]interface{} {
	requestFormError := reqR.ParseForm()
	if requestFormError != nil {
		return Message(false, "Oops something wents wrong!")
	}
	requestItem := reqR.Form.Get("item")
	// [ RequestValidation ] : Validating request value
	validationStatus := ItemRequestValidation(requestItem)
	if validationStatus["status"] != true {
		return Message(false, validationStatus["message"].(string))
	}
	filteredItem := map[string]interface{}{
		"requestItem": requestItem,
	}
	return RequestDataWithMessage(true, "None", filteredItem)
}
