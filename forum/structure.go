package forum

// Users représente la structure des utilisateurs dans la base de données.
type Users struct {
	Id       int
	Username string
	Email    string
	Password string
}

// Posts représente la structure des messages dans le forum.
type Posts struct {
	Id         int
	UserName   string
	Content    string
	Category1  string
	Category2  string
	Category3  string
	Category   []string
	Date       string
	Comment    []Comments
	CommentNum int
	Like       int
	Dislike    int
}

// Comments représente la structure des commentaires dans le forum.
type Comments struct {
	Id       int
	PostID   int
	UserName string
	Content  string
	Date     string
	Like     int
	Dislike  int
}

// LikePost représente la structure des likes et dislikes sur les messages.
type LikePost struct {
	PostId   int
	Username string
	Like     int
	Dislike  int
}

type LikeComment struct {
	CommentId int
	Username  string
	Like      int
	Dislike   int
}
