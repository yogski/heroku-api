package routers

import (
	c "employee_subscription/controllers"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	router := gin.Default()

	router.GET("/user", c.UserGet)
	router.GET("/user/:id", c.UserDetail)
	router.GET("/subs", c.SubsGet)
	router.GET("/subs/:id", c.SubsDetail)
	router.POST("/user", c.UserCreate)
	router.POST("subs", c.SubsCreate)
	router.PUT("/user/:id", c.UserUpdate)
	router.DELETE("/user/:id", c.UserDelete)

	return router
}
