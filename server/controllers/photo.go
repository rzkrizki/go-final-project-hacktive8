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

type PhotoController struct {
	photoService *services.PhotoService
	userService  *services.UserService
}

func NewPhotoController(photoService *services.PhotoService, userService *services.UserService) *PhotoController {
	return &PhotoController{photoService: photoService, userService: userService}
}

// Createphoto godoc
// @Summary Create Photo
// @Description Create Photo
// @Tags Photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param body body request.CreatePhotoRequest true "Create Photo"
// @Success 201 {object} view.ResponseCreatePhoto
// @Failure 400 {object} view.Response
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /photos [post]
func (c *PhotoController) Create(ctx *gin.Context) {
	var req request.CreatePhotoRequest
	email := ctx.GetString("email")

	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
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
		ctx.JSON(http.StatusBadRequest, view.ErrorValidation(http.StatusBadRequest, "Error Validation", res))
		return
	}

	data, err := c.photoService.Create(&req, idUser)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, data)
}

// GetAllPhoto godoc
// @Summary Get All Photo
// @Description Get All Photo
// @Tags Photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} view.ResponseGetAllPhoto
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /photos [get]
func (c *PhotoController) GetAll(ctx *gin.Context) {
	data, err := c.photoService.GetAll()

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		if err.Error() == "No Data" {
			ctx.JSON(http.StatusNotFound, view.Error(http.StatusNotFound, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// UpdatePhoto godoc
// @Summary Update Photo
// @Description Update Photo
// @Tags Photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoid path int true "Photo ID"
// @Param body body request.UpdatePhotoRequest true "Update Photo"
// @Success 200 {object} view.ResponseUpdatePhoto
// @Failure 400 {object} view.Response
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /photos/{photoid} [put]
func (c PhotoController) Update(ctx *gin.Context) {
	var req request.UpdatePhotoRequest
	id := ctx.Param("photoid")
	email := ctx.GetString("email")

	idPhoto, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

	userId, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
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
		ctx.JSON(http.StatusBadRequest, view.ErrorValidation(http.StatusBadRequest, "Error Validation", res))
		return
	}

	data, err := c.photoService.Update(&req, idPhoto, userId)

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

// DeletePhoto godoc
// @Summary Delete Photo
// @Description Delete Photo
// @Tags Photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoid path int true "Photo ID"
// @Success 200 {object} view.ResponseDeletePhoto
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /photos/{photoid} [delete]
func (c PhotoController) Delete(ctx *gin.Context) {
	id := ctx.Param("photoid")
	email := ctx.GetString("email")

	idPhoto, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

	userId, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

	data, err := c.photoService.Delete(idPhoto, userId)

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
