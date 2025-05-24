package main

import (
	"challengeGO/config"
	"challengeGO/handler"
	"challengeGO/middleware"
	"challengeGO/repository"
	"challengeGO/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	r := gin.Default()

	// Repository
	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// Service
	userService := service.NewUserService(userRepo)
	bookService := service.NewBookService(bookRepo)
	categoryService := service.NewCategoryService(categoryRepo)

	// Handler
	authHandler := handler.NewAuthHandler(userService)
	bookHandler := handler.NewBookHandler(bookService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// API Grouping
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		books := api.Group("/books")
		books.Use(middleware.JWTAuth())
		{
			books.GET("/", middleware.RoleAuthorization("user", "admin"), bookHandler.GetBooks) // User & Admin
			books.POST("/", middleware.RoleAuthorization("admin"), bookHandler.CreateBook)      // Admin only
			books.PUT("/:id", middleware.RoleAuthorization("admin"), bookHandler.UpdateBook)    // Admin only
			books.DELETE("/:id", middleware.RoleAuthorization("admin"), bookHandler.DeleteBook) // Admin only
		}

		categories := api.Group("/categories")
		categories.Use(middleware.JWTAuth())
		{
			categories.GET("/", middleware.RoleAuthorization("admin"), categoryHandler.GetCategories)        // Admin only
			categories.POST("/", middleware.RoleAuthorization("admin"), categoryHandler.CreateCategory)      // Admin only
			categories.PUT("/:id", middleware.RoleAuthorization("admin"), categoryHandler.UpdateCategory)    // Admin only
			categories.DELETE("/:id", middleware.RoleAuthorization("admin"), categoryHandler.DeleteCategory) // Admin only
		}

	}

	r.Run()
}
