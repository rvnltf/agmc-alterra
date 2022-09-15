package controllers

import (
	"agmc-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func V1GetBookControllers(c echo.Context) error {
	books := models.Books

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Successfully get data.",
		"data":    books,
	})
}

func V1GetBookByIdControllers(c echo.Context) error {
	bookId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return err
	}

	var book models.Book
	var status bool
	var message string
	var httpStat int

	for _, singleBook := range models.Books {
		if int64(singleBook.ID) == bookId {
			book = singleBook
			status = true
			message = "Successfully get data."
			httpStat = http.StatusOK
		} else {
			status = true
			message = "Data not found"
			httpStat = http.StatusNotFound
		}
	}

	return c.JSON(httpStat, map[string]interface{}{
		"success": status,
		"message": message,
		"data":    book,
	})
}

func V1CreateBookController(c echo.Context) error {
	var newBook models.Book
	c.Bind(&newBook)
	log.Println(newBook)
	models.Books = append(models.Books, newBook)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Successfully insert data.",
		"data":    newBook,
	})
}

func V1UpdateBookByIdControllers(c echo.Context) error {
	bookId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	var updateBook models.Book
	var status bool
	var message string
	var httpStat int

	if err != nil {
		return err
	}

	for i, singleBook := range models.Books {
		if int64(singleBook.ID) == bookId {
			singleBook.Title = updateBook.Title
			singleBook.Description = updateBook.Description
			singleBook.Author = updateBook.Author
			models.Books = append(models.Books[:i], singleBook)
			status = true
			message = "Successfully update data."
			httpStat = http.StatusOK
		} else {
			status = false
			message = "Data not found"
			httpStat = http.StatusNotFound
		}
	}

	return c.JSON(httpStat, map[string]interface{}{
		"success": status,
		"message": message,
		"data":    updateBook,
	})
}

func V1DeleteBookControllers(c echo.Context) error {
	bookId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	var status bool
	var message string
	var httpStat int

	if err != nil {
		return err
	}

	for i, singleBook := range models.Books {
		if int64(singleBook.ID) == bookId {
			models.Books = append(models.Books[:i], models.Books[i+1:]...)
			status = true
			message = "Successfully delete data."
			httpStat = http.StatusOK
		} else {
			status = false
			message = "Data not found"
			httpStat = http.StatusNotFound
		}
	}

	return c.JSON(httpStat, map[string]interface{}{
		"success": status,
		"message": message,
	})
}
