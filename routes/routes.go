package routes

import (
	"net/http"

	"github.com/wilker/loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
