package utils

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRequestBody(c *gin.Context) ([]byte, error) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, errors.New("invalid request body")
	}
	return jsonData, nil
}
func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
func IsPhoneValid(phone string) bool {
	phoneRegex := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return phoneRegex.MatchString(strings.TrimSpace(phone))
}
