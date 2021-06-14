package fichiers

import "database/sql"

type Donnees struct {
	Base_de_donnees_du_profil *sql.DB
	Connecte                  bool
	Name                      string
	Mail                      string
	password                  string
	Photo                     string
	Choix                     Nom_provisoire
}

type Nom_provisoire struct {
	Name     string
	Mail     string
	Password string
	V1       string
	V2       string
}

type Sign_up struct {
	Name             string
	Mail             string
	Password         string
	Bandeau_name     bool
	Bandeau_mail     bool
	Bandeau_password bool
	Valider          bool
}
