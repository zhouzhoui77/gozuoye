package routers

import "github.com/gin-gonic/gin"

func ShouyeRoutersInit(r *gin.Engine) {
	shouyeRouters := r.Group("/shouye")
	{
		shouyeRouters.GET("/", func(c *gin.Context) {
			c.HTML(200, "shouye.html", gin.H{
				"Title": "登陆/注册",
			})
		})
	}
}
