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
		"message":    "success",
		"order_list": orderList,
	})
	return
}

func GetRegisterOrder(ctx *gin.Context) {
	patientID, err := middleware.GetUserID(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	orderList, err := biz.GetRegisterOrder(ctx, patientID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message":    err.Error(),
			"order_list": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":    "success",
		"order_list": orderList,
	})
	return
}

func GetPrescriptionOrder(ctx *gin.Context) {
	req := new(cqe.GetPrescriptionOrder)
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}
	orderList, err := biz.GetPrescriptionOrder(ctx, req.RegisterOrderID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message":    err.Error(),
			"order_list": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":    "success",
		"order_list": orderList,
	})
	return
}

func ToAppoint(ctx *gin.Context) {
	patientID, err := middleware.GetUserID(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	req := new(cqe.ToAppointRequest)
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}
	res, err := biz.ToAppoint(ctx, patientID, req.DoctorID, req.AppointTime)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message":     err.Error(),
			"appoint_res": false,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":     "success",
		"appoint_res": res,
	})
}

func GetCheckOrder(ctx *gin.Context) {
	req := new(cqe.GetCheckOrderRequest)
	if ctx.ShouldBindJSON(req) != nil {
		ctx.JSON(403, gin.H{"message": "wrong param"})
		return
	}
	res, err := biz.GetCheckOrder(ctx, req.RegisterOrderID)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message":          err.Error(),
			"check_order_list": nil,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":          "success",
		"check_order_list": res,
	})
}
