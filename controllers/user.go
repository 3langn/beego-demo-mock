package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mock/helpers"
	"mock/models"
	"mock/utils"
	"net/http"
)

type UserController struct {
	beego.Controller
}

// @Title Get all users
// @Description Get all users
// @Success 200 {int} body []*models.User
// @Failure 403 body is empty
// @router / [get]
func (u *UserController) GetAll() {
	var users models.User
	all, err := users.GetAll()
	if err != nil {
		helpers.NewHttpException(&u.Controller, utils.GetUsersError, err, http.StatusInternalServerError)
		return
	} else {
		u.Data["json"] = map[string][]*models.User{"users": all}
	}
	u.ServeJSON()
}
