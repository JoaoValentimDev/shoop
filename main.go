package main

import (
	"net/http"

	"github.com/JoaoValentimDev/shoop/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
