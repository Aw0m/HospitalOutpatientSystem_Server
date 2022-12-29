package service

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/repo"
	"github.com/pkg/errors"
)

func RegisterService(ctx context.Context, doctor bdm.Doctor, repo repo.DoctorRepoInterface) (retError error) {
	_, err := repo.FindNonNil(ctx, doctor.ID)
	if err == nil {
		return errors.New("repeated userID")
	}

	return repo.Save(ctx, &doctor)
}
