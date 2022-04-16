// @APIVersion 1.0.0
// @Name beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Security ApiKeyAuth
// @SecurityDefinition ApiKeyAuth apiKey Authorization header "Authorization token"
package routers

import (
	"mock/controllers"
	"mock/handlers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter("/v1/*", beego.BeforeRouter, handlers.JwtFilter)
	beego.InsertFilter("/v1/*", beego.BeforeRouter, handlers.RoleFilter)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(&controllers.AuthController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(&controllers.CategoryController{}),
		),
		beego.NSNamespace("/food",
			beego.NSInclude(&controllers.FoodController{}),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(&controllers.OrderController{}),
		),
		beego.NSNamespace("/basket",
			beego.NSInclude(&controllers.BasketController{}),
		),
	)

	beego.AddNamespace(ns)
}
