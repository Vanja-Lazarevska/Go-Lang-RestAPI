package routes

import (
	"appointment-tracking/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/staff", handleGetAllDoctors)
	server.POST("/staff", handleCreateDoctor)
	server.GET("/staff/:id", handleGetDoctorsById)
	server.DELETE("/staff/:id", handleDeleteDoctor)
	server.PUT("/staff/update/:id", handleUpdateDoctor)
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.GET("/get_appointments", handleGetAppointments)
	server.GET("/client", handleGetAllClients)

	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.POST("/create_appointment", createAppointment)
	authenticatedRoutes.PUT("/update_appointment/:id", handleUpdateAppointment)
	authenticatedRoutes.DELETE("/delete_appointment/:id", handleDeleteAppointment)

}
