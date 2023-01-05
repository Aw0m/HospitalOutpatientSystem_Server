package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/bdm"

	"gorm.io/gorm"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/utils"

	"github.com/pkg/errors"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/repo"
)

func DoctorLoginService(ctx context.Context, userId, password string, repo repo.DoctorRepoInterface) (token string, retError error) {
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

// PatientServiceLogin 用户登录服务，生成Token。并且在数据库中检索该用户是否已经注册，如果没有则还会在数据库中创建该用户
func PatientServiceLogin(ctx context.Context, code, userName string, repo repo.PatientRepoInterface) (token string, retError error) {
	url := fmt.Sprintf(utils.OpenIDURL, code)
	// 通过code获取用户的唯一标识符openid
	res, err := http.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "fail to get openid")
	}

	// 解析微信服务端的response，获得openid并查询是否已经存入数据库，如果没有则在数据库中生成一个user
	body, _ := utils.ParseResponse(res)
	openidAny := body["openid"]

	if openid, ok := openidAny.(string); ok {
		// 生成token。如果数据库里没有该用户，则在该数据库生成该user
		patient, err := repo.FindNonNil(ctx, openid)
		if err == nil {
			return patient.CreateToken(), nil
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.WithMessage(err, "find patient err")
		}

		// 未注册过，需要先注册
		patient = &bdm.Patient{
			UserBase: bdm.UserBase{
				ID:       openid,
				Name:     userName,
				Birthday: time.Now(),
			},
		}
		err = repo.Save(ctx, patient)
		if err != nil {
			return "", errors.WithMessage(err, "register patient fail")
		}
		return patient.CreateToken(), nil

	} else {
		return "", errors.Errorf("openid 不为 string. openid:%+v", openid)
	}
}
