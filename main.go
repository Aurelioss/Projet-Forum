package main

import (
	"fmt"
	"html/template"
	"net/http"

	fichiers "./go"
)

func main() {

	donnees := fichiers.Donnees{
		Base_de_donnees_du_profil: nil,
		Connecte:                  false,
		Name:                      "",
		Photo:                     "",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  entrée dans normal.html")

		if donnees.Connecte {
			template.Must(template.ParseFiles("html/normal_connecte.html")).Execute(w, donnees)
		} else {
			template.Must(template.ParseFiles("html/normal_non_connecte.html")).Execute(w, donnees)
		}
	})

	http.HandleFunc("/sign_up", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  entrée dans sign_up.html")

		template.Must(template.ParseFiles("html/sign_up.html")).Execute(w, fichiers.ValiderProfil(r))
	})

	http.HandleFunc("/load_profil", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  entrée dans load_profil.html")
		echec := ""

		if donnees.Connecte {
			http.Redirect(w, r, "/profil", http.StatusFound)
		}

		donnees, echec = fichiers.LoadProfil(r)
		fmt.Println(donnees.Base_de_donnees_du_profil, "  ", donnees.Connecte)

		if donnees.Connecte {
			http.Redirect(w, r, "/profil", http.StatusFound)
		}

		fmt.Println("fin")
		template.Must(template.ParseFiles("html/load_profil.html")).Execute(w, echec)

	})

	http.HandleFunc("/profil", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  entrée dans profil.html")
		fmt.Println("1-----------------------------------------")
		choix := fichiers.Nom_provisoire{}
		fmt.Println("2-----------------------------------------")
		donnees, choix = fichiers.Modifier_profil(r, donnees)
		fmt.Println("3-----------------------------------------")
		template.Must(template.ParseFiles("html/profil.html")).Execute(w, choix)
		fmt.Println(donnees, choix)
	})

	http.HandleFunc("/deconnexion", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  entrée dans deconnexion.html")

		donnees = fichiers.Deconnexion(donnees)
		template.Must(template.ParseFiles("html/normal_non_connecte.html")).Execute(w, donnees)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  entrée dans contact.html")

		template.Must(template.ParseFiles("html/contact.html")).Execute(w, donnees)
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("  entrée dans chat.html")

		template.Must(template.ParseFiles("html/chat.html")).Execute(w, donnees)
	})

	fmt.Println("debut")
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	http.ListenAndServe(":80", nil)
}
