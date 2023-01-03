package controller

import (
	"net/http"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz/cqe"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	req := new(cqe.RegisterRequest)
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}
	if req.ID == "" || req.Name == "" || req.Password == "" {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}

	err := biz.Register(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	return
}
