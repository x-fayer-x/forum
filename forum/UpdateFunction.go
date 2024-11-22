package forum

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// update like
func UpdateLike(postid int, username string) error {
	// Open DataBase
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("UpdateLike fail to open database: %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("UpdateLike failure to start transaction : %w ", err)
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// update like in database
	_, err = db.Exec("UPDATE like SET like = like + 1, dislike = dislike - 1 WHERE post_id = ? AND username = ?", postid, username)
	if err != nil {
		return fmt.Errorf("UpdateLike failure to update like and dislike in database : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("UpdateLike failure to commit : %w ", err)
	}

	return nil
}

// update Dislike
func UpdateDislike(postid int, username string) error {
	// open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("updateDislike failure to open Database : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("UpdateDislike failure to start transaction : %w ", err)
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// Update Dislike
	_, err = db.Exec("UPDATE like SET like = like - 1, dislike = dislike + 1 WHERE post_id = ? AND username = ?", postid, username)
	if err != nil {
		return fmt.Errorf("UpdateDislike failure to update Dislike : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("UpdateDislike failure to commit transaction : %w ", err)
	}
	return nil
}

// update Comment Like
func UpdateCommentLike(comment_id int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("UpdateCommentLike failure to open Database : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("UpdateCommentLike failure to start transaction : %w ", err)
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// update comment like
	_, err = db.Exec("UPDATE commentlike SET like = like + 1, dislike = dislike - 1 WHERE comment_id = ? AND username = ?", comment_id, username)
	if err != nil {
		return fmt.Errorf("UpdateCommentLike failure to update comment like : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("UpdateCommentLike failure to commit transaction : %w ", err)
	}
	return nil
}

// Update Comment Dislike
func UpdateCommentDislike(comment_id int, username string) error {
	// Open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("UpdateCommentDislike failure to open Database : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("UpdateCommentDislike failure to start transaction : %w ", err)
	}

	// defer fail commit test
	failCommit := true
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	// update comment dislike
	_, err = db.Exec("UPDATE commentlike SET like = like - 1, dislike = dislike + 1 WHERE comment_id = ? AND username = ?", comment_id, username)
	if err != nil {
		return fmt.Errorf("UpdateCommentDislike failure to update comment dislike : %w ", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("UpdateCommentDislike failure to commit transaction : %w ", err)
	}
	return nil
}

func UpdateName(username string, UUID string) error {
	// open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("update name failure to open db : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("update name failure to start transaction : %w ", err)
	}

	failCommit := false
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("UPDATE users SET username = ? WHERE UUID = ? ", username, UUID)
	if err != nil {
		return fmt.Errorf("update name failure to set usename  : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update name failure to commit transaction : %w ", err)
	}

	return nil
}

func UpdateEmail(email string, UUID string) error {
	// open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("update email failure to open db : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("update email failure to start transaction : %w ", err)
	}

	failCommit := false
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("UPDATE users SET email = ? WHERE UUID = ? ", email, UUID)
	if err != nil {
		return fmt.Errorf("update email failure to set usename  : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("update email failure to commit transaction : %w ", err)
	}

	return nil
}

func UpdatePassWord(password string, UUID string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("update password failure to encrypte password %w ", err)
	}

	// open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("update password failure to open db : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("update password failure to start transaction : %w ", err)
	}

	failCommit := false
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec("UPDATE users SET password = ? WHERE UUID = ? ", hashedPassword, UUID)
	if err != nil {
		return fmt.Errorf("update password failure to set usename  : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		failCommit = false
		return fmt.Errorf("update password failure to commit transaction : %w ", err)
	}

	return nil
}

func UpdateCommentNumber(PostID int) error {
	// open Database
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("update password failure to open db : %w ", err)
	}
	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("update password failure to start transaction : %w ", err)
	}

	failCommit := false
	defer func() {
		if !failCommit {
			tx.Rollback()
		}
	}()

	tx.Exec("UPDATE INTO posts SET com_number += 1 WHERE post_id = ?", PostID)

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("update password failure to commit transaction : %w ", err)
	}

	return nil
}
