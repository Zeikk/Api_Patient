package control

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	jwt "github.com/dgrijalva/jwt-go"
	db "api_patient/db"
	"log"
)

func getId(w http.ResponseWriter, r *http.Request) string{

	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintf(w, "\nNécessite Authentification")
			w.WriteHeader(http.StatusUnauthorized)
			return ""
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "\nNécessite Authentification")
		return ""
	}

	tokenStr := c.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("grain_de_sel"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Fprintf(w, "\nNécessite Authentification")
			w.WriteHeader(http.StatusUnauthorized)
			return ""
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "\nNécessite Authentification")
		return ""
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "\nNécessite Authentification")
		return ""
	}

	log.Println(claims.IdMedecin)
	return claims.IdMedecin
}	

func GetPatients(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Affichage Patients")
	
	idMedecin := getId(w, r)
	db := db.OpenDB()

	rows, err := db.Query("select prenomPersonne, nomPersonne, sexePersonne, dateDeNaissance from personne join etre_patient using(idPersonne) join suivre using(idPatient) where idMedecin = ?", idMedecin)

	defer rows.Close()

	var nom, prenom, date, sexe string
	for rows.Next() {
		err := rows.Scan(&prenom, &nom, &sexe, &date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "\nNom: %s", nom)
		fmt.Fprintf(w, "\nPrénom: %s", prenom)
		fmt.Fprintf(w, "\nDate de Naissance: %s", date)
		fmt.Fprintf(w, "\nSexe: %s", sexe)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()

}

func GetPatientById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idPatient := vars["patientId"]

	idMedecin := getId(w, r)
		
	db := db.OpenDB()

	rows, err := db.Query("select prenomPersonne, nomPersonne, sexePersonne, dateDeNaissance from personne join etre_patient using(idPersonne) join suivre using(idPatient) where idPatient = ? and idMedecin = ?", idPatient, idMedecin)

	defer rows.Close()

	var nom, prenom, date, sexe string
	for rows.Next() {
		err := rows.Scan(&prenom, &nom, &sexe, &date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "\nNom: %s", nom)
		fmt.Fprintf(w, "\nPrénom: %s", prenom)
		fmt.Fprintf(w, "\nDate de Naissance: %s", date)
		fmt.Fprintf(w, "\nSexe: %s", sexe)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}