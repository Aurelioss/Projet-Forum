package fichiers

import (
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoadProfil(r *http.Request) (Donnees, string) {


	result := Donnees{
		Base_de_donnees_du_profil: nil,
		Connecte:                  false,
		Name:                      "",
		Photo:                     "",
	}

	name := r.PostFormValue("name")
	password := r.PostFormValue("password")

	fmt.Println("name = ", name, " password = ", password, ".")

	if name == "" && password == "" {
		return result, ""
	}

	database, _ := sql.Open("sqlite3", "./database/login/login.db")

	rows, _ := database.Query("SELECT password FROM login WHERE name = '" + name + "'")

	defer database.Close()
	defer rows.Close()

	password2 := ""

	if rows.Next() {
		rows.Scan(&password2)
	} else {
		return result, name
	}

	fmt.Println("p2 = ", password2)

	if CheckPasswordHash(password, password2) {

		database, _ := sql.Open("sqlite3", "./dataBase/utilisateurs/"+name+".db")

		rows, _ := database.Query("SELECT * FROM informations")

		if rows.Next() {
			rows.Scan(&result.Name, &result.Photo)
		}

		result.Base_de_donnees_du_profil = database
		result.Connecte = true

		fmt.Println("???", result.Base_de_donnees_du_profil)

		return result, ""
	}

	return result, name
}
