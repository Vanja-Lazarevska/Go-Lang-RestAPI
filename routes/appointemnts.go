package routes

import (
	"appointment-tracking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createAppointment(context *gin.Context) {
	var appointment models.Appointments

	err := context.ShouldBindJSON(&appointment)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse data"})
		return
	}

	clientId := context.GetInt64("clientId")
	appointment.Client_id = clientId
	err = appointment.CreateNewAppointment()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create appointment."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Appointment created successfully"})

}

func handleUpdateAppointment(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id sent"})
		return
	}

	appointment, err := models.GetAppointmentById(id)
	clientId := context.GetInt64("clientId")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find staff member by this id"})
		return
	}

	if appointment.Client_id != clientId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update an appointment"})
		return
	}

	var updatedAppointment models.Appointments

	err = context.ShouldBindJSON(&updatedAppointment)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	updatedAppointment.ID = id
	err = updatedAppointment.UpdateAppointment(updatedAppointment.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update appointment by this id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Appointment updated succesfully"})
}

func handleGetAppointments(context *gin.Context) {
	appointments, err := models.GetAllAppointments()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get appointments"})
		return
	}
	context.JSON(http.StatusOK, appointments)
}

func handleDeleteAppointment(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id sent"})
		return
	}

	clientId := context.GetInt64("clientId")
	appointment, err := models.GetAppointmentById(id)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find appointment by this id"})
		return
	}

	if appointment.Client_id != clientId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update an appointment"})
		return
	}

	err = appointment.DeleteAppointment()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete appointment by this id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})

}
