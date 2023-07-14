package routes

import (
    f "github.com/gofiber/fiber/v2"
    c "github.com/arcsolace/ak-skin-tracker/controllers"
)

func SkinsRoute(r f.Router) {
    r.Get("/", c.GetAllSkins)
	r.Get("/:id", c.GetSkinByID)
}

func UserRoute(r f.Router) {
	r.Get("/:user_code", c.GetUserSkins)
	r.Post("/:user_code", c.UpdateUserSkins)
	r.Post("/create", c.CreateUser)
	r.Post("/:user_code", c.RemoveUserSkins)
	r.Delete("/delete/:user_code", c.DeleteUser)
}

func ShareRoute(r f.Router) {
	r.Get("/:share_code", c.GetSharedSkins)
}