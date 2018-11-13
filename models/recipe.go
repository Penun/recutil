package models

import (
    "errors"
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

func GetRecipes() []Recipe {
    o := orm.NewOrm()
    var recipes []Recipe
    o.QueryTable("recipe").OrderBy("name").All(&recipes)
    if len(recipes) > 0 {
        return recipes
    } else {
        return []Recipe{}
    }
}

func GetTenRecipes() ([]Recipe, error) {
    o := orm.NewOrm()
    var recipes []Recipe
    o.QueryTable("recipe").OrderBy("-id").Limit(10).All(&recipes)
    if len(recipes) > 0 {
        return recipes, nil
    } else {
        return nil, errors.New("Empty Ten")
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

func GetRecipeIngredients_R(r_id int64) ([]RecipeIngredient, error) {
    o := orm.NewOrm()
    var recIngrs []RecipeIngredient
    o.QueryTable("recipe_ingredient").Filter("recipe_id", r_id).RelatedSel("ingredient").All(&recIngrs)
    if len(recIngrs) > 0 {
        return recIngrs, nil
    } else {
        return nil, errors.New("Missing Ingredients")
    }
}

func GetMethod_R(r_id int64) (Method, error) {
    o := orm.NewOrm()
    method := Method{}
    o.QueryTable("method").Filter("recipe_id", r_id).One(&method)
    if method != (Method{}) {
        return method, nil
    } else {
        return Method{}, errors.New("No Method")
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
