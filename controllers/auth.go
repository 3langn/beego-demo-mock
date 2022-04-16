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

type AuthController struct {
	beego.Controller
}

// @Name Register
// @Description Register users
// @Param	body body models.LoginDto	true	"body for user content"
// @Success 200 {int} body models.User
// @Failure 403 body is empty
// @router /register [post]
func (a *AuthController) Register() {
	var user models.User
	var login models.LoginDto
	json.Unmarshal(a.Ctx.Input.RequestBody, &login)
	err := user.Create(login.Username, login.Password)
	if err != nil {
		helpers.NewHttpException(&a.Controller, err.Error(), err, http.StatusBadRequest)
		panic(err)
		return
	} else {
		a.Data["json"] = map[string]string{"message": "success"}
	}
	a.ServeJSON()
}

// @Name Login
// @Description Logs user into the system
// @Param	body body models.LoginDto	true	"body for user "
// @Success 200 {int} body models.User
// @Failure 400 Login failed
// @router /login [post]
func (a *AuthController) Login() {
	var login models.LoginDto
	var user models.User
	json.Unmarshal(a.Ctx.Input.RequestBody, &login)

	err := user.FindByUsername(login.Username)

	if err != nil {
		if gorm.ErrRecordNotFound == err {
			helpers.NewHttpException(&a.Controller, utils.UserNotFoundError, err, http.StatusNotFound)
			return
		}
		helpers.NewHttpException(&a.Controller, utils.EmailNotFoundError, err, http.StatusInternalServerError)
		return
	}
	err = user.CheckPassword(login.Password)
	if err != nil {
		helpers.NewHttpException(&a.Controller, utils.InvalidPasswordError, err, http.StatusUnauthorized)
		return
	}
	var t models.Token

	token, err := t.GenerateAuthToken(user.ID)

	if err != nil {
		helpers.NewHttpException(&a.Controller, utils.InvalidPasswordError, err, http.StatusUnauthorized)
		return
	}
	a.Data["json"] = &models.LoginResponseDto{Message: utils.LoginSuccess, User: &user, Token: token}
	a.ServeJSON()
}
