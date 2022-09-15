package database

import (
	"agmc-api/config"
	"agmc-api/models"
	"log"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() ([]models.Users, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func GetUserById(c echo.Context) ([]models.User, error) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return nil, err
	}

	var user []models.User
	log.Println(user)

	if e := config.DB.Where("id=?", userId).First(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.Users{}
	c.Bind(&user)
	user.Password = getHash([]byte(user.Password))

	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func UpdateUser(c echo.Context) (interface{}, error) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return nil, err
	}

	user := models.Users{}
	c.Bind(&user)

	if user.Password != "" {
		user.Password = getHash([]byte(user.Password))
	}

	if e := config.DB.Model(&user).Where("id=?", userId).Updates(user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func DeleteById(c echo.Context) ([]models.User, error) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return nil, err
	}

	var user []models.User

	if e := config.DB.Where("id=?", userId).Delete(user).Error; e != nil {
		return nil, e
	}

	return user, nil
}
