package system

import (
	"association/global"
	models "association/modules"
	"association/utils"
)

type RegisterService struct{}

func (r *RegisterService) InsertUser(u models.User) (err error) {
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.ASS_DB.Create(&u).Error
	return err
}

func (r *RegisterService) IsUsername(username string) (err error, s string) {
	var userStruct models.User
	err = global.ASS_DB.Table("users u").Where("u.username = ? AND u.is_deleted=0", username).Scan(&userStruct).Error
	if err != nil {
		return err, ""
	}

	return nil, userStruct.Username
}

func (r *RegisterService) InsertTeacher(t models.Teacher) (err error) {
	err = global.ASS_DB.Create(&t).Error
	return err
}

func (r *RegisterService) InsertStudent(s models.Student) (err error) {
	err = global.ASS_DB.Create(&s).Error
	return err
}
