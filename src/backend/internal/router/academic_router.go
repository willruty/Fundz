package router

import (
	"github.com/gin-gonic/gin"

	controller "fundz/internal/controller/academic"
)

func SetupAcademicRouter(rg *gin.RouterGroup) {

	academic := rg.Group("/academic")
	{

		// === Student ===
		// student := academic.Group("/student")
		// {
		// 	// === Student CRUD ===
		// 	student.GET("/:id", controller.StudentById)
		// 	student.POST("/", controller.CreateStudent)
		// 	student.PUT("/", controller.UpdateStudentById)
		// 	student.DELETE("/:id", controller.DeleteStudentById)
		// }

		// === Semester ===
		// semester := academic.Group("/semester")
		// {
		// 	// === Semester CRUD ===
		// 	semester.GET("/:id", controller.SemesterById)
		// 	semester.POST("/", controller.CreateSemester)
		// 	semester.PUT("/", controller.UpdateSemesterById)
		// 	semester.DELETE("/:id", controller.DeleteSemesterById)
		// }

		// === Professor ===
		// professor := academic.Group("/professor")
		// {
		// 	// === Professor CRUD ===
		// 	professor.GET("/", controller.GetAllProfessors)
		// 	professor.GET("/:id", controller.GetProfessorById)
		// 	professor.POST("/", controller.CreateProfessor)
		// 	professor.PUT("/", controller.UpdateProfessorById)
		// 	professor.DELETE("/:id", controller.DeleteProfessorById)
		// }

		// === Assignment ===
		assignment := academic.Group("/assignment")
		{
			// === Student CRUD ===
			assignment.GET("/", controller.GetAllAssignments)
			assignment.GET("/:id", controller.GetAssignmentById)
			assignment.POST("/", controller.CreateAssignment)
			assignment.PUT("/", controller.UpdateAssignmentById)
			assignment.DELETE("/:id", controller.DeleteAssignmentById)
		}

		// === Subject ===
		subject := academic.Group("/subject")
		{
			// === Subject CRUD ===
			subject.GET("/", controller.GetAllSubjects)
			subject.GET("/:id", controller.GetSubjectById)
			subject.POST("/", controller.CreateSubject)
			subject.PUT("/", controller.UpdateSubjectById)
			subject.DELETE("/:id", controller.DeleteSubjectById)
		}

		// === Study_session ===
		study_session := academic.Group("/study_session")
		{
			// === Study_session CRUD ===
			study_session.GET("/", controller.GetAllStudySessions)
			study_session.GET("/:id", controller.GetStudySessionById)
			study_session.POST("/", controller.CreateStudySession)
			study_session.PUT("/", controller.UpdateStudySessionById)
			study_session.DELETE("/:id", controller.DeleteStudySessionById)
		}

	}
	
}
