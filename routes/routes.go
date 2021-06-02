package routes

import (
	"net/http"

	"github.com/JoaoValentimDev/shoop/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/update", controllers.UpdateForm)
	http.HandleFunc("/updateProduct", controllers.Update)
}
