package controllers

import (
	"github.com/astaxie/beego"
	"github.com/Penun/recutil/models"
    "encoding/json"
)

type RecipeController struct {
	beego.Controller
}

type AddRecReq struct {
	Recipe models.Recipe `json:"recipe"`
	Ingredients []models.Ingredient `json:"ingredients"`
	RecipeIngredients []models.RecipeIngredient `json:"recipe_ingredients"`
    Method models.Method `json:"method"`
}

type AddRecResp struct{
    Success bool `json:"success"`
}

func (this *RecipeController) AddRecipe() {
    var insReq AddRecReq
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &insReq)
	resp := AddRecResp{Success: false}
	if err == nil {
		rec_id := models.AddRecipe(insReq.Recipe)
		insReq.Recipe.Id = rec_id
		for i := 0; i < len(insReq.Ingredients); i++ {
            ing_id, _ := models.ReadCreateIngredient(insReq.Ingredients[i].Name)
            insReq.RecipeIngredients[i].Recipe = &insReq.Recipe
            insReq.RecipeIngredients[i].Ingredient = &models.Ingredient{Id: ing_id}
            insReq.RecipeIngredients[i].Active = true
            go models.AddRecipeIngredient(insReq.RecipeIngredients[i])
        }
        insReq.Method.Recipe = &insReq.Recipe
        insReq.Method.Active = true
        go models.AddMethod(insReq.Method)
		resp.Success = true
	} else {
        beego.Error(err.Error())
    }
	this.Data["json"] = resp
	this.ServeJSON()
}
