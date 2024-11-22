package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	f "forum/forum"
	_ "github.com/mattn/go-sqlite3"
	
)

const port = ":8080"

func main() {
	// Ouvrir la base de donnée SQLite
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	f.CreateUsersTable(db)
	f.CreateCommentsTable(db)
	f.CreatePostsTable(db)
	f.CreateLikeTable(db)
	f.CreateCategoryTable(db)

	go f.CleanupVisitors()
	mux := http.NewServeMux()
	mux.Handle("/", f.LimitMiddleware(http.HandlerFunc(f.Handler)))
	mux.Handle("/home", f.LimitMiddleware(http.HandlerFunc(f.Handlerin)))
	mux.Handle("/team", f.LimitMiddleware(http.HandlerFunc(f.HandleTeam)))
	mux.Handle("/teamin", f.LimitMiddleware(http.HandlerFunc(f.HandleTeamin)))
	mux.Handle("/forum", f.LimitMiddleware(http.HandlerFunc(f.HandleForum)))
	mux.Handle("/forumin", f.LimitMiddleware(http.HandlerFunc(f.HandleForumin)))
	mux.Handle("/edit-post", f.LimitMiddleware(http.HandlerFunc(f.HandleEdit)))
	mux.Handle("/login", f.LimitMiddleware(http.HandlerFunc(f.HandleLogin)))
	mux.Handle("/edit-comment", f.LimitMiddleware(http.HandlerFunc(f.HandleComment)))
	mux.Handle("/results", f.LimitMiddleware(http.HandlerFunc(f.HandleResults)))
	mux.Handle("/like", f.LimitMiddleware(http.HandlerFunc(f.LikePostHandler)))
	mux.Handle("/dislike", f.LimitMiddleware(http.HandlerFunc(f.DislikePostHandler)))
	mux.Handle("/users", f.LimitMiddleware(http.HandlerFunc(f.HandlerUser)))
	mux.Handle("/commentslike", f.LimitMiddleware(http.HandlerFunc(f.LikeCommentHandler)))
	mux.Handle("/commentsdislike", f.LimitMiddleware(http.HandlerFunc(f.DislikeCommentHandler)))
	mux.HandleFunc("/UserFilter", f.Handlerfilter)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	server := &http.Server{
        Addr:         port,
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    fmt.Println("http://localhost:8080 - server started on port", port)
    server.ListenAndServe()
}
// http.Handle("/", limitMiddleware(http.HandlerFunc(f.Handler)))
// http.Handle("/home", limitMiddleware(http.HandlerFunc(f.Handlerin)))
// http.Handle("/team", limitMiddleware(http.HandlerFunc(f.HandleTeam)))
// http.Handle("/teamin", limitMiddleware(http.HandlerFunc(f.HandleTeamin)))
// http.Handle("/forum", limitMiddleware(http.HandlerFunc(f.HandleForum)))
// http.Handle("/forumin", limitMiddleware(http.HandlerFunc(f.HandleForumin)))
// http.Handle("/edit-post", limitMiddleware(http.HandlerFunc(f.HandleEdit)))
// http.Handle("/login", limitMiddleware(http.HandlerFunc(f.HandleLogin)))
// http.Handle("/edit-comment", limitMiddleware(http.HandlerFunc(f.HandleComment)))
// http.Handle("/results", limitMiddleware(http.HandlerFunc(f.HandleResults)))
// http.Handle("/like", limitMiddleware(http.HandlerFunc(f.LikePostHandler)))
// http.Handle("/dislike", limitMiddleware(http.HandlerFunc(f.DislikePostHandler)))
// http.Handle("/users", limitMiddleware(http.HandlerFunc(f.HandlerUser)))
// http.Handle("/commentslike", limitMiddleware(http.HandlerFunc(f.LikeCommentHandler)))
// http.Handle("/commentsdislike", limitMiddleware(http.HandlerFunc(f.DislikeCommentHandler)))