package main

import (
	"fmt"
	"net/http"
	"server/router"

	"github.com/rs/cors"
)

func main() {
	//Config the router
	r := router.InitRouter()
	//condigure CORS
	handler := cors.Default().Handler(r)
	//Start server
	if err := http.ListenAndServe(":8000", handler); err != nil {
		fmt.Printf("Error al inciar el servidor. %+v", err)
		return
	}

}
