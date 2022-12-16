package controllers

import "github.com/gin-gonic/gin"

func HasPassword(password string) string {

	return password
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {

}
