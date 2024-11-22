package forum

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Handler gère les requêtes GET et POST pour la page d'accueil.
func Handler(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodGet {
			// Gestion des requêtes GET pour la page d'accueil
			test, err := checkCookie(r)
			if test && err == nil {
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}

			t, err := template.ParseFiles("./assets/templates/index.html")
			if err != nil {
				http.Error(w, "Erreur de chargement de la page", http.StatusInternalServerError)
				return
			}

			t.Execute(w, t)

		} else if r.Method == http.MethodPost {
			// Gestion des requêtes POST pour la page d'accueil

			// Récupération des infos du formulaire
			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")

			// Génération d'un UUID pour l'utilisateur
			userUUID, _ := uuid.NewV4()

			// Création d'un cookie avec l'UUID
			cookie := http.Cookie{
				Name:     "UserToken",
				Value:    userUUID.String(),
				Path:     "/",
				Expires:  time.Now().Add(1 * time.Hour),
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			// Hachage du mot de passe
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Mot de passe non sécurisé", http.StatusInternalServerError)
				return
			}

			// Insertion des infos du formulaire dans la BDD
			AddUser(w, r, userUUID, name, email, hashedPassword)

			// Redirection vers la page du forum
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

// HandleForum gère les requêtes pour la page du forum.
func HandleForum(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodGet {

			// Gestion des requêtes pour la page du forum
			test, err := checkCookie(r)
			if test && err == nil {
				http.Redirect(w, r, "/forumin", http.StatusSeeOther)
			}

			data, err := Data_HandlerForum()
			if err != nil {
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			}

			// Chargement du modèle forum.html avec les messages récupérés
			t, err := template.ParseFiles("./assets/templates/forum.html")
			if err != nil {
				http.Error(w, "Erreur de chargement de la page", http.StatusInternalServerError)
				return
			}
			t.Execute(w, data)
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

// HandleLogin gère les requêtes GET et POST pour la page de connexion.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodPost {

			// Gestion des requêtes POST pour la page de connexion

			// Récupération des données du formulaire
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Récupération du mot de passe haché de l'utilisateur depuis la base de données
			hashedPassword, err := GetPassWord(r, username)
			if err != nil {
				fmt.Println(err)
			}
			if hashedPassword == "" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}

			// Comparaison du mot de passe saisi avec le mot de passe haché
			err = matchPassWord(hashedPassword, password)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusUnauthorized)
			}

			// Création du cookie de session
			createAndSetCookie(w, username)

			// Redirection vers la page du forum
			http.Redirect(w, r, "/forumin", http.StatusSeeOther)
		} else if r.Method == http.MethodGet {
			// Gestion des requêtes GET pour la page de connexion

			// Déconnexion de l'utilisateur
			HandleLogout(w)

			// Affichage du formulaire de connexion
			t, err := template.ParseFiles("./assets/templates/login.html")
			if err != nil {
				http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
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

// HandleTeam gère les requêtes pour la page de l'équipe.
func HandleTeam(w http.ResponseWriter, r *http.Request) {
	if CheckURl(r) {
		if r.Method == http.MethodGet {
			test, err := checkCookie(r)
			if test && err == nil {
				http.Redirect(w, r, "/teamin", http.StatusSeeOther)
			}

			t, err := template.ParseFiles("./assets/templates/team.html")
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

// HandleResults gère les requêtes pour la page des résultats.
func HandleResults(w http.ResponseWriter, r *http.Request) {
	// Tableau pour stocker les catégories sélectionnées par l'utilisateur
	if CheckURl(r) {
		if r.Method == http.MethodPost {
			var t *template.Template
			var selectedCategories []string

			test, err := checkCookie(r)

			// Gestion des requêtes GET pour la page d'accueil
			if test && err == nil {
				tmpl, err := template.ParseFiles("./assets/templates/resultsin.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				t = tmpl
			} else {
				tmpl, err := template.ParseFiles("./assets/templates/results.html")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				t = tmpl
			}

			for i := 0; i < 5; i++ {
				categoryName := r.FormValue("category-" + strconv.Itoa(i))

				// Ajouter la catégorie non vide au tableau
				if categoryName != "" {
					selectedCategories = append(selectedCategories, categoryName)
				}
			}
			// retrieve data in db
			data, err := Data_HandlerResults(selectedCategories)
			if err != nil {
				fmt.Println(err)
			}

			// Passer les messages filtrés au modèle results.html
			t.Execute(w, data)
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
