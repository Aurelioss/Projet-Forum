package fichiers

import (
	"fmt"
	"net/http"
)

func Modifier_profil(r *http.Request, donnees Donnees) (Donnees, Nom_provisoire) {

	name := ""
	photo := ""

	choix := Nom_provisoire{
		Name: name,
		V1: r.PostFormValue("v1"),
		V2: r.PostFormValue("v2"),
	}

	rows, _ := donnees.Base_de_donnees_du_profil.Query("SELECT * FROM informations")

	if rows.Next() {
		rows.Scan(&name, &photo)
	}

	fmt.Println(name, "   ", photo)
	return donnees, choix
}
