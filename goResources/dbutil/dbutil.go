package dbutil

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "strconv"
)

type Dbutil struct {
	db        *sql.DB
	BlogPosts []BlogPost
}

type BlogPost struct {
	PostId  int
	Name    string
	Message string
}

// CreateDatabase creates a connection to the database
func CreateDbutil() Dbutil {
	// Database stuff
	connectionStr := "user=blackdartq dbname=brocode sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	return Dbutil{db, make([]BlogPost, 0)}
}

// GetAllBlogPosts() gets all the post names and messages from the blog table in the database
func (dbutils *Dbutil) GetAllBlogPosts() {
	rows, err := dbutils.db.Query("SELECT * FROM blog;")
	if err != nil {
		panic("error querying")
	}
	for rows.Next() {
		var id int
		var name string
		var post string
		if err := rows.Scan(&id, &name, &post); err != nil {
			panic("couldn't get fields")
		}
		dbutils.BlogPosts = append(dbutils.BlogPosts, BlogPost{id, name, post})

	}
}

// AddBlogPost adds a post to the blog database
func (dbutils *Dbutil) AddBlogPost(blogPost BlogPost) {
	_, err := dbutils.db.Exec("INSERT INTO blog VALUES(DEFAULT, $1, $2);", blogPost.Name, blogPost.Message)
	if err != nil {
		//panic("couldn't add blog post to the database")
		panic(err)
	}
	dbutils.GetAllBlogPosts()
}

// DeleteBlogPost will delete a post by its ID
func (dbutils *Dbutil) DeleteBlogPost(blogPost BlogPost) {
	_, err := dbutils.db.Exec("DELETE FROM blog WHERE post_id = $1;", blogPost.PostId)
	if err != nil {
		//panic("couldn't add blog post to the database")
		panic(err)
	}
	dbutils.GetAllBlogPosts()
}

// DeleteBlogPost will delete a post by its ID
func (dbutils *Dbutil) DeleteBlogPostById(postId int64) {
	_, err := dbutils.db.Exec("DELETE FROM blog WHERE post_id = $1;", postId)
	if err != nil {
		//panic("couldn't add blog post to the database")
		panic(err)
	}
	dbutils.GetAllBlogPosts()
}

func (dbutils Dbutil) PrintAllBlogPosts() {
	for _, blogPost := range dbutils.BlogPosts {
		fmt.Printf("%v\t%s\n%s\n\n", blogPost.PostId, blogPost.Name, blogPost.Message)
	}
}
