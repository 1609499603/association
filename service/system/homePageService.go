package system

import (
	"association/global"
	models "association/modules"
	"association/modules/dto"
)

type HomePageService struct{}

func (h *HomePageService) AssociationNumber() (ass []models.Association, number int64) {
	var a []models.Association
	m := global.ASS_DB.Find(&a).Order("created_at desc").Scan(&ass)
	return ass, m.RowsAffected
}

func (h *HomePageService) AssociationName(name string) (ass []models.Association, number int64) {
	var a []models.Association
	m := global.ASS_DB.Find(&a).Where("name LIKE ?", "%"+name+"%").Order("created_at desc").Scan(&ass)
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

func (h *HomePageService) SelectUserByAssociationId(id string) (a []dto.AssociationContent, number int64) {
	var user []models.User
	global.ASS_DB.Find(&user).Where("association_id = ?", id).Order("created_at desc").Scan(&user)

	var content []dto.AssociationContent

	for _, v := range user {
		if v.Role == 0 {
			var student dto.AssociationContent
			student.Role = "0"
			global.ASS_DB.Model(&models.Student{}).Where("user_id = ?", v.ID).Scan(&student)
			content = append(content, student)
		} else {
			var teacher dto.AssociationContent
			teacher.Role = "1"
			global.ASS_DB.Model(&models.Teacher{}).Where("user_id = ?", v.ID).Scan(&teacher)
			content = append(content, teacher)
		}
	}
	return content, int64(len(content))
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

func (h *HomePageService) CreateNotice(notice models.Notice) error {
	err := global.ASS_DB.Create(&notice).Error
	return err
}

func (h *HomePageService) SelectNoticeByAssociationId(associationId string) (n []models.Notice, number int64) {
	var notice []models.Notice
	m := global.ASS_DB.Find(&notice).Where("is_system = ?", associationId).Order("created_at desc").Scan(&n)
	return n, m.RowsAffected
}

func (h *HomePageService) SelectNoticeById(noticeId string) (notice models.Notice, err error) {
	err = global.ASS_DB.Model(&notice).Where("id = ?", noticeId).Scan(notice).Error
	return notice, err
}
