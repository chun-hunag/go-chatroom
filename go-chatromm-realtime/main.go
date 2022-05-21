package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("public/index.html")
	router.Static("public", "public")
	router.GET("/room/:roomId", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(context *gin.Context) {
		roomId := context.Param("roomId")
		serveWs(context.Writer, context.Request, roomId)
	})

	router.Run("localhost:8081")
}
