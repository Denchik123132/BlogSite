package database

import (
	"BlogSite/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreateTable(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		image TEXT
	);
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "internal/storage/database/posts.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AllPosts() []models.Post {
	db := GetDB()
	var allPosts []models.Post
	rows, err := db.Query("select * from posts")
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.Image)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allPosts = append(allPosts, post)
	}
	if err != nil {
		log.Println("Error: ", err)
	}
	return allPosts
}

func PostByID(id int) models.Post {
	db := GetDB()
	var post models.Post
	err := db.QueryRow("SELECT * FROM posts WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Description, &post.Image)
	if err != nil {
		log.Println("Error: ", err)
	}
	return post

}

func AddPost(title string, description string, imageName string) {
	db := GetDB()
	var lastPostID int
	err := db.QueryRow("SELECT MAX(ID) FROM posts").Scan(&lastPostID)
	if err != nil {
		log.Println("Error: ", err)
	}
	newPostID := lastPostID + 1

	relativeImagePath := ""
	if imageName != "" {
		relativeImagePath = "/assets/" + imageName
	}

	_, err = db.Exec("insert into posts (id, title, description, image) values ($1, $2, $3, $4)", newPostID, title, description, relativeImagePath)
	if err != nil {
		return
	}
}
