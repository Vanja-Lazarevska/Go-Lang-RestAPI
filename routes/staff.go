package routes

import (
	"appointment-tracking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGetAllDoctors(context *gin.Context) {
	staff, err := models.GetAllStaff()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get staff members"})
		return
	}
	context.JSON(http.StatusOK, staff)
}

func handleCreateDoctor(context *gin.Context) {
	var staff models.Staff
	if err := context.ShouldBindJSON(&staff); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	if err := staff.CreateStaff(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create staff member"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"staffCreated": staff})
}

func handleGetDoctorsById(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id sent"})
		return
	}

	row, err := models.GetStaffById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find staff member by this id"})
		return
	}

	context.JSON(http.StatusOK, row)

}

func handleUpdateDoctor(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id sent"})
		return
	}

	_, err = models.GetStaffById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find staff member by this id"})
		return
	}
	var updatedDoctor models.Staff

	err = context.ShouldBindJSON(&updatedDoctor)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	updatedDoctor.ID = id
	err = updatedDoctor.UpdateDoctor()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update staff member by this id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Staff memeber updated succesfully"})
}

func handleDeleteDoctor(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id sent"})
		return
	}

	doctor, err := models.GetStaffById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find staff member by this id"})
		return
	}

	err = doctor.DeleteDoctor()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete staff member by this id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Staff member deleted successfully"})
}
