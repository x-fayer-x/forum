package forum

import (
	"fmt"
	"net/http"
)

// retrieve Username by uuid in cookie
func GetUsername(r *http.Request) (string, error) {
	var username string

	// Get the cookie from the user
	cookie, err := r.Cookie("UserToken")
	if err != nil {
		return "", fmt.Errorf("GetUsername failure to get cookie : %w ", err)
	}

	cookieUUID := cookie.Value

	// Open Database
	db, err := openDB()
	if err != nil {
		return "", fmt.Errorf("GetUsername failure to open database : %w ", err)
	}
	defer db.Close()

	// Open transction in database
	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("GetUsername failure to start transaction : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// query the username in users where UUID a la valuer de cookieUUID
	err = tx.QueryRow("SELECT username FROM users WHERE UUID = ?", cookieUUID).Scan(&username)
	if err != nil {
		return "", fmt.Errorf("GetUsername failure to query : %w ", err)
	}

	// end of the transaction
	tx.Commit()
	if err != nil {
		failCommit = false
		return "", fmt.Errorf("GetUsername failure to commit transaction : %w ", err)
	}

	return username, nil
}

func GetEmail(r *http.Request) (string, error) {
	var email string

	// Get the cookie from the user
	cookie, err := r.Cookie("UserToken")
	if err != nil {
		return "", fmt.Errorf("GetUsername failure to get cookie : %w ", err)
	}

	cookieUUID := cookie.Value

	// Open database
	db, err := openDB()
	if err != nil {
		return "", fmt.Errorf("GetEmail failure to open Database : %w ", err)
	}
	defer db.Close()

	// Open transaction
	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("GetEmail failure to start transaction")
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	err = tx.QueryRow("SELECT email FROM users WHERE UUID = ?", cookieUUID).Scan(&email)
	if err != nil {
		return "", fmt.Errorf("GetEmail failure to query email")
	}

	err = tx.Commit()
	if err != nil {
		failCommit = false
		return "", fmt.Errorf("GetEmail failure to commit transaction : %w ", err)
	}

	return email, nil
}

// Retrieve password with username
func GetPassWord(r *http.Request, username string) (string, error) {
	var passWordDB string = ""

	// Open database
	db, err := openDB()
	if err != nil {
		return "", fmt.Errorf("GetPassWord failure to open Database : %w ", err)
	}
	defer db.Close()

	// Open transaction
	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("GetPassWord failure to start transaction")
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// query the password in users by a username
	err = tx.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&passWordDB)
	if err != nil {
		return "", fmt.Errorf("GetPassWord failure to query : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		failCommit = false
		return "", fmt.Errorf("GetPassWord failure to commit transaction : %w ", err)
	}

	return passWordDB, nil
}

func GetAllPost() ([]Posts, error) {
	var allpost []Posts
	var thepost Posts
	var query string = "SELECT id, user_id, content, category1, category2, category3, date FROM posts "
	// Open DataBase
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("GetallPost failure to open DataBase : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("GetallPost failure to start transaction : %w", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	row, err := tx.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetallPost failure to query : %w ", err)
	}
	for row.Next() {

		thepost = Posts{}

		err = row.Scan(&thepost.Id, &thepost.UserName, &thepost.Content, &thepost.Category1, &thepost.Category2, &thepost.Category3, &thepost.Date)
		if err != nil {
			return nil, fmt.Errorf("GetallPost failure to scan : %w ", err)
		}

		// Ajouter les catégories non vides à la liste de catégories
		if thepost.Category1 != "" {
			thepost.Category = append(thepost.Category, thepost.Category1)
		}
		if thepost.Category2 != "" {
			thepost.Category = append(thepost.Category, thepost.Category2)
		}
		if thepost.Category3 != "" {
			thepost.Category = append(thepost.Category, thepost.Category3)
		}

		allpost = append(allpost, thepost)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return nil, fmt.Errorf("GetallPost failure to commit transaction : %w ", err)
	}

	return allpost, nil
}

func GetPost(postID int) (Posts, error) {
	var thepost Posts
	var query string = "SELECT id, user_id, content, category1, category2, category3, date FROM posts WHERE id = ?"

	// Open DataBase
	db, err := openDB()
	if err != nil {
		return Posts{}, fmt.Errorf("GetPost failure to open DataBase : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return Posts{}, fmt.Errorf("GetPost failure to start transaction : %w", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	err = tx.QueryRow(query, postID).Scan(&thepost.Id, &thepost.UserName, &thepost.Content, &thepost.Category1, &thepost.Category2, &thepost.Category3, &thepost.Date)
	if err != nil {
		return Posts{}, fmt.Errorf("GetPost failure to query : %w ", err)
	}

	// Ajouter les catégories non vides à la liste de catégories
	if thepost.Category1 != "" {
		thepost.Category = append(thepost.Category, thepost.Category1)
	}
	if thepost.Category2 != "" {
		thepost.Category = append(thepost.Category, thepost.Category2)
	}
	if thepost.Category3 != "" {
		thepost.Category = append(thepost.Category, thepost.Category3)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return Posts{}, fmt.Errorf("GetPost failure to commit transaction : %w ", err)
	}

	return thepost, nil
}

func GetLikePost(PostsID int) ([]LikePost, error) {
	var likes []LikePost
	var like LikePost
	var query string = "SELECT post_id, username, like, dislike FROM like WHERE post_id = ?"

	// open Database
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("GetLikePost failure to open database : %w ", err)
	}
	defer db.Close()

	// start transaction in database
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("GetLikePost failure to start transaction : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// query like
	row, err := tx.Query(query, PostsID)
	if err != nil {
		return nil, fmt.Errorf("GetLikePost failure to query")
	}

	for row.Next() {
		err = row.Scan(&like.PostId, &like.Username, &like.Like, &like.Dislike)
		if err != nil {
			return nil, fmt.Errorf("GetLikePost failure to scan query : %w ", err)
		}
		likes = append(likes, like)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return nil, fmt.Errorf("GetLikePost failure to commit transaction : %w ", err)
	}

	return likes, nil
}

func GetComment(PostID int) ([]Comments, error) {
	var comments []Comments
	var comment Comments
	var query string = "SELECT id, post_id, auteur, comment, date FROM comments WHERE post_id = ?"

	// Open Database
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("GetComment failure to open Database : %w ", err)
	}
	defer db.Close()

	// start transaction in Database
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("GetComment failure to open Database : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// query comment
	row, err := tx.Query(query, PostID)
	if err != nil {
		return nil, fmt.Errorf("GetComment failure to query db : %w", err)
	}

	for row.Next() {
		err = row.Scan(&comment.Id, &comment.PostID, &comment.UserName, &comment.Content, &comment.Date)
		if err != nil {
			return nil, fmt.Errorf("GetComment failure to scan query : %w ", err)
		}
		comments = append(comments, comment)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return nil, fmt.Errorf("GetComment failure to commit transaction : %w ", err)
	}

	return comments, nil
}

func GetLikeComment(CommentID int) ([]LikeComment, error) {
	var likes []LikeComment
	var like LikeComment
	var query string = "SELECT comment_id, username, like, dislike FROM commentlike WHERE comment_id = ?"

	// Open Database
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("GetlikeComment failure to open database : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("GetlikeComment failure to start transaction : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// query like of comment
	row, err := tx.Query(query, CommentID)
	if err != nil {
		return nil, fmt.Errorf("GetlikeComment failure to query db : %w ", err)
	}

	for row.Next() {
		err = row.Scan(&like.CommentId, &like.Username, &like.Like, &like.Dislike)
		if err != nil {
			return nil, fmt.Errorf("GetlikeComment failure to scan query : %w ", err)
		}
		likes = append(likes, like)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return nil, fmt.Errorf("GetlikeComment failure to commimt transaction : %w ", err)
	}
	return likes, nil
}

func GetLikeUser(Username string) ([]LikePost, error) {
	var likes []LikePost
	var like LikePost
	var query string = "SELECT post_id, username, like, dislike FROM like WHERE username = ?"

	// open Database
	db, err := openDB()
	if err != nil {
		return nil, fmt.Errorf("GetLikePost failure to open database : %w ", err)
	}
	defer db.Close()

	// start transaction in database
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("GetLikePost failure to start transaction : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// query like
	row, err := tx.Query(query, Username)
	if err != nil {
		return nil, fmt.Errorf("GetLikePost failure to query")
	}

	for row.Next() {
		err = row.Scan(&like.PostId, &like.Username, &like.Like, &like.Dislike)
		if err != nil {
			return nil, fmt.Errorf("GetLikePost failure to scan query : %w ", err)
		}
		likes = append(likes, like)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return nil, fmt.Errorf("GetLikePost failure to commit transaction : %w ", err)
	}

	return likes, nil
}
