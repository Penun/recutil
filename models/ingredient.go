package models

import (
    "github.com/astaxie/beego/orm"
)

func init() {
    orm.RegisterModel(new(Ingredient))
}

type Ingredient struct {
    Id int64 `orm:"pk" json:"id"`
    Name string `orm:"unique" json:"name"`
}

func GetIngredients() []Ingredient{
    o := orm.NewOrm()
    var ingreds []Ingredient
    o.QueryTable("recipe").OrderBy("name").All(&ingreds)
    if len(ingreds) > 0 {
        return ingreds
    } else {
        return []Ingredient{}
    }
}

func AddIngredient(ing Ingredient) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&ing)
	if err == nil {
		return id
	} else {
		return 0
	}
}

func ReadCreateIngredient(name string) (int64, Ingredient) {
    o := orm.NewOrm()
    ingred := Ingredient{Name: name}
    if created, id, err := o.ReadOrCreate(&ingred, "name"); err == nil {
        if created {
             return id, Ingredient{}
        } else {
            return id, ingred
        }
    }
    return 0, Ingredient{}
}
