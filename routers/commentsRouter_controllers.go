package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["mock/controllers:AuthController"] = append(beego.GlobalControllerRouter["mock/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:AuthController"] = append(beego.GlobalControllerRouter["mock/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:BasketController"] = append(beego.GlobalControllerRouter["mock/controllers:BasketController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:BasketController"] = append(beego.GlobalControllerRouter["mock/controllers:BasketController"],
        beego.ControllerComments{
            Method: "AddFood",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:CategoryController"] = append(beego.GlobalControllerRouter["mock/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:CategoryController"] = append(beego.GlobalControllerRouter["mock/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAllCategory",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:CategoryController"] = append(beego.GlobalControllerRouter["mock/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:CategoryController"] = append(beego.GlobalControllerRouter["mock/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:CategoryController"] = append(beego.GlobalControllerRouter["mock/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:FoodController"] = append(beego.GlobalControllerRouter["mock/controllers:FoodController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:FoodController"] = append(beego.GlobalControllerRouter["mock/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:FoodController"] = append(beego.GlobalControllerRouter["mock/controllers:FoodController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:OrderController"] = append(beego.GlobalControllerRouter["mock/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:OrderController"] = append(beego.GlobalControllerRouter["mock/controllers:OrderController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:OrderController"] = append(beego.GlobalControllerRouter["mock/controllers:OrderController"],
        beego.ControllerComments{
            Method: "CancelOrder",
            Router: "/cancel/:id",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:OrderController"] = append(beego.GlobalControllerRouter["mock/controllers:OrderController"],
        beego.ControllerComments{
            Method: "UpdateOrderStatus",
            Router: "/update-status/:id",
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mock/controllers:UserController"] = append(beego.GlobalControllerRouter["mock/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
