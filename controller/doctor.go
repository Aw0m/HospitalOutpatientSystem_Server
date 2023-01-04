package controller

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz/cqe"
	"github.com/gin-gonic/gin"
)

func GetDocFromDepartment(ctx *gin.Context) {
	req := new(cqe.GetDocFromDepartmentRequest)
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}

	doctorList, err := biz.GetDocFromDepartment(ctx, req.DepartmentID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
			"doctors": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
		"doctors": doctorList,
	})
	return
}

func GetAppointOrder(ctx *gin.Context) {
	req := new(cqe.GetAppointTimeRequest)
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}

	registerList, err := biz.GetAppointOrder(ctx, req.DoctorID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message":      err.Error(),
			"registerList": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":      "success",
		"registerList": registerList,
	})
	return
}
