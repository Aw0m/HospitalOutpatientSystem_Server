package controller

import (
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz"
	"github.com/gin-gonic/gin"
)

func GetAllDepartment(ctx *gin.Context) {
	departmentList, err := biz.GetAllDepartment(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message":         err.Error(),
			"department_list": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":         "success",
		"department_list": departmentList,
	})
	return
}
