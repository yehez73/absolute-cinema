package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	// FOR PUBLIC
	r := echo.New()
	////// Auth
	r.POST("/register", controllers.RegisterUser) // Register for member
	r.POST("/login", controllers.Login)
	////// Movie
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movie/:id", controllers.GetSpecMovie)
	r.GET("/movie/nowshowing", controllers.GetNowShowing)
	r.GET("/movie/upcoming", controllers.GetUpcoming)
	////// Showtime
	r.GET("/showtimes", controllers.GetShowtimes)
	r.GET("/showtime/:id", controllers.GetSpecShowtime)
	// Get specfic showtime group by theaters by movie and date
	r.GET("/showtime/:movie_id/:showdate", controllers.GetSpecShowtimeByMovieDate)
	
	////// Theater
	r.GET("/theaters", controllers.GetTheaters)
	r.GET("/theater/:id", controllers.GetSpecTheater)


	// FOR AUTHENTICATED USER
	protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware)
	////// Auth
	protected.POST("/logout", controllers.Logout)
	////// Booking
	protected.GET("/bookings", controllers.GetBookings)
	protected.GET("/booking/:id", controllers.GetSpecBooking)
	protected.POST("/booking/add", controllers.CreateBooking)
	protected.DELETE("/booking/delete/:id", controllers.DeleteBooking)


	// FOR ADMIN
	protectedAdmin := r.Group("/admin")
	protectedAdmin.Use(middleware.AdminMiddleware)
	////// User
	protectedAdmin.GET("/users", controllers.GetUsers)
	protectedAdmin.GET("/user/:id", controllers.GetSpecUser)
	protectedAdmin.POST("/user/add", controllers.CreateUser) // Register for both role but only admin can access
	protectedAdmin.PUT("/user/update/:id", controllers.UpdateUser)
	protectedAdmin.DELETE("/user/delete/:id", controllers.DeleteUser)
	////// Movie
	protectedAdmin.POST("/movie/add", controllers.CreateMovie)
	protectedAdmin.PUT("/movie/update/:id", controllers.UpdateMovie)
	protectedAdmin.DELETE("/movie/delete/:id", controllers.DeleteMovie)
	////// Showtime
	protectedAdmin.POST("/showtime/add", controllers.CreateShowtime)
	protectedAdmin.PUT("/showtime/update/:id", controllers.UpdateShowtime)
	protectedAdmin.DELETE("/showtime/delete/:id", controllers.DeleteShowtime)
	////// Theater
	protectedAdmin.POST("/theater/add", controllers.CreateTheater)
	protectedAdmin.PUT("/theater/update/:id", controllers.UpdateTheater)
	protectedAdmin.DELETE("/theater/delete/:id", controllers.DeleteTheater)
	////// Booking (Not really needed, but just in case)
	protectedAdmin.PUT("/booking/update/:id", controllers.UpdateBooking)

	return r
}