package system

import (
	"association/global"
	models "association/modules"
	"association/modules/dto"
	"association/utils"
)

type UserLoginService struct{}

func (*UserLoginService) LoginUser(user dto.LoginUser) (u models.User, err error) {
	user.Password = utils.MD5V([]byte(user.Password))
	err = global.ASS_DB.
		Table("users u").
		Where("u.username = ? AND u.password = ?", user.Username, user.Password).Scan(&u).Error
	return u, err
}
