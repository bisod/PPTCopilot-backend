package routers

import (
	"backend/controllers"
	"encoding/json"

	"github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {

		if ctx.Request.RequestURI != "/login" && ctx.Request.RequestURI != "/register" {
			cookie := ctx.Request.Header.Get("token")

			//进入/projects和/logout必须要有token和cookie
			if ctx.Request.RequestURI == "/projects" || ctx.Request.RequestURI == "/logout" || ctx.Request.Method != "GET" {
				if cookie == "" {
					resp := map[string]interface{}{
						"code":    50000,
						"message": "缺少token",
					}
					ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
					ctx.ResponseWriter.WriteHeader(200)
					json.NewEncoder(ctx.ResponseWriter).Encode(resp)
				}
			}
		}
	})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/projects", &controllers.ProjectsController{})
	beego.Router("/projects/:project_id", &controllers.ProjectController{})
	beego.Router("/projects/:project_id/files", &controllers.FilesController{})
	beego.Router("/projects/:project_id/files/:file_name", &controllers.FileController{})
	beego.Router("/search_project?:filter_words", &controllers.SearchController{})
	beego.Router("/resetpassword", &controllers.ResetPasswordController{})
	beego.Router("/sendverifyemail", &controllers.ForgetPasswdController{}, "post:SendVerificationEmail")
	beego.Router("/checkverifyemail", &controllers.ForgetPasswdController{}, "post:CheckVerificationEmail")
	beego.Router("/checklogin", &controllers.CheckLoginController{})

	gptService := beego.NewNamespace("/gpt",
		beego.NSRouter("/set_api_key", &controllers.GptController{}, "post:SetApiKey"),
		beego.NSRouter("/is_api_key_set", &controllers.GptController{}, "get:IsApiKeySet"),
		beego.NSRouter("/get_catalog", &controllers.GptController{}, "post:GetCatalog"),
		beego.NSRouter("/update_slides", &controllers.GptController{}, "post:UpdateSides"),
		beego.NSRouter("/chat", &controllers.GptController{}, "post:Chat"),
	)
	beego.AddNamespace(gptService)

}
