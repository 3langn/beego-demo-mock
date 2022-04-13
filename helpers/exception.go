package helpers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HttpException struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewHttpException(c *beego.Controller, message string, error error, code int) {
	c.Ctx.Output.SetStatus(code)
	c.Data["json"] = &HttpException{message, error.Error()}
	c.ServeJSON()
}
