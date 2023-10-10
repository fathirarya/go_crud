package controllers

import (
	"Praktikum/configs"
	"Praktikum/helpers"
	"Praktikum/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Get All Book
func GetBooksController(c echo.Context) error {
	books := []models.Book{}

	result := configs.DB.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Failed to get all book",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Success get all book",
		Data:    books,
	})
}

// Get Book By Id
func GetBookController(c echo.Context) error {
	book := models.Book{}

	id, _ := strconv.Atoi(c.Param("id"))
	result := configs.DB.Find(&book, id)
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
		Message: "Success get book by Id",
		Data:    book,
	})
}

// Post Create New User
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	result := configs.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Invalid create book",
			Data:    nil,
		})

	}
	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Success create new book",
		Data:    book,
	})
}

// Delete User By Id
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := models.Book{}
	err := configs.DB.Delete(&book, id).Error
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
		Message: "Success delete book by Id",
		Data:    nil,
	})
}

// Update User By Id
func UpdateBookController(c echo.Context) error {
	var updateBook models.Book
	inputBook := models.Book{}

	c.Bind(&inputBook)
	id, _ := strconv.Atoi(c.Param("id"))
	result := configs.DB.First(&updateBook, "id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, helpers.BaseResponse{
			Status:  false,
			Message: "Invalid book Id",
			Data:    nil,
		})
	}

	updateBook.Title = inputBook.Title
	updateBook.Author = inputBook.Author
	updateBook.Publisher = inputBook.Publisher
	configs.DB.Save(&updateBook)

	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Status:  true,
		Message: "Book updates successfully",
		Data:    updateBook.ResponseConvertBook(),
	})
}
