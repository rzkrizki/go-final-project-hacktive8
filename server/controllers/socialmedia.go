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

type SocmedController struct {
	socmedService *services.SocmedService
	userService   *services.UserService
}

func NewSocmedController(socmedService *services.SocmedService, userService *services.UserService) *SocmedController {
	return &SocmedController{socmedService: socmedService, userService: userService}
}

// CreateSocialMedia godoc
// @Summary Create Social Media
// @Description Create Social Media
// @Tags Social Media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param body body request.CreateSocialMedia true "Create Social Media"
// @Success 201 {object} view.ResponseCreateSocmed
// @Failure 400 {object} view.Response
// @Failure 500 {object} view.ResponseError
// @Router /socialmedias [post]
func (c *SocmedController) Create(ctx *gin.Context) {
	var req request.CreateSocialMedia

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

	data, err := c.socmedService.Create(&req, idUser)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, data)

}

// GetSocialMedia godoc
// @Summary Get Social Media
// @Description Get Social Media
// @Tags Social Media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} view.ResponseGetSocmed
// @Failure 401 {object} view.ResponseError
// @Failure 404 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /socialmedias [get]
func (c *SocmedController) Get(ctx *gin.Context) {
	email := ctx.GetString("email")

	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := c.socmedService.Get(idUser)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
		} else if err.Error() == "Social Media Not Found" {
			ctx.JSON(http.StatusNotFound, view.Error(http.StatusNotFound, err.Error()))
		} else {
			ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// UpdateSocialMedia godoc
// @Summary Update Social Media
// @Description Update Social Media
// @Tags Social Media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param socialMediaId path int true "Social Media Id"
// @Param body body request.UpdateSocialMedia true "Update Social Media"
// @Success 200 {object} view.ResponseUpdateSocmed
// @Failure 400 {object} view.Response
// @Failure 401 {object} view.ResponseError
// @Failure 404 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /socialmedias/{socialMediaId} [put]
func (c *SocmedController) Update(ctx *gin.Context) {
	var req request.UpdateSocialMedia

	idSocmed := ctx.Param("socialMediaId")

	socmedId, err := strconv.Atoi(idSocmed)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

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

	data, err := c.socmedService.Update(&req, idUser, socmedId)

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

// DeleteSocialMedia godoc
// @Summary Delete Social Media
// @Description Delete Social Media
// @Tags Social Media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param socialMediaId path int true "Social Media Id"
// @Success 200 {object} view.ResponseDeleteSocmed
// @Failure 401 {object} view.ResponseError
// @Failure 404 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /socialmedias/{socialMediaId} [delete]
func (c *SocmedController) Delete(ctx *gin.Context) {
	idSocmed := ctx.Param("socialMediaId")

	socmedId, err := strconv.Atoi(idSocmed)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, err.Error()))
		return
	}

	email := ctx.GetString("email")
	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	data, err := c.socmedService.Delete(idUser, socmedId)

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
