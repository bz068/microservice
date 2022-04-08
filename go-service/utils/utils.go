package utils

import (
	"log"
	"fmt"
	"net/http"
	"regexp"
	"encoding/json"
)


func LogServerLive(){
	fmt.Println("*********************************************************")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("")
	log.Println("Server live: http://localhost:8080/")
	fmt.Println("")
	fmt.Println("--------------------------------------------------------")
	fmt.Println("*********************************************************")

}

func PathLogger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	}
}

func MatchRoute(expression string, path string) bool {
	if matched, err := regexp.MatchString(expression, path); err != nil {
		return false
	} else {
		return matched
	}
}

func CreateResponseObject(data interface{}, message string, success bool) []uint8 {
	res := make(map[string]interface{})

	res["success"] = success
	res["message"] = message
	res["data"] = data

	jsonBytes, err := json.Marshal(res)

	if err != nil {
		res["success"] = false
		res["message"] = "Internal Server Error"
		res["data"] = []int{}
	}

	return jsonBytes
}

func ValidateFeild(field string, msg string) (bool, []uint8) {
	if len(field) == 0 { 
		resObj := CreateResponseObject([]int{}, msg, false)
		return true, resObj
	 } else {
		 return false, nil
	 }
}