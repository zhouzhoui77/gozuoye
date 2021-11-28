package routers

import "github.com/gin-gonic/gin"

func AddRoutersInit(r *gin.Engine) {
	addRouters := r.Group("AddUser")
	addRouters.GET("", func(c *gin.Context) {
		c.HTML(200, "add.html", gin.H{
			"Msg": "注册成功",
		})
	})
}
