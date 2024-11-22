package forum

import "fmt"

// Sub like to database
func subLike(postid int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("subLike fail to open database: %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("subLike failure to start transaction : %w ", err)
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// delete like
	_, err = tx.Exec("DELETE FROM like WHERE post_id = ? AND username = ?", postid, username)
	if err != nil {
		return fmt.Errorf("subLike failure to delete like in database (tx.Query) : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("subLike failure to commit : %w ", err)
	}

	return nil
}

// Sub Dislike to database
func subDislike(postid int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("subDislike fail to open database: %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("subDislike failure to start transaction : %w ", err)
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// delete dislike
	_, err = tx.Exec("DELETE FROM like WHERE post_id = ? AND username = ?", postid, username)
	if err != nil {
		return fmt.Errorf("subDislike failure to delete dislike in database (tx.Query) : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("subDislike failure to commit : %w ", err)
	}
	return nil
}

// Sub like of comment to database
func subCommentLike(comment_id int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("subCommentLike fail to open database: %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("subCommentLike failure to start transaction : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("DELETE FROM commentlike WHERE comment_id = ? AND username = ?", comment_id, username)
	if err != nil {
		return fmt.Errorf("subCommentLike failure to delete commentlike in database (tx.Query) : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("subCommentLike failure to commit : %w ", err)
	}
	return nil
}

// Sub dislike of comment to database
func subCommentDislike(comment_id int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("subCommentDislike fail to open database: %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("subCommentDislike failure to start transaction : %w ", err)
	}

	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// delete dislike
	_, err = tx.Exec("DELETE FROM commentlike WHERE comment_id = ? AND username = ?", comment_id, username)
	if err != nil {
		return fmt.Errorf("subCommentDislike failure to delete commentlike in database (tx.Query) : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("subCommentDislike failure to commit : %w ", err)

	}

	return nil
}
