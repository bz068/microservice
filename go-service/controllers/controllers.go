package controllers

import (
	"net/http"
	"posts/utils"
	"time"
	"encoding/json"
	"math/rand"
	"fmt"
)

type post struct {
	Id 			int 			`json:"id"`
	Title 		string			`json:"title"`
	Body 		string			`json:"body"`
	Date 		time.Time 		`json:"date"`
	Comments 	[]comment 		`json:"comments"`
}

type comment struct {
	Id 			int 		`json:"id"`
	PostId 		int 		`json:"postId"`
	Comment 	string 		`json:"comment"`
}

type postsStore struct {
	store []post
}

type resp struct {
	Data 		[]comment 		`json:"data"`
	Msg 		string 			`json:"msg"`
	Success 	bool 			`json:"success"`

}

func returnPostsFromStore() *postsStore{
	return &postsStore{
		store: []post{post{1, "title", "body", time.Now(), []comment{}}},}
	}


var store = returnPostsFromStore()

func IndexController(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to golang-node Microservice V1</h1>"))
}

func GetPosts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	for i, singlePost := range store.store {
		res, err := http.Get(fmt.Sprintf("http://localhost:9000/api/v1/posts/%d/comments", singlePost.Id))
		if err != nil {
			resObj := utils.CreateResponseObject([]int{}, "Request to Comments MS Failed", false)
			w.Write(resObj)
			return
		}

		var responseObj resp
		jsonErr := json.NewDecoder(res.Body).Decode(&responseObj)
		if jsonErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			resObj := utils.CreateResponseObject([]int{}, "Malformed Request", false)
			w.Write(resObj)
			return
		}
		store.store[i].Comments = responseObj.Data
}
	resObj := utils.CreateResponseObject(store.store, "Posts Fetched Successfully", true)
	w.Write(resObj)
}
func generateId() int {
	return rand.Int()
}

func AddPosts(w http.ResponseWriter, r *http.Request){

	var p post
	
	err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
		resObj := utils.CreateResponseObject([]int{}, "Malformed Request", false)
		w.Write(resObj)
        return
    }

	if err, res := utils.ValidateFeild(p.Title, "Please Enter Title"); err {
		w.Write(res)
		return
	}
	if err, res := utils.ValidateFeild(p.Body, "Please Enter Body Text"); err {
		w.Write(res)
		return
	}

	p.Id = generateId()
	p.Date = time.Now()

	store.store = append(store.store, p)

	resObj := utils.CreateResponseObject(p, "Posts Added Successfully", true)
	w.Write(resObj)
}