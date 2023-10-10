package controllers

import (
	"Praktikum/configs"
	"Praktikum/helpers"
	"Praktikum/middlewares"
	"Praktikum/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Get All User
func GetUsersController(c echo.Context) error {
	users := []models.User{}

	result := configs.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Failed to get all user",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Success get all user",
		Data:    users,
	})
}

// Get user By Id
func GetUserController(c echo.Context) error {
	user := models.User{}

	id, _ := strconv.Atoi(c.Param("id"))
	result := configs.DB.Find(&user, id)
	if result.Error == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusNotFound, helpers.BaseResponse{
			Status:  false,
			Message: "Id not found",
			Data:    nil,
		})

	}
	var err error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.BaseResponse{
			Status:  false,
			Message: "Failed to fetch Id",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Success get user by Id",
		Data:    user,
	})
}

// Post Create New User
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user.Password = helpers.HashPassword(user.Password)
	result := configs.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Invalid create user",
			Data:    nil,
		})

	}
	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Success create new user",
		Data:    user,
	})
}

// Delete User By Id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.User{}
	err := configs.DB.Delete(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, helpers.BaseResponse{
				Status:  false,
				Message: "Id not found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, helpers.BaseResponse{
			Status:  false,
			Message: "Failed to fetch Id",
			Data:    nil,
		})

	}

	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Success delete user by Id",
		Data:    nil,
	})
}

// Update User By Id
func UpdateUserController(c echo.Context) error {
	var updateUser models.User
	inputUser := models.User{}

	c.Bind(&inputUser)
	id, _ := strconv.Atoi(c.Param("id"))
	result := configs.DB.First(&updateUser, "id = ?", id)
	fmt.Println(updateUser.ID)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Invalid user Id",
			Data:    nil,
		})
	}

	updateUser.Email = inputUser.Email
	updateUser.Name = inputUser.Name
	updateUser.Password = helpers.HashPassword(inputUser.Password)
	configs.DB.Save(&updateUser)

	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "User updates successfully",
		Data:    updateUser.ResponseConvertUser(),
	})
}

// Login User
func LoginController(c echo.Context) error {
	request := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.User{}
	err := configs.DB.Where("email = ? ", request.Email).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Invalid username and password",
			Data:    nil,
		})
	}
	err = helpers.ComparePassword(user.Password, request.Password)

	token, err := middlewares.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Error create token",
			Data:    nil,
		})

	}
	userResponse := models.UserResponseLogin{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}
	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Success login",
		Data:    userResponse,
	})

}
