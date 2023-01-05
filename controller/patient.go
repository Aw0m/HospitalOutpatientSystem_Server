package controller

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz/cqe"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/middleware"
	"github.com/gin-gonic/gin"
)

func GetPlace(ctx *gin.Context) {
	req := new(cqe.GetPlaceRequest)
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}
	patientID, err := middleware.GetUserID(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	place, err := biz.GetPlace(ctx, req.DoctorID, patientID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
			"place":   -1,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
		"place":   place,
	})
	return
}

func GetPayment(ctx *gin.Context) {
	patientID, err := middleware.GetUserID(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	orderList, err := biz.GetPayment(ctx, patientID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message":    err.Error(),
			"order_list": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
		"place":   orderList,
	})
	return
}