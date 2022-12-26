package bdm

import "time"

type UserBase struct {
	ID       string    `json:"id"`       // 用户ID
	Name     string    `json:"name"`     // 用户姓名
	Birthday time.Time `json:"birthday"` // 出生日期

	GenderEnum `json:"gender_enum"` // 用户性别
}

type GenderEnum struct {
	Gender string
}

func (u UserBase) GetAge() int32 {
	interval := time.Now().Sub(u.Birthday)
	return int32(interval.Hours() / 24 / 365)
}
