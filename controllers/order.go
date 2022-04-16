package controllers

import (
	"encoding/json"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"mock/helpers"
	"mock/models"
	"net/http"
)

type OrderRequest struct {
	BasketItemId string `json:"basket_item_id"`
}

type OrderController struct {
	beego.Controller
}

// @router / [get]
func (c *OrderController) Get() {
	var o models.Order
	requestBody := OrderRequest{}
	user_id := c.Ctx.Input.GetData("user_id").(string)

	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	order, err := o.GetOrder(user_id)
	if err != nil {
		helpers.NewHttpException(&c.Controller, err.Error(), err, http.StatusInternalServerError)
		return
	}
	c.Data["json"] = order
	c.ServeJSON()
}

// @router / [post]
func (c *OrderController) Post() {
	var o models.Order
	requestBody := OrderRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
	err := o.CreateOrder(requestBody.BasketItemId)
	if err != nil {
		helpers.NewHttpException(&c.Controller, err.Error(), err, http.StatusInternalServerError)
		return
	}
	c.Data["json"] = models.ResponseDto{Data: nil, Message: "Order created"}
	c.ServeJSON()
}

// @router /update-status/:id [patch]
func (c *OrderController) UpdateOrderStatus() {
	status := c.GetString("status")
	user_id := c.Ctx.Input.GetData("user_id").(string)

	itemId := c.Ctx.Input.Param(":id")
	if itemId == "" {
		err := errors.New("Order id is required")
		helpers.NewHttpException(&c.Controller, "Invalid order id", err, http.StatusBadRequest)
		return
	}

	var o models.Order
	err := o.UpdateOrderStatus(itemId, user_id, status)
	if err != nil {
		helpers.NewHttpException(&c.Controller, err.Error(), err, http.StatusInternalServerError)
		return
	}

	c.Data["json"] = models.ResponseDto{Data: nil, Message: "Order status updated"}
	c.ServeJSON()
}

// @router /cancel/:id [patch]
func (c *OrderController) CancelOrder() {
	itemId := c.Ctx.Input.Param(":id")
	if itemId == "" {
		err := errors.New("Order id is required")
		helpers.NewHttpException(&c.Controller, "Invalid order id", err, http.StatusBadRequest)
		return
	}

	var o models.Order
	err := o.CancelOrder(itemId)
	if err != nil {
		helpers.NewHttpException(&c.Controller, err.Error(), err, http.StatusInternalServerError)
		return
	}

	c.Data["json"] = models.ResponseDto{Data: nil, Message: "Order cancelled"}
	c.ServeJSON()
}
