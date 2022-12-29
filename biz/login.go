package biz

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/service"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra"
)

func Login(ctx context.Context, userId, password string) (token string, retErr error) {
	repo := new(infra.DoctorRepo)
	token, retErr = service.LoginService(ctx, userId, password, repo)
	if retErr != nil {
		return "", retErr
	}

	return
}
