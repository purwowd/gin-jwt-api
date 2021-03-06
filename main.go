package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/purwowd/go-jwt-api/auth"
	"github.com/purwowd/go-jwt-api/handler"
	"github.com/purwowd/go-jwt-api/middleware"
	"github.com/purwowd/go-jwt-api/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dsn, _ := os.LookupEnv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	authMiddleware := middleware.AuthMiddleware

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("images", "./images")

	api := router.Group("/api/v1")
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.Login)
	api.POST("/users/email-checkers", userHandler.CheckEmailAvailability)
	api.POST("/users/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	_ = router.Run(":1337")

}
