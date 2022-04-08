package routes

import (
	"net/http"
	"posts/controllers"
	"posts/utils"
)


func RouteHandler(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path

	if(utils.MatchRoute(`^/$`, path)){
		controllers.IndexController(w, r)
		return
	} else if (utils.MatchRoute(`^/api/v1/allposts$`, path) && r.Method == "GET"){
		controllers.GetPosts(w,r)
		return
	} else if (utils.MatchRoute(`^/api/v1/addpost$`, path) && r.Method == "POST"){
		controllers.AddPosts(w,r)
		return
	} else {
		resObj := utils.CreateResponseObject([]int{}, "Route not found, Please try again!", false)
		w.Write(resObj)
	}
}