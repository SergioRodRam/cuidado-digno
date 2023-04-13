package main

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	"net/http"

	//"text/template"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Host := "whiteshark.mysql.database.azure.com:3306"
	Usuario := "whiteshark"
	Contraseña := "alberto2023$%"
	BaseDeDatos := "whitesharkbd"

	//parseTime=True (Para pasar datos tipo date de la base de datos)
	conexion, err := sql.Open(Driver, Usuario+":"+Contraseña+"@tcp("+Host+")/"+BaseDeDatos+"?parseTime=True")
	if err != nil {
		panic(err.Error())
	}

	return conexion
}

type Receta struct {
	IdReceta     int64
	Receta       string 
}

var plantillas = template.Must(template.ParseGlob("plantillas/*.html"))

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	
	http.HandleFunc("/", MostrarHome)
	http.HandleFunc("/home", MostrarHome)
	http.HandleFunc("/checate", MostrarChecate)
	http.HandleFunc("/registro", MostrarRegistro)
	http.HandleFunc("/contacto", MostrarContacto)
	http.HandleFunc("/dietas", MostrarDietas)

	fmt.Printf("Servidor corriendo ...\n")

	http.ListenAndServe(":8080", nil)
}

func MostrarHome(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "home", nil) //nil
}

func MostrarChecate(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "checate", nil) //nil
}

func MostrarRegistro(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "registro", nil) //nil
}

func MostrarContacto(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "contacto", nil) //nil
}

func MostrarDietas(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := conexionBD()

	recetas, err := conexionEstablecida.Query("SELECT * FROM whitesharkbd.receta_chatgpt WHERE id_receta=1;")

	if err != nil {
		panic(err.Error())
	}

	receta := Receta{}
	arrRecetas := []Receta{}

	for recetas.Next() {

		err = recetas.Scan(&receta.IdReceta, &receta.Receta)

		if err != nil {
			panic(err.Error())
		}

		arrRecetas = append(arrRecetas,receta)
	}

	//json.NewEncoder(w).Encode(receta)

	defer conexionEstablecida.Close()

	plantillas.ExecuteTemplate(w, "dietas", arrRecetas) //nil
}
