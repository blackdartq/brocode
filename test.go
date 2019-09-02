package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	router := gin.Default()
	router.StaticFS("/html", http.Dir("./web/html"))
	router.StaticFS("/css", http.Dir("./web/css"))
	router.StaticFS("/js", http.Dir("./web/js"))
	router.StaticFS("/img", http.Dir("./web/img"))
	router.Run(":8080")

	// Database stuff
	connectionStr := "user=blackdartq dbname=brocode sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	names := make([]string, 0)
	posts := make([]string, 0)

	rows, err := db.Query("SELECT * FROM blog;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id string
		var name string
		var post string
		if err := rows.Scan(&id, &name, &post); err != nil {
		}
		names = append(names, name)
		posts = append(posts, post)
	}
	fmt.Println(names[0], posts[1])
}
