package routes

import (
	"gin-crud/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)


func UrlMaps(){
	r.GET("/api/v1/health", func (ctx * gin.Context)  {
		ctx.String(http.StatusOK, "API is Alive")
	})
	v1 := r.Group("/v1/users")
	{
		v1.GET("/users", controller.GetUsers)
		v1.GET("/user/:id", controller.GetUser)
		v1.DELETE("user/:id", controller.DeleteUser)
		v1.PATCH("/user/:id", controller.UpdateUser)
		v1.POST("/user/", controller.CreatUser)
	}
}