package handlers

import (
	"api/models"
	"api/pkg/utils/responser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	userId   = "userId"
	userName = "userName"
	email    = "email"
)

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "no authorization header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
	}
	token := headerParts[1]
	user, err := h.service.Auth.ParseToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	c.Set(userId, user.Id)
	c.Set(userName, user.Name)
	c.Set(email, user.Email)
}

func getUserId(c *gin.Context) (int, error) {
	strId, ok := c.Get(userId)
	if !ok {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "no user id found")
		return 0, fmt.Errorf("no user id found")
	}
	id, ok := strId.(int)
	if !ok {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "user id not int")
		return 0, fmt.Errorf("user id not int")
	}
	return id, nil
}

func getUserInfo(c *gin.Context) (models.UserInfo, error) {
	strId, _ := c.Get(userId)
	id, _ := strId.(int)
	name, _ := c.Get(userName)
	email, _ := c.Get(email)
	user := models.UserInfo{
		Id:    id,
		Name:  name.(string),
		Email: email.(string),
	}
	return user, nil
}
