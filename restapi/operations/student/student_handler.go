package student

import (
	"github.com/Cinojose/StudentEnrollmentApi/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
)

type StudentOrm struct {
	Db *gorm.DB
}

// GetStudentTableName will return the name of the student table
func (s *StudentOrm) GetStudentTableName() string {

	tablename := "student"

	return tablename

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
	err := s.Db.Debug().Table(s.GetStudentTableName()).Create(&student).Error

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

	s.Db.Debug().Table(s.GetStudentTableName()).Find(&student)

	student.Class = &params.Body.Class

	s.Db.Debug().Table(s.GetStudentTableName()).Save(&student)

	return NewUpdateStudentOK()

}

// DeteleStudentHandler function takes care of removing user from db
func (s *StudentOrm) DeleteStudentHandler(params DeleteStudentParams) middleware.Responder {

	student := models.Student{
		ID: &params.Body.ID,
	}

	s.Db.Debug().Table(s.GetStudentTableName()).Find(&student)

	s.Db.Debug().Table(s.GetStudentTableName()).Delete(&student)

	return NewDeleteStudentOK()
}

//Function to retrive data based on the request.
func (s *StudentOrm) FindStudentsHandler(params FindStudentsByStatusParams) middleware.Responder {

	id := params.ID

	class := params.Class
	students := []*models.Student{}

	if class != nil {
		// Fetch students based on class
		s.Db.Debug().Table(s.GetStudentTableName()).Where("class = ?", class).Find(&students)

		return NewFindStudentsByStatusOK().WithPayload(students)
	} else if id != nil {
		s.Db.Debug().Table(s.GetStudentTableName()).Where("id = ?", id).Find(&students)

		return NewFindStudentsByStatusOK().WithPayload(students)
	}

	return NewFindStudentsByStatusBadRequest()
}
