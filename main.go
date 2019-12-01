package main

import (
	routes "./routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/new", routes.NewTask)
	r.GET("/tasks", routes.GetTasks)
	r.DELETE("/delete", routes.DeleteTask)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
