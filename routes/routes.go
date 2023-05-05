package routes

import (
	"time"
	"toko-buah/controller"
	"toko-buah/controller/c_buah"
	"toko-buah/controller/c_user"
	"toko-buah/handler"
	"toko-buah/middleware"
	"toko-buah/repository/r_buah"
	"toko-buah/repository/r_user"
	"toko-buah/service/s_buah"
	"toko-buah/service/s_user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Secure JSON prefix
	r.SecureJsonPrefix(")]}',\n")

	// Setup auth middleware
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		panic(err)
	}

	registerUserRepo := r_user.NewUserRepository()
	registerUser := c_user.NewUserController(registerUserRepo)

	getUserRepo := r_user.NewGetUserRepository()
	userGetService := s_user.NewGetUserService(getUserRepo)

	updateUserRepo := r_user.NewUpdateUserRepository()
	userUpdateService := s_user.NewUpdateUserService(updateUserRepo)

	userDeleteRepo := r_user.NewDeleteUserRepository()
	userDeleteService := s_user.NewDeleteUserService(userDeleteRepo)

	getBuahRepo := r_buah.NewGetBuahRepository()
	buahGetService := s_buah.NewGetBuahService(getBuahRepo)

	addBuahRepo := r_buah.NewAddBuahRepository()
	buahAddService := s_buah.NewAddBuahService(addBuahRepo)

	updateBuahRepo := r_buah.NewUpdateBuahRepository()
	buahUpdateService := s_buah.NewUpdateBuahService(updateBuahRepo)

	buahDeleteRepo := r_buah.NewDeleteBuahRepository()
	buahDeleteService := s_buah.NewDeleteBuahService(buahDeleteRepo)

	// Create controller instance
	userGetController := c_user.NewGetUserController(userGetService)
	userUpdateController := c_user.NewUpdateUserController(userUpdateService)
	userDeleteController := c_user.NewUserDeleteController(userDeleteService)

	buahGetController := c_buah.NewGetBuahController(buahGetService)
	buahAddController := c_buah.NewBuahAddController(buahAddService)
	buahUpdateController := c_buah.NewUpdateBuahController(buahUpdateService)
	buahDeleteController := c_buah.NewBuahDeleteController(buahDeleteService)

	// Apply to public routes
	r.GET("/", controller.Helloworld)
	r.POST("/login", handler.LoginHandler)
	r.POST("/register", registerUser.RegisterUser)

	// Apply auth middleware to routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users", userGetController.GetAllUser)
		auth.GET("/users/:id_user", userGetController.GetUserByID)
		auth.PUT("/users/:id_user", userUpdateController.UpdateUser)
		auth.DELETE("/users/:id_user", userDeleteController.DeleteUser)

		auth.GET("/buah", buahGetController.GetAllBuah)
		auth.POST("buah", buahAddController.AddBuah)
		auth.PUT("buah/:buah_id", buahUpdateController.UpdateBuah)
		auth.DELETE("buah/:buah_id", buahDeleteController.DeleteBuah)
	}
	return r
}
