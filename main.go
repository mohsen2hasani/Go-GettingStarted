package main

import (
	"github.com/mohsen2hasani/Go-GettingStarted/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
