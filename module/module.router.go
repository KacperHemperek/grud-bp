package module

import "www.github.com/KacperHemperek/grud/utils"

func Router() {
	utils.Get("/api/module", GetBase)
	utils.Post("/api/module", PostBase)
	utils.Get("/api/module/test", GetTestHandler)
}
