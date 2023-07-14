package routes

import (
    f "github.com/gofiber/fiber/v2"
    c "github.com/arcsolace/ak-skin-tracker/controllers"
)

func SkinsRoute(r f.Router) {
    r.Get("/", c.GetAllSkins)
	r.Get("/:id", c.GetSkinByID)
}