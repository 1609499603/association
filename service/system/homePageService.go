package system

import (
	"association/global"
	models "association/modules"
	"association/modules/dto"
)

type HomePageService struct{}

func (h *HomePageService) AssociationPage(pageNo, pageSize int) (page []dto.AssPage, err error) {

	var a []models.Association

	global.ASS_DB.Limit(pageSize).Offset(pageNo).Find(&a).Scan(&page)

	return page, err
}
