package fichiers

func Deconnexion(donnees Donnees) Donnees {

	donnees.Base_de_donnees_du_profil.Close()
	donnees.Base_de_donnees_du_profil = nil
	donnees.Connecte = false
	donnees.Name = ""
	donnees.Photo = ""

	//fmt.Println(donnees.Base_de_donnees_du_profil.Ping().Error())

	return donnees
}
