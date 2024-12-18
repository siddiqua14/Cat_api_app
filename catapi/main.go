package main

import (
	"catapi/controllers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Router("/", &controllers.CatController{}, "get:GetCatImage")
	web.Router("/api/catimage", &controllers.CatController{}, "get:GetCatImageAPI")
	// Serve static files (CSS and JS)
	web.SetStaticPath("/static", "static")

	web.Run()
}
