package main

import (
	_ "fmt"
	_ "github.com/blackdartq/brocode/goResources/dbutil"
	"github.com/blackdartq/brocode/goResources/routes"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func main() {
	router := gin.Default()
	router.StaticFile("", "./web/index.html")
	router.StaticFile("blog.html", "./web/blog.html")
	router.StaticFile("wiki.html", "./web/wiki.html")
	router.StaticFile("goals.html", "./web/goals.html")
	router.StaticFile("projects.html", "./web/projects.html")
	router.Static("/css", "./web/css")
	router.Static("/js", "./web/js")
	router.Static("/img", "./web/img")
	router.GET("/getBlogPosts", routes.GetAllBlogPosts)

	router.Run(":8080")

	//database := dbutil.CreateDbutil()
	//database.GetAllBlogPosts()
	//testing := dbutil.BlogPost{3, "added by Golang", "I just added a post by Golang"}
	//database.AddBlogPost(testing)
	//database.PrintAllBlogPosts()
	//database.DeleteBlogPostById(2)
	//database.PrintAllBlogPosts()

}
