package main

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/controller"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra/dal/client_db"
	"github.com/gin-gonic/gin"
)

func Init() {
	client_db.InitDB()
}

func main() {
	Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
