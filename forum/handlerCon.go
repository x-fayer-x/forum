package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Handlerin gère les requêtes pour la page Home Connecté
func Handlerin(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodGet {

			test, err := checkCookie(r)

			// Vérifier si l'utilisateur est connecté en vérifiant le cookie UserToken
			if !test || err != nil {
				// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}

			// Charger le modèle de la page d'accueil connectée
			t, err := template.ParseFiles("./assets/templates/indexLogin.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			t.Execute(w, t)
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// HandleForumin gère les requêtes pour la page Forum Connecté
func HandleForumin(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		test, err := checkCookie(r)

		// Vérifier si l'utilisateur est connecté en vérifiant le cookie UserToken
		if !test || err != nil {
			// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if r.Method == http.MethodPost {
			// Traitement du formulaire de publication de message

			InsertPost(w, r)

			http.Redirect(w, r, "/forumin", http.StatusSeeOther)

		} else if r.Method == http.MethodGet {
			// Affichage de la page Forum Connecté avec les messages

			// Ouvrir la base de données
			data, err := Data_HandlerForum()
			if err != nil {
				fmt.Println(err)
			}

			// Charger le modèle de la page Forum Connecté avec les messages
			t, err := template.ParseFiles("./assets/templates/forumLogin.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			t.Execute(w, data)
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, t)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// HandleTeamin gère les requêtes pour la page Team Connecté
func HandleTeamin(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodGet {
			// Vérifier si l'utilisateur est connecté en vérifiant le cookie UserToken
			ut, err := r.Cookie("UserToken")
			TimeOut := time.Now()

			if err != nil || ut.Value == "" || ut.Expires.After(TimeOut) {
				// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}

			// Charger le modèle de la page Team Connecté
			t, err := template.ParseFiles("./assets/templates/teamIn.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			t.Execute(w, t)
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// HandleEdit gère les requêtes pour la page d'édition de message Connecté
func HandleEdit(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodGet {
			// Vérifier si l'utilisateur est connecté en vérifiant le cookie UserToken
			ut, err := r.Cookie("UserToken")
			TimeOut := time.Now()

			if err != nil || ut.Value == "" || ut.Expires.After(TimeOut) {
				// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}

			// Charger le modèle de la page d'édition de message
			t, err := template.ParseFiles("./assets/templates/edit-post.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			t.Execute(w, nil)
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// HandleComment gère les requêtes pour la page de commentaire
func HandleComment(w http.ResponseWriter, r *http.Request) {
	// Initialiser les variables

	// Vérifier le chemin de l'URL
	if CheckURl(r) {

		// Vérifier si l'utilisateur est connecté en vérifiant le cookie UserToken
		ut, err := r.Cookie("UserToken")
		TimeOut := time.Now()

		if err != nil || ut.Value == "" || ut.Expires.After(TimeOut) {
			// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if r.Method == http.MethodPost {
			// Traitement du formulaire de commentaire

			// Récupérer les valeurs du formulaire
			forumpage := r.FormValue("new-comment-1")
			editpage := r.FormValue("new-comment-2")

			if forumpage != "" {
				// convert postid from string to int
				temp, _ := strconv.Atoi(forumpage)

				// retrieve post
				data, err := GetPost(temp)
				if err != nil {
					fmt.Println(err)
				}

				// Charger le modèle de la page de commentaire
				t, err := template.ParseFiles("./assets/templates/Comment.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				t.Execute(w, data)

			} else if editpage != "" {
				err = InsertComment(w, r)
				if err != nil {
					fmt.Println(err)
				}
				http.Redirect(w, r, "/forumin", http.StatusSeeOther)
			}
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodPost {
			postId := r.FormValue("postId")
			username, err := GetUsername(r)
			if err != nil {
				fmt.Println(err)
			}
			postIdInt, err := strconv.Atoi(postId)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}

			db, err := openDB()
			if err != nil {
				fmt.Println(err)
			}
			defer db.Close()

			var liked LikePost

			err = db.QueryRow("SELECT like, dislike FROM like WHERE post_id = ? AND username = ?", postIdInt, username).Scan(&liked.Like, &liked.Dislike)
			if err == sql.ErrNoRows {
				addLike(postIdInt, username)
			}

			if liked.Like == 1 {
				subLike(postIdInt, username)
			} else if liked.Like == 0 && liked.Dislike == 1 {
				UpdateLike(postIdInt, username)
			}

			http.Redirect(w, r, "/forumin", http.StatusSeeOther)
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DislikePostHandler(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodPost {
			postId := r.FormValue("postId")
			username, err := GetUsername(r)
			if err != nil {
				fmt.Println(err)
			}
			postIdInt, err := strconv.Atoi(postId)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}

			db, err := openDB()
			if err != nil {
				fmt.Println(err)
			}
			defer db.Close()

			var liked LikePost

			err = db.QueryRow("SELECT  like, dislike FROM like WHERE post_id = ? AND username = ?", postIdInt, username).Scan(&liked.Like, &liked.Dislike)
			if err == sql.ErrNoRows {
				addDislike(postIdInt, username)
			}

			if liked.Dislike == 1 {
				subDislike(postIdInt, username)
			} else if liked.Like == 1 && liked.Dislike == 0 {
				UpdateDislike(postIdInt, username)
			}

			http.Redirect(w, r, "/forumin", http.StatusSeeOther)

		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodGet {
			username, err := GetUsername(r)
			if err != nil {
				fmt.Println(err)
			}
			t, err := template.ParseFiles("./assets/templates/user.html")
			if err != nil {
				fmt.Println("error for parse file")
			}
			t.Execute(w, username)
		} else if r.Method == http.MethodPost {
			name := r.FormValue("username")
			email := r.FormValue("email")
			password := r.FormValue("password")
			cookie, _ := r.Cookie("UserToken")
			uuid := cookie.Value

			username, err := GetUsername(r)
			if err != nil {
				fmt.Println(err)
			}

			if name != "" {
				username, err := GetUsername(r)
				if err != nil {
					fmt.Println(err, uuid)
				}
				if username != name {
					UpdateName(name, uuid)
				}
			}
			if email != "" {
				mail, err := GetEmail(r)
				if err != nil {
					fmt.Println(err)
				}
				if mail != email {
					err := UpdateEmail(email, uuid)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			if password != "" {
				mdp, err := GetPassWord(r, username)
				if err != nil {
					fmt.Println(err)
				}

				if bcrypt.CompareHashAndPassword([]byte(mdp), []byte(password)) != nil {
					err = UpdatePassWord(password, uuid)
					if err != nil {
						fmt.Println(err)
					}
					HandleLogout(w)
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}

			t, err := template.ParseFiles("./assets/templates/user.html")
			if err != nil {
				fmt.Println("error for parse file")
			}
			t.Execute(w, username)
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {

		ut, err := r.Cookie("UserToken")
		TimeOut := time.Now()

		if err != nil || ut.Value == "" || ut.Expires.After(TimeOut) {
			// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		if r.Method == http.MethodPost {
			comment_id := r.FormValue("CommentId")
			username, err := GetUsername(r)
			if err != nil {
				fmt.Println(err)
			}
			commentIdInt, err := strconv.Atoi(comment_id)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}
			db, err := openDB()
			if err != nil {
				fmt.Println(err)
			}
			defer db.Close()

			var liked LikeComment

			err = db.QueryRow("SELECT  like, dislike FROM commentlike WHERE comment_id = ? AND username = ?", commentIdInt, username).Scan(&liked.Like, &liked.Dislike)
			if err == sql.ErrNoRows {
				addCommentLike(commentIdInt, username)
			}

			if liked.Like == 1 {
				subCommentLike(commentIdInt, username)
			} else if liked.Like == 0 && liked.Dislike == 1 {
				UpdateCommentLike(commentIdInt, username)
			}

			http.Redirect(w, r, "/forumin", http.StatusSeeOther)

		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func DislikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {

		ut, err := r.Cookie("UserToken")
		TimeOut := time.Now()

		if err != nil || ut.Value == "" || ut.Expires.After(TimeOut) {
			// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		if r.Method == http.MethodPost {
			comment_id := r.FormValue("CommentId")
			username, err := GetUsername(r)
			if err != nil {
				fmt.Println(err)
			}
			commentIdInt, err := strconv.Atoi(comment_id)
			if err != nil {
				http.Error(w, "Invalid post ID", http.StatusBadRequest)
				return
			}
			db, err := openDB()
			if err != nil {
				fmt.Println(err)
			}
			defer db.Close()

			var liked LikeComment

			err = db.QueryRow("SELECT  like, dislike FROM commentlike WHERE comment_id = ? AND username = ?", commentIdInt, username).Scan(&liked.Like, &liked.Dislike)
			if err == sql.ErrNoRows {
				addCommentDislike(commentIdInt, username)
			}

			if liked.Dislike == 1 {
				subCommentDislike(commentIdInt, username)
			} else if liked.Like == 1 && liked.Dislike == 0 {
				UpdateCommentDislike(commentIdInt, username)
			}

			http.Redirect(w, r, "/forumin", http.StatusSeeOther)

		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Handlerfilter(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {

		ut, err := r.Cookie("UserToken")
		TimeOut := time.Now()

		if err != nil || ut.Value == "" || ut.Expires.After(TimeOut) {
			// Rediriger vers la page d'accueil si l'utilisateur n'est pas connecté
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		if r.Method == http.MethodPost {

			value := r.FormValue("filteruser")

			username, err := GetUsername(r)
			if err != nil {
				fmt.Println(err)
			}

			switch value {
			case "Mypost":
				data, err := Data_HandlerFilter_Post(username)
				if err != nil {
					fmt.Println(err)
				}

				t, err := template.ParseFiles("./assets/templates/UserData.html")
				if err != nil {
					fmt.Println("error in filterHandler failure for found html file ", err)
				}
				t.Execute(w, data)

			case "likedpost":
				data, err := Data_HandlerFilter_Like(username)
				if err != nil {
					fmt.Println(err)
				}

				t, err := template.ParseFiles("./assets/templates/UserData.html")
				if err != nil {
					fmt.Println("error in filterHandler failure to found html file")
				}
				t.Execute(w, data)
			case "Comment":
			}
		} else {
			t, err := template.ParseFiles("./assets/templates/400.html")
			if err != nil {
				fmt.Println(err)
			}
			t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
