package biz

import (
	"context"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz/assembler"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/biz/cqe"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/service"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/infra"
)

func Register(ctx context.Context, request *cqe.RegisterRequest) error {
	doctor, err := assembler.RegisterTransferDTO(request)
	if err != nil {
		return err
	}
	repo := new(infra.DoctorRepo)
	return service.RegisterService(ctx, doctor, repo)
}
