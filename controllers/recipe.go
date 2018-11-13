package controllers

import (
	"github.com/astaxie/beego"
	"github.com/Penun/recutil/models"
    "encoding/json"
)

type RecipeController struct {
	beego.Controller
}

type GetRecResp struct {
	BaseResponse
	Recipes []models.Recipe `json:"recipes"`
	Ingredients [][]models.RecipeIngredient `json:"ingredients"`
	Methods []models.Method `json:"instructions"`
}

type AddRecReq struct {
	Recipe models.Recipe `json:"recipe"`
	Ingredients []models.Ingredient `json:"ingredients"`
	RecipeIngredients []models.RecipeIngredient `json:"recipe_ingredients"`
    Method models.Method `json:"method"`
}

func (this *RecipeController) GetAll() {

}

func (this *RecipeController) GetTen() {
	var resp GetRecResp
	resp.Success = false
	recs, err := models.GetTenRecipes()
	if err != nil {
		resp.Error = "M_GT_P_01"
	} else {
		resp.Recipes = recs
		resp.Ingredients = make([][]models.RecipeIngredient, len(recs))
		resp.Methods = make([]models.Method, len(recs))
		resp.Success = true
		for i := 0; i < len(recs); i++ {
			if resp.Ingredients[i], err = models.GetRecipeIngredients_R(recs[i].Id); err != nil {
				resp.Error = "M_GT_P_02"
				resp.Success = false
				break
			}
			if resp.Methods[i], err = models.GetMethod_R(recs[i].Id); err != nil {
				resp.Error = "M_GT_P_03"
				resp.Success = false
				break
			}
		}
	}
	if !resp.Success {
		resp.Recipes = nil
		resp.Ingredients = nil
		resp.Methods = nil
	}
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *RecipeController) AddRecipe() {
    var insReq AddRecReq
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &insReq)
	resp := BaseResponse{Success: false}
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
		resp.Error = "P_01"
    }
	this.Data["json"] = resp
	this.ServeJSON()
}
