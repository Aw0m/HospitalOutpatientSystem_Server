package bdm

import (
	"time"

	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/domain/entity"
	"gitee.com/CQU-2022CurriculumProject/HospitalOutpatientSystem_Server/utils"
	"github.com/golang-jwt/jwt"
)

type UserBase struct {
	ID         string               `json:"id"`       // 用户ID
	Name       string               `json:"name"`     // 用户姓名
	Birthday   time.Time            `json:"birthday"` // 出生日期
	Password   string               `json:"password"`
	GenderEnum `json:"gender_enum"` // 用户性别
}

type GenderEnum struct {
	Gender string
}

func (u *UserBase) GetAge() int32 {
	interval := time.Now().Sub(u.Birthday)
	return int32(interval.Hours() / 24 / 365)
}

func (u *UserBase) CreateToken() string {
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.MyCustomClaims{
		UserId:   u.ID,
		UserName: u.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * utils.Duration).Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "test",
		},
	})
	ss, _ := myToken.SignedString(utils.MySigningKey)
	return ss
}
