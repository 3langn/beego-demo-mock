package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"mock/helpers"
	"mock/models"
	"net/http"
)

type BasketController struct {
	beego.Controller
}

// @router / [get]
func (c *BasketController) Get() {
	user_id := c.Ctx.Input.GetData("user_id").(string)
	b := models.Basket{}
	basketData, err := b.GetBasket(user_id)
	if err != nil {
		helpers.NewHttpException(&c.Controller, "Error getting basket", err, http.StatusInternalServerError)
	}
	c.Data["json"] = models.ResponseDto{
		Message: "success",
		Data:    basketData,
	}
	c.ServeJSON()
}

// @router / [post]
func (c *BasketController) AddFood() {
	user_id := c.Ctx.Input.GetData("user_id").(string)

	b := models.Basket{}
	var foodId = map[string]string{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &foodId)

	err := b.AddFood(user_id, foodId["food_id"])
	if err != nil {
		helpers.NewHttpException(&c.Controller, "Error getting basket", err, http.StatusInternalServerError)
	}
	c.Data["json"] = models.ResponseDto{
		Message: "success",
		Data:    nil,
	}
	c.ServeJSON()
}
