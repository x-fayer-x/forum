package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// OpenDB ouvre et retourne la connexion à la base de données SQLite.
func openDB() (*sql.DB, error) {
	// Open database
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return nil, fmt.Errorf("openDB : failure to open the database %w", err)
	}
	return db, nil
}

func createAndSetCookie(w http.ResponseWriter, username string) {
	// Remplacer l'UUID dans la base de données avec un nouveau UUID
	sessionUUID, err := changeUUIDinDB(username)
	if err != nil {
		fmt.Println(err)
	}

	// load timezone
	loc, _ := time.LoadLocation("Europe/Paris")

	fmt.Println("dans createandsetcookie",sessionUUID)
	// create new cookie for the users
	cookie := &http.Cookie{
		Name:     "UserToken",
		Value:    sessionUUID.String(),
		Expires:  time.Now().In(loc).Add(1 * time.Hour),
		HttpOnly: true,
	}

	// set the cookie in the responce
	http.SetCookie(w, cookie)
}

// Check if there is a valide cookie in the browser
func checkCookie(r *http.Request) (bool, error) {
	var dbUUID string
	var timeOut time.Time
	var CorrectUUId bool

	// Get the time
	timeOut = time.Now()

	// Get the cookie from the user
	cookie, err := r.Cookie("UserToken")
	if err != nil {
		return false, nil
	}

	// Get the value of the cookie
	cookieUUID := cookie.Value

	// Open database
	db, err := openDB()
	if err != nil {
		return false, fmt.Errorf("CheckCookie failure to open db : %w ", err)
	}
	defer db.Close()

	// Open transction in database
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error in CheckCookie : failure to open transaction (db.begin), ", err)
	}
	commit := false
	defer func() {
		if commit {
			tx.Rollback()
		}
	}()

	// query the UUID in users tabel
	rows, err := tx.Query("SELECT UUID FROM users")
	if err != nil {
		fmt.Println("error in CheckCookie : failure to exctract uuid (tx.Query), ", err)
	}

	// loop on the next line and check if a UUID from the database match with the UUID from the cookie
	for rows.Next() {

		err = rows.Scan(&dbUUID)
		if err != nil {
			fmt.Println("error in CheckCookie : failure to retrieve user uuid, ", err)
		}
		if dbUUID == cookieUUID {
			// end of the transaction
			CorrectUUId = true
			break
		}
	}
	// end of the transaction
	err = tx.Commit()
	if err != nil {
		commit = true
		fmt.Println("error in  CheckCookie : failure to Commit transaction (tx.Commit), ", err)
	}

	// check if the expiration time is not past
	if CorrectUUId && cookie.Expires.Before(timeOut) {
		return true, nil
	}

	return false, nil
}

// check if the password from the db and that from the form match
func matchPassWord(passwordDB, passwordForm string) error {
	// use bcrypt package function to check if the 2 password are matching
	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordForm))
	if err != nil {
		fmt.Println("error in matchPassWord : password are not matching, ", err)
		return err
	}

	return nil
}

// change the UUID for a user in database
func changeUUIDinDB(username string) (uuid.UUID, error) {
	// Generate a new UUID
	sessionUUID, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, fmt.Errorf("ChangeUUIDinDB failure to generate new uuid %w ", err)
	}

	// Open database
	db, err := openDB()
	if err != nil {
		return uuid.Nil, fmt.Errorf("ChangeUUIDindb failure to open db : %w ", err)
	}
	defer db.Close()

	// Open a transaction
	tx, err := db.Begin()
	if err != nil {
		return uuid.Nil, fmt.Errorf("ChangeUUIDinDB failure to start transaction : %w ", err)
	}
	commit := false
	defer func() {
		if commit {
			tx.Rollback()
		}
	}()

	// exec the transaction
	_, err = tx.Exec("UPDATE users SET UUID = ? WHERE username = ?", sessionUUID, username)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ChangeUUIDinDB failure to exec transaction : %w ", err)
	}

	err = tx.Commit()
	if err != nil {
		commit = true

		return uuid.Nil, fmt.Errorf("changeUUIDinDB failure to commit transaction %w ", err)
	}

	return sessionUUID, nil
}

// HandleLogout gère la déconnexion de l'utilisateur en supprimant le cookie de session.
func HandleLogout(w http.ResponseWriter) {
	// Créer un nouveau cookie avec le nom "UserToken", une valeur vide, et une expiration immédiate pour supprimer le cookie
	cookie := &http.Cookie{
		Name:    "UserToken",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour), // Expiration immédiate
	}

	// Définir le cookie dans la réponse HTTP pour supprimer le cookie de session du navigateur
	http.SetCookie(w, cookie)
}

// check if the Url match with one of the page
func CheckURl(r *http.Request) bool {
	// variable
	var good bool = true
	var notgood bool = false

	// switch for the path
	switch r.URL.Path {
	case "/":
		return good

	case "/home":
		return good

	case "/team":
		return good

	case "/teamin":
		return good

	case "/forum":
		return good

	case "/forumin":
		return good

	case "/edit-post":
		return good

	case "/edit-comment":
		return good

	case "/login":
		return good

	case "/results":
		return good

	case "/like":
		return good

	case "/dislike":
		return good

	case "/users":
		return good
	case "/commentslike":
		return good
	case "/commentsdislike":
		return good
	case "/UserFilter":
		return good
	}
	return notgood
}
