package forum

import "fmt"

func Data_HandlerForum() ([]Posts, error) {
	posts, err := GetAllPost()
	if err != nil {
		return nil, fmt.Errorf("Data_HandlerForum failure to Get All Posts : %w ", err)
	}

	for i := range posts {

		like, err := GetLikePost(posts[i].Id)
		if err != nil {
			return nil, fmt.Errorf("Data_HandlerForum failure to get like of post : %w ", err)
		}
		posts[i] = CountLikePost(posts[i], like)

		coms, err := GetComment(posts[i].Id)
		if err != nil {
			return nil, fmt.Errorf("Data_HandlerForum failure to get comment : %w ", err)
		}

		for z := range coms {
			like, err := GetLikeComment(coms[z].Id)
			if err != nil {
				return nil, fmt.Errorf("Data_HandlerForum failure to get like of comment : %w ", err)
			}
			coms[z] = CountLikeComment(coms[z], like)
		}
		posts[i].Comment = coms
		posts[i].CommentNum = len(coms)
	}
	return posts, nil
}

func Data_HandlerResults(Catagories []string) ([]Posts, error) {
	var posts []Posts

	post, err := GetAllPost()
	if err != nil {
		return nil, fmt.Errorf("Data_HandlerResults failure to get all posts : %w ", err)
	}
	for i := range post {
		for z := range Catagories {
			if post[i].Category1 == Catagories[z] || post[i].Category2 == Catagories[z] || post[i].Category3 == Catagories[z] {
				posts = append(posts, post[i])
				break
			}
		}
	}
	if len(posts) != 0 {
		for i := range posts {
			like, err := GetLikePost(posts[i].Id)
			if err != nil {
				return nil, fmt.Errorf("Data_HandlerResults failure to get like of post : %w ", err)
			}
			posts[i] = CountLikePost(posts[i], like)

			coms, err := GetComment(posts[i].Id)
			if err != nil {
				return nil, fmt.Errorf("Data_HandlerResults failure to get comment : %w ", err)
			}

			for z := range coms {
				like, err := GetLikeComment(coms[z].Id)
				if err != nil {
					return nil, fmt.Errorf("Data_HandlerResults failure to get like of comment : %w ", err)
				}
				coms[z] = CountLikeComment(coms[z], like)
			}
			posts[i].Comment = coms
			posts[i].CommentNum = len(coms)
		}
		return posts, nil
	}
	return nil, fmt.Errorf("Data_HandlerResults no post valid : %w ", err)
}

func Data_HandlerFilter_Post(username string) ([]Posts, error) {
	var posts []Posts
	var SendPost []Posts

	posts, err := GetAllPost()
	if err != nil {
		return nil, fmt.Errorf("Data_HandlerFilter_Post failure to get all post %w ", err)
	}

	for i := range posts {
		if posts[i].UserName == username {
			like, err := GetLikePost(posts[i].Id)
			if err != nil {
				return nil, fmt.Errorf("Data_HandlerResults failure to get like of post : %w ", err)
			}
			posts[i] = CountLikePost(posts[i], like)

			coms, err := GetComment(posts[i].Id)
			if err != nil {
				return nil, fmt.Errorf("Data_HandlerResults failure to get comment : %w ", err)
			}

			for z := range coms {
				like, err := GetLikeComment(coms[z].Id)
				if err != nil {
					return nil, fmt.Errorf("Data_HandlerResults failure to get like of comment : %w ", err)
				}
				coms[z] = CountLikeComment(coms[z], like)
			}
			posts[i].Comment = coms
			posts[i].CommentNum = len(coms)

			SendPost = append(SendPost, posts[i])
		}
	}

	return SendPost, nil
}

func Data_HandlerFilter_Like(username string) ([]Posts, error) {
	var posts []Posts

	like, err := GetLikeUser(username)
	if err != nil {
		return nil, fmt.Errorf("Data_HandlerFilter_Like failure to get like of  the user : %w ", err)
	}

	for i := range like {
		post, err := GetPost(like[i].PostId)
		if err != nil {
			return nil, fmt.Errorf("Data_HandlerFilter_Like failure to get post of  the user : %w ", err)
		}
		posts = append(posts, post)
	}

	for i := range posts {
		like, err := GetLikePost(posts[i].Id)
		if err != nil {
			return nil, fmt.Errorf("Data_HandlerFilter_Like failure to get like of post : %w ", err)
		}
		posts[i] = CountLikePost(posts[i], like)

		coms, err := GetComment(posts[i].Id)
		if err != nil {
			return nil, fmt.Errorf("Data_HandlerFilter_Like failure to get comment : %w ", err)
		}

		for z := range coms {
			like, err := GetLikeComment(coms[z].Id)
			if err != nil {
				return nil, fmt.Errorf("Data_HandlerFilter_Like failure to get like of comment : %w ", err)
			}
			coms[z] = CountLikeComment(coms[z], like)
		}
		posts[i].Comment = coms
		posts[i].CommentNum = len(coms)
	}

	return posts, nil
}
