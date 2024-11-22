package forum

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateUsersTable(db *sql.DB) {
	// Creation de la Table Users
	createUserTableSql := `
		CREATE TABLE IF NOT EXISTS users(
			UUID TEXT NOT NULL,
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := db.Exec(createUserTableSql)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreatePostsTable(db *sql.DB) {
	// Creation de la Table Posts
	createPostsTableSql := `
		CREATE TABLE IF NOT EXISTS posts(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id TEXT NOT NULL,
			content TEXT NOT NULL,
			category1 TEXT NOT NULL,
			category2 TEXT NOT NULL,
			category3 TEXT NOT NULL,
			date TEXT NOT NULL,
			com_number INT DEFAULT 0
		)
	`
	_, err := db.Exec(createPostsTableSql)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateLikeTable(db *sql.DB) {
	// Creation de la Table Like
	createLikeTableSql := `
		CREATE TABLE IF NOT EXISTS like(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER NOT NULL,
			username TEXT NOT NULL,
			like INTEGER NOT NULL,	
			dislike INTEGER NOT NULL
		)
	`
	_, err := db.Exec(createLikeTableSql)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateCommentsTable(db *sql.DB) {
	// Creation de la Table Comments
	createCommentTableSql := `
		CREATE TABLE IF NOT EXISTS comments(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			auteur TEXT NOT NULL,
			post_id INTEGER NOT NULL,
			comment TEXT NOT NULL,
			date TEXT NOT NULL
		)
	`
	_, err := db.Exec(createCommentTableSql)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateCategoryTable(db *sql.DB) {
	// Creation de la Table Category
	createLikeCommentTableSql := `
		CREATE TABLE IF NOT EXISTS commentlike(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			comment_id INTEGER NOT NULL,
			username TEXT NOT NULL,
			like INTEGER NOT NULL,	
			dislike INTEGER NOT NULL
		)
	`
	_, err := db.Exec(createLikeCommentTableSql)
	if err != nil {
		fmt.Println(err)
		return
	}
}
