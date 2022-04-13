package handlers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"mock/utils"
	"net/http"
	"strings"
)

func RoleFilter(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	if strings.HasPrefix(ctx.Input.URL(), "/v1/auth") || strings.HasPrefix(ctx.Input.URL(), "/v1/user") {
		return
	}
	role := ctx.Input.GetData("role").(string)
	fmt.Println("role:", role)
	for _, url := range utils.ProtectUrl {
		if strings.HasPrefix(ctx.Input.URL(), url.Url) && ctx.Request.Method == url.Method && role != url.Role {
			ctx.Output.SetStatus(http.StatusForbidden)
			ctx.Output.Body([]byte("{\"code\":403,\"message\":\"Forbidden\"}"))
		}
	}
}
