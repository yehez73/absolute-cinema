package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	// For regular user
	r := echo.New()

	////// Auth
	r.POST("/register", controllers.RegisterUser) // Register for member
	r.POST("/login", controllers.Login)

	// For authenticated user
	protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware)

	////// Auth
	protected.POST("/logout", controllers.Logout)

	// For admin
	protectedAdmin := r.Group("/admin")
	protectedAdmin.Use(middleware.AdminMiddleware)

	////// User
	protectedAdmin.GET("/users", controllers.GetUsers)
	protectedAdmin.GET("/user/:id", controllers.GetUser)
	protectedAdmin.POST("/user/add", controllers.CreateUser) // Register for both role but only admin can access
	protectedAdmin.PUT("/user/update/:id", controllers.UpdateUser)
	protectedAdmin.DELETE("/user/delete/:id", controllers.DeleteUser)

	return r
}
