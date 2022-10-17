package router

import (
	"final-project/server/controllers"
	"final-project/server/middleware"

	docs "final-project/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type router struct {
	router  *gin.Engine
	user    *controllers.UserController
	photo   *controllers.PhotoController
	comment *controllers.CommentController
	socmed  *controllers.SocmedController
}

func NewRouter(user *controllers.UserController, photo *controllers.PhotoController, comment *controllers.CommentController, socmed *controllers.SocmedController) *router {
	return &router{
		router:  gin.Default(),
		user:    user,
		photo:   photo,
		comment: comment,
		socmed:  socmed,
	}
}

func (r *router) SetupRouter(port string) {
	docs.SwaggerInfo.BasePath = ""
	user := r.router.Group("/users")
	user.POST("/register", r.user.Register)
	user.POST("/login", r.user.Login)
	user.PUT("/:userid", middleware.Authentication, r.user.Update)
	user.DELETE("", middleware.Authentication, r.user.Delete)

	photo := r.router.Group("/photos")
	photo.POST("", middleware.Authentication, r.photo.Create)
	photo.GET("", middleware.Authentication, r.photo.GetAll)
	photo.PUT(":photoid", middleware.Authentication, r.photo.Update)
	photo.DELETE(":photoid", middleware.Authentication, r.photo.Delete)

	comment := r.router.Group("/comments")
	comment.POST("", middleware.Authentication, r.comment.Create)
	comment.GET("", middleware.Authentication, r.comment.GetAll)
	comment.PUT(":commentid", middleware.Authentication, r.comment.Update)
	comment.DELETE(":commentid", middleware.Authentication, r.comment.Delete)

	socmed := r.router.Group("/socialmedias")
	socmed.POST("", middleware.Authentication, r.socmed.Create)
	socmed.GET("", middleware.Authentication, r.socmed.Get)
	socmed.PUT(":socialMediaId", middleware.Authentication, r.socmed.Update)
	socmed.DELETE(":socialMediaId", middleware.Authentication, r.socmed.Delete)

	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.router.Run(port)
}
