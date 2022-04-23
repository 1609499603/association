package system

import (
	"association/global"
	models "association/modules"
	"association/modules/dto"
)

type UserLoginService struct{}

func (*UserLoginService) LoginUser(user dto.LoginUser) (u models.User, err error) {
	err = global.ASS_DB.
		Table("users u").
		Where("u.username = ? AND u.password = ? AND u.is_deleted = 0", user.Username, user.Password).Scan(&u).Error
	return u, err
}
