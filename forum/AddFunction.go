package forum

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

// Add User to Database
func AddUser(w http.ResponseWriter, r *http.Request, UUID uuid.UUID, username, email string, hashedPassword []byte) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("AddUser failure to open database : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("AddUser failure to start transaction : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// execute transaction
	_, err = tx.Exec("INSERT INTO users (UUID, username, email, password) VALUES (?, ?, ?, ?)", UUID, username, email, hashedPassword)
	if err != nil {
		return fmt.Errorf("AddUser failure to execute transaction : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("AddUser failure to commit transaction : %w ", err)
	}

	return nil
}

// Add Post to Database
func InsertPost(w http.ResponseWriter, r *http.Request) error {
	// variable Declaration
	var checkedCategories []string

	username, err := GetUsername(r)
	if err != nil {
		return fmt.Errorf("InsertPost failure to get username : %w ", err)
	} // get username
	postContent := r.FormValue("postContent") // get the content of the post by Form

	for i := 0; i < 5; i++ {
		categorieName := r.FormValue("category-" + strconv.Itoa(i))
		if categorieName != "" {
			checkedCategories = append(checkedCategories, categorieName)
		}
	}

	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("InsertPost fail to open db : %w", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("InsertPost fail to start transaction : %w", err)
	}

	// defer test for failing commit
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()
	// add depend of the number of categories
	switch len(checkedCategories) {
	case 1:
		InserPost_OneCategories(tx, username, postContent, checkedCategories)
	case 2:
		InserPost_TwoCategories(tx, username, postContent, checkedCategories)
	case 3:
		InserPost_threeCategories(tx, username, postContent, checkedCategories)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("InsertPost fail to commit transaction : %w ", err)
	}

	return nil
}

// Add Comment to Database
func InsertComment(w http.ResponseWriter, r *http.Request) error {
	// get the username and the time of the comment
	username, err := GetUsername(r)
	if err != nil {
		return fmt.Errorf("InsertComment failure to get username : %w ", err)
	}
	CommenteTime := time.Now().Format("2006-01-02 : 15:04")

	// get value by the Form
	commentContent := r.FormValue("comContent")
	postID, _ := strconv.Atoi(r.FormValue("new-comment-2"))

	// Exec for Comment
	exec := "INSERT INTO comments (post_id, comment, auteur, date) VALUES (?,?,?,?)"

	// open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("InsertComment failure to open database : %w ", err)
	}
	defer db.Close()

	// start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("InsertComment start transaction : %w ", err)
	}

	// defer test for failure of commit transaction
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(exec, postID, commentContent, username, CommenteTime)
	if err != nil {
		return fmt.Errorf("InsertComment insert in comment table : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("InsertComment fail to commit : %w ", err)
	}
	err = UpdateCommentNumber(postID)
	if err != nil {
		return fmt.Errorf("insert Comment failure for increase number of comment : %w ", err)
	}

	return nil
}

func addLike(postid int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("addLike fail to open database: %w ", err)
	}
	defer db.Close()

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("AddLike failure to start transaction : %w ", err)
	}

	// defer fail Commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// Insert like in Database
	_, err = tx.Exec("INSERT INTO like (post_id, username, like, dislike) VALUES (?,?,?,?) ", postid, username, 1, 0)
	if err != nil {
		return fmt.Errorf("AddLike failure to insert like into like table : %w ", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("AddLike failure to commit : %w ", err)
	}
	return nil
}

func addDislike(postid int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("addDislike fail to open database: %w ", err)
	}
	defer db.Close()

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("AddDislike failure to start transaction : %w ", err)
	}

	// defer fail Commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// Insert Dislike in Database
	_, err = tx.Exec("INSERT INTO like (post_id, username, like, dislike) VALUES (?,?,?,?)", postid, username, 0, 1)
	if err != nil {
		return fmt.Errorf("AddDislike failure to insert like into like table : %w ", err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("AddDislike failure to commit : %w ", err)
	}

	return nil
}

// add comment like to database
func addCommentLike(comment_id int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("addCommentLike fail to open database: %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("AddcommentLike fail to start transaction : %w ", err)
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// insert like
	_, err = tx.Exec("INSERT INTO commentlike (comment_id, username, like, dislike) VALUES (?,?,?,?) ", comment_id, username, 1, 0)
	if err != nil {
		return fmt.Errorf("AddcommentLike fail to insert like in comment like table : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("AddcommentLike fail to commit transaction : %w ", err)
	}
	return nil
}

// Add dislike of comment to database
func addCommentDislike(comment_id int, username string) error {
	//	Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("addCommentDislike fail to open database: %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("addCommentDislike failure to start transaction : %w ", err)
	}
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("INSERT INTO commentlike (comment_id, username, like, dislike) VALUES (?,?,?,?)", comment_id, username, 0, 1)
	if err != nil {
		return fmt.Errorf("addCommentDislike failure to insert commentlike in database (tx.Query) : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("addCommentDislike failure to commit : %w ", err)
	}
	return nil
}
