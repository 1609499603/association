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
