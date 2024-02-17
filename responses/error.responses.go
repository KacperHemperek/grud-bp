package responses

import (
	"net/http"
	"www.github.com/KacperHemperek/grud/types"
)

var InternalServerError = types.ApiError{
	Message: "Internal server error",
	Status:  http.StatusInternalServerError,
}

var NotFoundError = types.ApiError{
	Message: "Not found",
	Status:  http.StatusNotFound,
}
