package controllers

import (
	"encoding/json"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
	"mock/helpers"
	"mock/models"
	"mock/utils"
	"net/http"
)

type CategoryController struct {
	beego.Controller
}

//// CRUD
//// @Name Create
//// @Description create category
//// @Param	body		body 	models.Category	true		"body for category content"
//// @Success 200 {int} models.Category.Id
//// @Failure 500 internal server error
//// @router / [post]
func (c *CategoryController) Create() {
	var category models.Category

	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &category)

	err := models.GetDB().Create(&category).Error
	if err != nil {
		helpers.NewHttpException(&c.Controller, utils.CreateCategoryError, nil, http.StatusBadRequest)
		return
	}

	c.Data["json"] = models.ResponseDto{
		Message: "Category created successfully",
		Data:    category,
	}
	c.ServeJSON()
}

// @Name Get
// @Description get category by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoryController) Get() {
	errMessage := "Can't get category"
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		helpers.NewHttpException(&c.Controller, errMessage, errors.New(utils.InvalidInputParam), http.StatusBadRequest)
		return
	}

	var category models.Category
	err := category.GetById(id)
	if err == gorm.ErrRecordNotFound {
		helpers.NewHttpException(&c.Controller, "Category not found", err, http.StatusNotFound)
		return
	}
	c.Data["json"] = models.ResponseDto{
		Message: "success",
		Data:    category,
	}
	c.ServeJSON()
}

// @Name GetAll
// @Description get category
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Category
// @Failure 500 Internal Server Error
// @router / [get]
func (c *CategoryController) GetAllCategory() {
	var category models.Category
	sort, offset, limit := helpers.GetSortOffsetLimit(&c.Controller)
	l, err := category.GetAllCategory(sort, offset, limit)
	if err != nil {
		helpers.NewHttpException(&c.Controller, utils.GetCategoriesError, err, http.StatusNotFound)
		return
	}
	c.Data["json"] = l

	c.ServeJSON()
}

// @router /:id [delete]
func (c *CategoryController) Delete() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		helpers.NewHttpException(&c.Controller, utils.DeleteCategoryError, errors.New(utils.InvalidInputParam), http.StatusBadRequest)
		return
	}

	var category models.Category
	err := category.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.NewHttpException(&c.Controller, "Category not found", err, http.StatusNotFound)
			return
		}
		helpers.NewHttpException(&c.Controller, utils.DeleteCategoryError, err, http.StatusBadRequest)
	}
	c.Data["json"] = models.ResponseDto{
		Message: "Category deleted successfully",
	}
	c.ServeJSON()
}

// @router /:id [put]
func (c *CategoryController) Update() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		helpers.NewHttpException(&c.Controller, utils.DeleteCategoryError, errors.New(utils.InvalidInputParam), http.StatusBadRequest)
		return
	}

	var category models.Category
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &category)
	err = category.Update(id, category)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.NewHttpException(&c.Controller, "Category not found", err, http.StatusNotFound)
			return
		}
		helpers.NewHttpException(&c.Controller, utils.UPDATE_CATEGORY_ERROR, err, http.StatusBadRequest)
	}
	c.Data["json"] = models.ResponseDto{
		Message: "Category updated successfully",
	}
	c.ServeJSON()
}
