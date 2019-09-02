package routes

import (
	"github.com/blackdartq/brocode/goResources/dbutil"
	"github.com/gin-gonic/gin"
)

func GetAllBlogPosts(c *gin.Context) {
	database := dbutil.CreateDbutil()
	database.GetAllBlogPosts()
	c.JSON(200, database.BlogPosts)
	database.PrintAllBlogPosts()
}
