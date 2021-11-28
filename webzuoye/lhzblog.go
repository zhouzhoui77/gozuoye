/* 使用说明
输入用户名和密码，
如果用户名不存在则注册成功
如果用户名存在则判断密码是否正确，正确则登录成功，反之输出密码错误
*/
package main

import (
	"net/http"
	"webzuoye/routers"

	"github.com/gin-gonic/gin"
)

//储存用户信息的结构体
type userInfo struct {
	Username string
	Password string
}

//用于存储用户的切片
var Slice []userInfo

//判断是否存在用户
func IsExist(user string) bool {
	//如果长度为0说明尚未有用户注册
	if len(Slice) == 0 {
		return false
	} else {
		//遍历切片
		for _, v := range Slice {
			// return v.Name == user //此时只能和第一个比较，所以第一个之后全为false
			if v.Username == user {
				return true
			}
		}
	}
	return false
}

//判断密码是否正确
func IsRight(user string, passwd string) bool {
	for _, v := range Slice {
		if v.Username == user {
			//先确认姓名一致，密码相同返回true
			return v.Password == passwd
		}
	}
	return false
}

//添加用户
func AddStruct(name string, passwd string) {
	var user userInfo
	user.Username = name
	user.Password = passwd
	//将结构体user加入到全局切片Slice
	Slice = append(Slice, user)
}
func main() {
	r := gin.Default()
	//调用templates目录下所有模板
	r.LoadHTMLGlob("templates/*")

	//初始化分组路由
	routers.LoginRoutersInit(r)
	routers.AddRoutersInit(r)
	routers.ShouyeRoutersInit(r)

	//定义结构体临时储存用户信息
	var user *userInfo

	//用户注册
	r.POST("/Submit", func(c *gin.Context) {
		user = &userInfo{
			c.PostForm("username"),
			c.PostForm("password"),
		}

	}, func(c *gin.Context) {
		//判断用户名是否存在
		Jug := IsExist(user.Username)
		if !Jug {
			AddStruct(user.Username, user.Password)
			c.JSON(200, "注册成功")
		} else {
			Jug = IsRight(user.Username, user.Password)
			if Jug {
				c.JSON(200, "登录成功")
			} else {
				c.JSON(200, "用户名存在，密码错误，请重试")
				//
				r.GET("/submit", func(c *gin.Context) {
					c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/shouye/")
				})
			}
		}
	})
	r.Run()
}
