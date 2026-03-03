package routers

import (
	"OperationAndMonitoring/controller/room"
	"github.com/gin-gonic/gin"
)

func RoomRouter(r *gin.RouterGroup) {
	roomRouter := r.Group("/rooms")
	{
		roomRouter.GET("/options", room.Option)
		roomRouter.GET("", room.Page)
		roomRouter.GET("/:roomId/form", room.Form)
		roomRouter.POST("", room.Add)
		roomRouter.PUT("/:roomId", room.Update)
		roomRouter.DELETE("/:roomIds", room.Delete)
	}
}
