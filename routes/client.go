package routes

import (
	"appointment-tracking/models"
	"appointment-tracking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var client models.Client

	err := context.ShouldBindJSON(&client)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse request data."})
		return
	}

	err = client.CreateNewClient()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Colud not create client"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "New client is created successfully"})

}

func login(context *gin.Context) {
	var client models.Client

	err := context.ShouldBindJSON(&client)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse request data."})
		return
	}

	err = client.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(client.Email, client.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successfull", "token": token})
}

func handleGetAllClients(context *gin.Context) {
	clients, err := models.GetAllClients()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get clients"})
		return
	}
	context.JSON(http.StatusOK, clients)

}
