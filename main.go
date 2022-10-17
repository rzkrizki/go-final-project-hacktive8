package main

import (
	"final-project/config/db"
	"final-project/server/controllers"
	"final-project/server/repositories/gorm"
	"final-project/server/router"
	"final-project/server/services"
	"os"

	"github.com/joho/godotenv"
)

// @title MyGram API
// @version 1.0
// @description Hacktiv8 Scalable Web Service with Golang Final Project
// @termsOfService http://swagger.io/terms/
// @contact.name Swagger API Team
// @contact.url http://swagger.io
// @contact.email fajaramaulana.dev@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081

func main() {
	db, err := db.ConnectMysqlGorm()

	if err != nil {
		panic(err)
	}

	userRepo := gorm.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	photoRepo := gorm.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService, userService)

	commentRepo := gorm.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService, userService, photoService)

	socmedRepo := gorm.NewSocialMediaRepository(db)
	socmedService := services.NewSocialMediaService(socmedRepo)
	socmedController := controllers.NewSocmedController(socmedService, userService)

	app := router.NewRouter(userController, photoController, commentController, socmedController)

	err = godotenv.Load()

	if err != nil {
		panic(err)
	}

	app.SetupRouter(os.Getenv("PORT"))
}
