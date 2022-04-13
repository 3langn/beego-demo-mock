package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
	"mock/helpers"
	"mock/models"
	"mock/utils"
	"net/http"
)

type FoodController struct {
	beego.Controller
}

// @router / [get]
func (c *FoodController) GetAll() {
	var food models.Food
	foods, _ := food.FindAll()

	c.Data["json"] = models.ResponseDto{
		Message: "success",
		Data:    foods,
	}
	c.ServeJSON()
}

// @router /:id [get]
func (c *FoodController) Get() {
	var food models.Food
	id := c.Ctx.Input.Param(":id")

	err := food.FindOne(id)
	if err == gorm.ErrRecordNotFound {
		helpers.NewHttpException(&c.Controller, utils.FoodNotFoundError, err, http.StatusNotFound)
		return
	}
	c.Data["json"] = models.ResponseDto{
		Message: "success",
		Data:    food,
	}
}

// @router / [post]
func (c *FoodController) Create() {
	var dto models.FoodDto
	var foodModel models.Food

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dto)
	if err != nil {
		helpers.NewHttpException(&c.Controller, utils.InvalidRequestError, err, http.StatusBadRequest)
		return
	}

	food, err := foodModel.Create(&dto)
	if err != nil {
		helpers.NewHttpException(&c.Controller, utils.InternalServerError, err, http.StatusInternalServerError)
		return
	}

	c.Data["json"] = models.ResponseDto{
		Message: "success",
		Data:    food,
	}
	c.ServeJSON()
}
