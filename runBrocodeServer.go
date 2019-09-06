package main

import (
	_ "fmt"
	_ "github.com/blackdartq/brocode/goResources/dbutil"
	"github.com/blackdartq/brocode/goResources/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/getBlogPosts", routes.GetAllBlogPosts)
	router.StaticFile("/", "./web/html/index.html")
	router.StaticFS("/html", http.Dir("./web/html"))
	router.StaticFS("/css", http.Dir("./web/css"))
	router.StaticFS("/js", http.Dir("./web/js"))
	router.StaticFS("/img", http.Dir("./web/img"))
	router.Run(":8080")

	//database := dbutil.CreateDbutil()
	//database.GetAllBlogPosts()
	//testing := dbutil.BlogPost{3, "added by Golang", "I just added a post by Golang"}
	//database.AddBlogPost(testing)
	//database.PrintAllBlogPosts()
	//database.DeleteBlogPostById(2)
	//database.PrintAllBlogPosts()

}
