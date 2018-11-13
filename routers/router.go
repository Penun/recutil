package routers

import (
	"github.com/Penun/recutil/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/recipe/add", &controllers.RecipeController{}, "post:AddRecipe")
}
