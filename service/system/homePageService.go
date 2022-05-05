package system

import (
	"association/global"
	models "association/modules"
)

type HomePageService struct{}

func (h *HomePageService) AssociationNumber() (ass []models.Association, number int64) {
	var a []models.Association
	m := global.ASS_DB.Find(&a).Scan(&ass)
	return ass, m.RowsAffected
}

func (h *HomePageService) AssociationName(name string) (ass []models.Association, number int64) {
	var a []models.Association
	m := global.ASS_DB.Find(&a).Where("name LIKE ?", "%"+name+"%").Scan(&ass)
	return ass, m.RowsAffected
}

func (h *HomePageService) SelectAssociationById(id uint) (models.Association, error) {
	var user models.User
	global.ASS_DB.Table("users").Where("id = ?", id).Scan(&user)
	associationId := user.AssociationId
	var association models.Association
	err := global.ASS_DB.Table("associations").Where("id = ?", associationId).Scan(&association).Error
	return association, err
}

func (h *HomePageService) SelectUserByAssociationId(id string) (u []models.User, number int64) {
	var user []models.User
	m := global.ASS_DB.Find(&user).Where("association_id = ?", id).Scan(&u)
	return u, m.RowsAffected
}

func (h *HomePageService) SelectUserById(id uint) (models.User, error) {
	var user models.User
	err := global.ASS_DB.Table("users").Where("id = ?", id).Scan(&user).Error
	return user, err
}

func (h *HomePageService) UpdateAssociationIdById(id, associationId uint) error {
	err := global.ASS_DB.Model(&models.User{}).Where("id = ?", id).Update("association_id", associationId).Error
	return err
}
