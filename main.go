package main

import (
	"fmt"
	"log"
	"net/http"
	"www.github.com/KacperHemperek/grud/module"
	"www.github.com/KacperHemperek/grud/utils"
)

func main() {

	utils.NotFoundRouter()
	module.Router()

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
