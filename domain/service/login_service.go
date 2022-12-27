package service

import (
	"context"
	"github.com/pkg/errors"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
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

func RegisterService(ctx context.Context, userId, password, username string, repo repo.DoctorRepoInterface) (retError error) {
	_, err := repo.FindNonNil(ctx, userId)
	if err == nil {
		return errors.New("repeated userID")
	}

	doctor := &bdm.Doctor{
		UserBase: bdm.UserBase{
			ID:       userId,
			Name:     username,
			Password: password,
		},
	}
	return repo.Save(ctx, doctor)
}
