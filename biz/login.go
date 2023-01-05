package biz

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/service"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra"
)

func Login(ctx context.Context, userId, password string) (token string, retErr error) {
	repo := new(infra.DoctorRepo)
	token, retErr = service.DoctorLoginService(ctx, userId, password, repo)
	if retErr != nil {
		return "", retErr
	}

	return
}

func WXLogin(ctx context.Context, code, username string) (token string, retErr error) {
	repo := new(infra.PatientRepo)
	token, retErr = service.PatientServiceLogin(ctx, code, username, repo)
	if retErr != nil {
		return "", retErr
	}

	return
}
