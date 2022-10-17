package controllers

import "github.com/gin-gonic/gin"

func WriteJsonResponse(ctx *gin.Context, status int, response map[string]interface{}) {
	ctx.JSON(status, response)
}
