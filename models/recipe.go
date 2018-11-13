package models

import (
    "github.com/astaxie/beego/orm"
)

func init() {
    orm.RegisterModel(new(Recipe), new(RecipeIngredient), new(Method))
}

type Recipe struct {
    Id int64 `orm:"pk" json:"id"`
    Name string `orm:"unique" json:"name"`
}

type RecipeIngredient struct {
    Id int64 `orm:"pk" json:"id"`
    Recipe *Recipe `orm:"rel(fk)" json:"recipe"`
    Ingredient *Ingredient `orm:"rel(fk)" json:"ingredient"`
    Quantity string `json:"quantity"`
    Unit string `json:"unit"`
    Active bool `json:"active"`
    Notes string `json:"notes"`
}

type Method struct {
    Id int64 `orm:"pk" json:"id"`
    Recipe *Recipe `orm:"rel(fk)" json:"recipe"`
    Details string `json:"details"`
    Active bool `json:"active"`
    Notes string `json:"notes"`
}

func GetRecipes() []Recipe{
    o := orm.NewOrm()
    var recipes []Recipe
    o.QueryTable("recipe").OrderBy("name").All(&recipes)
    if len(recipes) > 0 {
        return recipes
    } else {
        return []Recipe{}
    }
}

func GetRecipe(r_id int64) Recipe {
    o := orm.NewOrm()
    recipe := Recipe{Id: r_id}
    err := o.Read(&recipe)
    if err == nil {
        return recipe
    } else {
        return Recipe{}
    }
}

func AddRecipe(rec Recipe) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&rec)
	if err == nil {
		return id
	}
    return 0
}

func AddRecipeIngredient(rec_ing RecipeIngredient) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&rec_ing)
	if err == nil {
		return id
	}
    return 0
}

func AddMethod(met Method) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&met)
	if err == nil {
		return id
	}
    return 0
}
