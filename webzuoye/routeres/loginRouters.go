package routers

import "github.com/gin-gonic/gin"

func LoginRoutersInit(r *gin.Engine) {
	loginRouters := r.Group("LoginUser")
	loginRouters.GET("", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"Msg": "登录成功",
		})
	})
}
