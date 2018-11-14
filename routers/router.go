package routers

import (
	"github.com/Penun/recutil/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/recipe", &controllers.RecipeController{}, "get:GetAll")
	beego.Router("/recipe/10", &controllers.RecipeController{}, "get:GetTen")
    beego.Router("/recipe/add", &controllers.RecipeController{}, "post:AddRecipe")
    beego.Router("/recipe/delete", &controllers.RecipeController{}, "post:DeleteRecipe")
}
