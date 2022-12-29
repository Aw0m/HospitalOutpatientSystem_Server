package service

import (
	"context"

	"github.com/pkg/errors"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/repo"
)

func LoginService(ctx context.Context, userId, password string, repo repo.DoctorRepoInterface) (token string, retError error) {
	doctor, err := repo.FindNonNil(ctx, userId)
	if err != nil {
		return "", errors.Wrap(err, "invalid userID")
	}

	if password != doctor.Password {
		return "", errors.New("invalid password")
	}

	token = doctor.CreateToken()

	return token, nil
}
