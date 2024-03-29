package services

import (
	"errors"

	"github.com/br93/syllabus/syllabus-settings-go/pkg/models"
	"gorm.io/gorm/clause"
)

func CreateCourseProgram(req *models.CourseProgram) error {
	create := models.DB.Create(req)

	if create.Error != nil {
		return create.Error
	}

	return nil
}

func GetCourseProgramById(courseprogramId string) (*models.CourseProgram, error) {
	var courseprogram models.CourseProgram

	models.DB.Preload(clause.Associations).First(&courseprogram, "course_program_id", courseprogramId)

	if courseprogram.ID == 0 {
		return &courseprogram, errors.New("courseprogram not found")
	}

	return &courseprogram, nil
}

func GetCoursePrograms(preload ...string) (*[]models.CourseProgram, error) {

	var courseprograms []models.CourseProgram

	result := models.DBConfig(models.DB.Preload(clause.Associations), preload).Find(&courseprograms)

	if result.Error != nil {
		return &courseprograms, result.Error
	}

	return &courseprograms, nil
}

func GetCourseProgramsByCourseCodeIn(codes []string, preload ...string) (*[]models.CourseProgram, error) {

	var courseprograms []models.CourseProgram

	result := models.DBConfig(models.DB.Preload(clause.Associations), preload).Joins("JOIN tb_courses ON tb_courses.id = tb_course_programs.course_id").Where("tb_courses.course_code IN ?", codes).Find(&courseprograms)

	if result.Error != nil {
		return &courseprograms, result.Error
	}

	return &courseprograms, nil
}

func GetCourseProgramsByProgramAndCourseType(programCode string, courseType string, preload ...string) (*[]models.CourseProgram, error) {
	var courseprograms []models.CourseProgram

	result := models.DBConfig(models.DB.Preload(clause.Associations), preload).Joins("JOIN tb_programs ON tb_programs.id = tb_course_programs.program_id").Joins("JOIN tb_course_types ON tb_course_types.id = tb_course_programs.course_type_id").Where("tb_programs.program_code=? AND tb_course_types.type_name=?", programCode, courseType).Find(&courseprograms)

	if result.Error != nil {
		return &courseprograms, result.Error
	}

	return &courseprograms, nil
}

func GetCourseProgramsByProgramAndNotCourseType(programCode string, courseType string, preload ...string) (*[]models.CourseProgram, error) {
	var courseprograms []models.CourseProgram

	result := models.DBConfig(models.DB.Preload(clause.Associations), preload).Joins("JOIN tb_programs on tb_programs.id = tb_course_programs.program_id").Joins("JOIN tb_course_types ON tb_course_types.id = tb_course_programs.course_type_id").Where("tb_programs.program_code=? AND NOT tb_course_types.type_name=?", programCode, courseType).Find(&courseprograms)

	if result.Error != nil {
		return &courseprograms, result.Error
	}

	return &courseprograms, nil
}

func UpdateCourseProgram(courseprogramId string, req *models.CourseProgram) (*models.CourseProgram, error) {
	response, err := GetCourseProgramById(courseprogramId)

	if err != nil {
		return req, err
	}

	response.Course = req.Course
	response.Program = req.Program
	response.CourseType = req.CourseType
	response.Term = req.Term

	update := models.DB.Save(response)

	if update.Error != nil {
		return req, update.Error
	}

	return response, nil

}

func DeleteCourseProgram(courseprogramId string) error {
	get, err := GetCourseProgramById(courseprogramId)

	if err != nil {
		return err
	}

	delete := models.DB.Delete(get)

	if delete.Error != nil {
		return delete.Error
	}

	return nil
}
