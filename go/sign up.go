package fichiers

import (
	"database/sql"
	"net/http"
	adressemail "net/mail"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

//ajoute une entrée dans  "./dataBase/login/login.db"
func ajouter_utilisateur_login(name, mail, password string) {
	database, _ := sql.Open("sqlite3", "./dataBase/login/login.db")
	statement, _ := database.Prepare("INSERT INTO login (name, mail, password) VALUES (?, ?, ?)")
	statement.Exec(name, mail, password)
	database.Close()
	statement.Close()
}

//crée une nouvelle base de donnée dans "./dataBase/utilisateurs/"
func nouvelle_db_utilisateur(name string) {

	database, _ := sql.Open("sqlite3", "./dataBase/utilisateurs/"+name+".db")
	database.Exec("CREATE TABLE messages (id INTEGER PRIMARY KEY, message TEXT, like TEXT, retweet TEXT)")
	database.Exec("CREATE TABLE donnees (id INTEGER PRIMARY KEY, message TEXT, like TEXT, retweet TEXT)")

	database.Exec("CREATE TABLE informations (name TEXT, photo TEXT)")

	statement, _ := database.Prepare("INSERT INTO informations (name,photo) VALUES (?,?)")
	statement.Exec(name, "C:/Users/babouin/Pictures/superlupo.jpg")

	database.Close()
	statement.Close()
}

func hashPassword(password string) (string, error) {

	//fonction du package "golang.org/x/crypto/bcrypt" qui sert a crypter
	//exemple pour le même texte :
	// $2a$14$eZbQVUvVY82GCihf0rnFYu1W2XXjT2sxJcu.83madaRLNoyfCpUe.
	// $2a$14$W5aR2tyyVexndfSbNDgkg.7o8V/K/nTggkB03W/6ypzV6VX.ceO4G

	//bcrypt.GenerateFromPassword(password []byte, cost int)
	//cost est en corélation directe avec la qualité obtenue ainsi que du temps d'obtention
	// environ 1 seconde pour 14 et double a chaque incrémentation
	//connaitre cost est indispensable pour le décryptage
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

//fonction non finie, contient le minimum.
func validerName(name string) bool {

	//si le nom est vide ou comporte des espaces alors c'est faux
	if name == "" || strings.Contains(name, " ") {
		return false
	}

	//ouvrir la base de données
	database, _ := sql.Open("sqlite3", "./dataBase/login/login.db")
	defer database.Close()
	//si la base de donnée n'existe pas alors elle est crée
	database.Exec("CREATE TABLE IF NOT EXISTS login (id INTEGER PRIMARY KEY, name TEXT, mail TEXT, password TEXT)")
	//rechercher si il existe déjà un utilisateur avec ce pseudo
	rows, _ := database.Query("SELECT * FROM login WHERE name = '" + name + "'")

	//si c'est le cas
	if rows.Next() {
		//alors c'est faux
		return false
	}

	//sinon, le pseudonime est validé
	return true
}

//fonction non finie, contient le minimum
func validerMail(mail string) bool {
	//fonction du package "net/mail" qui vérifie si la syntaxe est bonne
	_, err := adressemail.ParseAddress(mail)

	//si il y a une erreur
	if err != nil {
		//alors c'est faux
		return false
	}
	//sinon, l'adressemail est validévalidéé
	//sous toute réserve qu'elle existe
	return true
}

//fonction non finie, contient le minimum
func validerPassword(password, password2 string) bool {

	//si ils sont différents ou vides
	if password != password2 || password == "" {
		//alors c'est faux
		return false
	}

	//rajouter des conditions ex : majuscule, taille, charactéres autorisés ...

	return true
}

//fonction principale de cette page
func ValiderProfil(r *http.Request) Sign_up {

	result := Sign_up{
		Name:             r.PostFormValue("name"),
		Mail:             r.PostFormValue("mail"),
		Password:         r.PostFormValue("password"),
		Bandeau_name:     false,
		Bandeau_mail:     false,
		Bandeau_password: false,
		Valider:          false,
	}

	//cas particulier lors du 1er passage pour ne pas afficher les bandeaux d'erreurs même si tous les champs sont faux
	if result.Name == "" && result.Mail == "" && result.Password == "" {
		return result
	}
	//récupération de la variable pour confirmer le password
	password2 := r.PostFormValue("password2")

	//fmt.Println("\nname =", result.Name, " mail =", result.Mail, " password =", result.Password, " password2 =", password2, ".")

	//vérifier si le pseudo utilisé est bien unique et respecte le format
	result.Bandeau_name = !validerName(result.Name)
	//vérifier si le mail utilisé respecte le format
	result.Bandeau_mail = !validerMail(result.Mail)
	//vérifier si le MDP utilisé respecte le format
	result.Bandeau_password = !validerPassword(result.Password, password2)

	//si l'un des champs est faux
	if result.Bandeau_name || result.Bandeau_mail || result.Bandeau_password {
		//ca va reboucler
		return result
	}
	//sinon, passage à l'étape suivante
	result.Valider = true

	//crypter le MDP
	password2, _ = hashPassword(password2)

	//base de données contenant le minimum pour différencier les différents utilisateurs
	ajouter_utilisateur_login(result.Name, result.Mail, password2)

	//ajouter une base de données propre a chaque utilisateur
	nouvelle_db_utilisateur(result.Name)

	return result
}
