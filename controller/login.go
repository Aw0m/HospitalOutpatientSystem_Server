package controller

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz/cqe"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	req := &cqe.LoginRequest{}
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}
	token, err := biz.Login(ctx, req.UserID, req.Password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
			"token":   "",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
		"token":   token,
	})
	return
}
