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

type CommentController struct {
	commentService *services.CommentService
	userService    *services.UserService
	photoService   *services.PhotoService
}

func NewCommentController(commentService *services.CommentService, userService *services.UserService, photoService *services.PhotoService) *CommentController {
	return &CommentController{commentService: commentService, userService: userService, photoService: photoService}
}

// CreateComment godoc
// @Summary Create Comment
// @Description Create Comment
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param body body request.CreateCommentRequest true "Create Comment"
// @Success 201 {object} view.ResponseCreateComment
// @Failure 400 {object} view.Response
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /comment [post]
func (c *CommentController) Create(ctx *gin.Context) {
	var req request.CreateCommentRequest

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

	checkIfPhotoIdExist, err := c.photoService.GetPhotoById(req.PhotoId)

	if !checkIfPhotoIdExist {
		ctx.JSON(http.StatusBadRequest, view.Error(http.StatusBadRequest, "Photo Id Not Found"))
		return
	}

	data, err := c.commentService.Create(idUser, &req)

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

// GetAllCommentByUserId godoc
// @Summary Get All Comment By User Id
// @Description Get All Comment By User Id
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} view.ResponseGetAllComment
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /comment [get]
func (c *CommentController) GetAll(ctx *gin.Context) {
	email := ctx.GetString("email")
	idUser, err := c.userService.GetUserIdByEmail(email)

	data, err := c.commentService.GetAll(idUser)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
		} else if err.Error() == "Comment Not Found" {
			ctx.JSON(http.StatusNotFound, view.Error(http.StatusNotFound, err.Error()))
		} else {
			ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// UpdateComment godoc
// @Summary Update Comment
// @Description Update Comment
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param commentid path int true "Comment Id"
// @Param body body request.UpdateCommentRequest true "Update Comment"
// @Success 200 {object} view.ResponseUpdateComment
// @Failure 400 {object} view.Response
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /comment/{commentid} [put]
func (c *CommentController) Update(ctx *gin.Context) {
	var req request.UpdateCommentRequest
	idComment := ctx.Param("commentid")

	commentId, err := strconv.Atoi(idComment)

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

	data, err := c.commentService.Update(idUser, commentId, &req)

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

// DeleteComment godoc
// @Summary Delete Comment
// @Description Delete Comment
// @Tags Comment
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param commentid path int true "Comment Id"
// @Success 200 {object} view.ResponseDeleteComment
// @Failure 401 {object} view.ResponseError
// @Failure 500 {object} view.ResponseError
// @Router /comment/{commentid} [delete]
func (c *CommentController) Delete(ctx *gin.Context) {
	idComment := ctx.Param("commentid")

	commentId, err := strconv.Atoi(idComment)

	email := ctx.GetString("email")

	idUser, err := c.userService.GetUserIdByEmail(email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	err = c.commentService.Delete(idUser, commentId)

	if err != nil {
		if err.Error() == "Unauthorized" {
			ctx.JSON(http.StatusUnauthorized, view.Error(http.StatusUnauthorized, err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, view.Error(http.StatusInternalServerError, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, view.ResponseDeleteComment{
		Message: "Your comment has been successfully deleted",
	})
}
