package main

import (
	"net/http"

	"github.com/wilker/loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
