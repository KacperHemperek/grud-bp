package module

import (
	"net/http"
	"www.github.com/KacperHemperek/grud/types"
	"www.github.com/KacperHemperek/grud/utils"
)

func GetBase(w http.ResponseWriter, _ *http.Request) error {
	utils.WriteJson(w, http.StatusOK, types.Json{"message": "Hello world"})
	return nil
}

func PostBase(w http.ResponseWriter, _ *http.Request) error {
	utils.WriteJson(w, http.StatusCreated, types.Json{"message": "Created"})
	return nil
}

func GetTestHandler(w http.ResponseWriter, _ *http.Request) error {
	utils.WriteJson(w, http.StatusOK, types.Json{"message": "Test"})
	return nil
}
