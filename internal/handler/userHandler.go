package handler

import (
	"net/http"

	db "github.com/Dima-Melnik/go-restapi-example-with-gin/internal/database"
	"github.com/Dima-Melnik/go-restapi-example-with-gin/internal/models"
	"github.com/Dima-Melnik/go-restapi-example-with-gin/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []*models.User

	result := db.DB.Find(&users)
	if result.Error != nil {
		utils.SendErrorResp(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	utils.SendJsonResp(c, http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	var user *models.User

	id, ok := utils.GetID(c)
	if !ok {
		return
	}

	result := db.DB.First(&user, id)
	if result.Error != nil {
		utils.SendErrorResp(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	utils.SendJsonResp(c, http.StatusOK, user)
}

func AddUser(c *gin.Context) {
	var createUser models.User
	if !utils.BindJSON(c, &createUser) {
		return
	}

	result := db.DB.Save(&createUser)
	if result.Error != nil {
		utils.SendErrorResp(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	utils.SendJsonResp(c, http.StatusCreated, &createUser)
}

func UpdateUser(c *gin.Context) {
	var updateUser models.User
	if !utils.BindJSON(c, &updateUser) {
		return
	}

	id, ok := utils.GetID(c)
	if !ok {
		return
	}

	result := db.DB.Model(&models.User{}).Where("id = ?", id).Updates(&updateUser)
	if result.Error != nil {
		utils.SendErrorResp(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	utils.SendJsonResp(c, http.StatusOK, &updateUser)
}

func DeleteUser(c *gin.Context) {
	id, ok := utils.GetID(c)
	if !ok {
		return
	}

	result := db.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		utils.SendErrorResp(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	utils.SendJsonResp(c, http.StatusOK, "Successfully deleted")
}
