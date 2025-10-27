package routes

import (
	"github.com/Dima-Melnik/go-restapi-example-with-gin/internal/handler"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/users", handler.GetAllUsers)
		user.GET("/user/:id", handler.GetUserByID)
		user.POST("/add", handler.AddUser)
		user.PUT("/update/:id", handler.UpdateUser)
		user.DELETE("/delete/:id", handler.DeleteUser)
	}

	return r
}
