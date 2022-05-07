package system

import (
	"association/global"
	models "association/modules"
	"association/modules/dto"
)

type PersonalService struct {
}

func (p *PersonalService) SelectPersonalUser(id string) (u models.User, err error) {
	err = global.ASS_DB.Model(&u).Where("id = ?", id).Scan(&u).Error
	return u, err
}

func (p *PersonalService) SelectPersonalTeacher(id string) (t models.Teacher, err error) {
	err = global.ASS_DB.Model(&t).Where("user_id = ?", id).Scan(&t).Error
	return t, err
}

func (p *PersonalService) SelectPersonalStudent(id string) (s models.Student, err error) {
	err = global.ASS_DB.Model(&s).Where("user_id = ?", id).Scan(&s).Error
	return s, err
}

func (p *PersonalService) UpdateTeacher(teacher dto.UpdateTeacher, userId string) (err error) {
	err = global.ASS_DB.Model(&models.Teacher{}).Where("user_id = ?", userId).
		Updates(models.Teacher{
			TeacherNumber: teacher.TeacherNumber,
			CollegeId:     teacher.CollegeId,
			Name:          teacher.Name,
			Gender:        teacher.Gender,
			Phone:         teacher.Phone,
			Email:         teacher.Email,
		}).
		Error
	return err
}

func (p *PersonalService) UpdateStudent(student dto.UpdateStudent, userId string) (err error) {
	err = global.ASS_DB.Model(&models.Student{}).Where("user_id = ?", userId).
		Updates(models.Student{
			CollegeId:     student.CollegeId,
			StudentNumber: student.StudentNumber,
			Name:          student.Name,
			Gender:        student.Gender,
			Phone:         student.Phone,
			Email:         student.Email,
			Major:         student.Major,
			Class:         student.Class,
		}).
		Error
	return err
}
