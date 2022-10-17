package controllers

import (
	"final-project/server/controllers/view"
	"final-project/server/helper"
	"final-project/server/request"
	"final-project/server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

// RegisterUser godoc
// @Summary Register User
// @Description Register User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body request.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 400 {object} view.Response
// @Failure 500 {object} view.ResponseError
// @Router /user/register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		ctx.JSON(http.StatusBadRequest, view.ErrorValidation(http.StatusBadRequest, "Error Validation", res))
		return
	}

	user, err := c.service.Register(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Login godoc
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body request.UserLoginRequest true "User"
// @Success 200 {object} view.ResponseLogin
// @Failure 400 {object} view.Response
// @Failure 500 {object} view.ResponseError
// @Router /user/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req request.UserLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		ctx.JSON(http.StatusBadRequest, view.ErrorValidation(http.StatusBadRequest, "Error Authentication", res))
		return
	}

	email, err := c.service.Login(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	token, err := helper.GenerateToken(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, view.ResponseLogin{
		Token: token,
	})
}

// UpdateUser godoc
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Param userid path int true "User ID"
// @Param user body request.UpdateUserRequest true "User"
// @Success 200 {object} view.ResponseUpdateUser
// @Failure 400 {object} view.Response
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /user/{userid} [put]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (c *UserController) Update(ctx *gin.Context) {
	var req request.UpdateUserRequest
	email := ctx.GetString("email")
	idUser, err := c.service.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	userIdParams := ctx.Param("userid")

	convertedUserIdParams, err := strconv.Atoi(userIdParams)

	if convertedUserIdParams != idUser {
		ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		ctx.JSON(http.StatusBadRequest, view.ErrorValidation(http.StatusBadRequest, "Error Authentication", res))
		return
	}

	user, err := c.service.Update(convertedUserIdParams, &req)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Param userid path int true "User ID"
// @Success 200 {object} view.ResponseDeleteUser
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /user/{userid} [delete]
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func (c *UserController) Delete(ctx *gin.Context) {
	email := ctx.GetString("email")

	data, err := c.service.Delete(email)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, data)
}
