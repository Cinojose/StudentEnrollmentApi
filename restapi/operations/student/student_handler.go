package student

import (
	"github.com/Cinojose/StudentEnrollmentApi/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

type StudentOrm struct {
	Db *gorm.DB
}

// AddStudent function will add the data to database
func (s *StudentOrm) AddStudentHandler(params AddStudentParams) middleware.Responder {

	student := models.Student{
		Nationality: params.Body.Nationality,
		FirstName:   params.Body.FirstName,
		LastName:    params.Body.LastName,
		Class:       params.Body.Class,
		ID:          params.Body.ID,
	}
	//Start to add the details to DB
	err := s.Db.Debug().Table("Student").Create(&student).Error

	if err != nil {
		return NewAddStudentInternalServerError()
	}

	return NewAddStudentOK()
}

//UpdateStudentHandler function takes care of updating student
func (s *StudentOrm) UpdateStudentHandler(params UpdateStudentParams) middleware.Responder {

	// id := params.Body.ID

	// class := params.Body.Class
	student := models.Student{
		ID: &params.Body.ID,
	}

	s.Db.Debug().Table("student").Find(&student)

	student.Class = &params.Body.Class

	s.Db.Debug().Table("student").Save(&student)

	return NewUpdateStudentOK()

}

// DeteleStudentHandler function takes care of removing user from db
func (s *StudentOrm) DeleteStudentHandler(params DeleteStudentParams) middleware.Responder {

	student := models.Student{
		ID: &params.Body.ID,
	}

	s.Db.Debug().Table("student").Find(&student)

	s.Db.Debug().Table("student").Delete(&student)

	return NewDeleteStudentOK()
}
