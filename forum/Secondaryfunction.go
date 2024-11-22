package forum

import (
	"database/sql"
	"fmt"
	"time"
)

// Insert post with 1 categories
func InserPost_OneCategories(tx *sql.Tx, username string, postContent string, Categories []string) (*sql.Tx, error) {
	postTime := time.Now().Format("2006-01-02 : 15:04") // get the time of post
	exec := "INSERT INTO posts (user_id, content, category1, category2, category3, date) VALUES (?,?,?,?,?,?)"

	_, err := tx.Exec(exec, username, postContent, Categories[0], "", "", postTime)
	if err != nil {
		return tx, fmt.Errorf("InserPost_OneCategories failure to insert post : %w", err)
	}

	return tx, nil
}

// Insert post with 2 categories
func InserPost_TwoCategories(tx *sql.Tx, username string, postContent string, Categories []string) (*sql.Tx, error) {
	postTime := time.Now().Format("2006-01-02 : 15:04") // get the time of post
	exec := "INSERT INTO posts (user_id, content, category1, category2, category3, date) VALUES (?,?,?,?,?,?)"

	_, err := tx.Exec(exec, username, postContent, Categories[0], Categories[1], "", postTime)
	if err != nil {
		return tx, fmt.Errorf("InserPost_OneCategories failure to insert post : %w", err)
	}
	return tx, nil
}

// Insert post with 3 categories
func InserPost_threeCategories(tx *sql.Tx, username string, postContent string, Categories []string) (*sql.Tx, error) {
	postTime := time.Now().Format("2006-01-02 : 15:04") // get the time of post
	exec := "INSERT INTO posts (user_id, content, category1, category2, category3, date) VALUES (?,?,?,?,?,?)"

	_, err := tx.Exec(exec, username, postContent, Categories[0], Categories[1], Categories[2], postTime)
	if err != nil {
		return tx, fmt.Errorf("InserPost_OneCategories failure to insert post : %w", err)
	}

	return tx, nil
}

func CountLikePost(post Posts, like []LikePost) Posts {
	post.Like = 0
	post.Dislike = 0

	for j := 0; j < len(like); j++ {
		if like[j].Like == 1 {
			post.Like += 1
		}
		if like[j].Dislike == 1 {
			post.Dislike += 1
		}
	}
	return post
}

func CountLikeComment(comment Comments, like []LikeComment) Comments {
	comment.Like = 0
	comment.Dislike = 0
	for i := range like {
		if like[i].Like == 1 {
			comment.Like += 1
		}
		if like[i].Dislike == 1 {
			comment.Dislike += 1
		}
	}
	return comment
}
