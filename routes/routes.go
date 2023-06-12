package routes

import (
	"time"
	"toko-buah/controller"
	"toko-buah/controller/c_about"
	c_bestseller "toko-buah/controller/c_bestseller"
	"toko-buah/controller/c_home"
	c_newarrival "toko-buah/controller/c_new_arrival"
	"toko-buah/controller/c_produk"
	"toko-buah/controller/c_send_email"
	"toko-buah/controller/c_testimoni"
	"toko-buah/controller/c_user"
	"toko-buah/handler"
	"toko-buah/middleware"
	"toko-buah/repository/r_about"
	r_bestseller "toko-buah/repository/r_best_seller"
	"toko-buah/repository/r_home"
	r_newarrival "toko-buah/repository/r_new_arrival"
	"toko-buah/repository/r_produk"
	"toko-buah/repository/r_send_email"
	"toko-buah/repository/r_testimoni"
	"toko-buah/repository/r_user"
	"toko-buah/service/s_about"
	s_bestseller "toko-buah/service/s_best_seller"
	"toko-buah/service/s_home"
	s_newarrival "toko-buah/service/s_new_arrival"
	"toko-buah/service/s_produk"
	"toko-buah/service/s_send_email"
	"toko-buah/service/s_testimoni"
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

	//User
	registerUserRepo := r_user.NewUserRepository()
	registerUser := c_user.NewUserController(registerUserRepo)

	getUserRepo := r_user.NewGetUserRepository()
	userGetService := s_user.NewGetUserService(getUserRepo)

	updateUserRepo := r_user.NewUpdateUserRepository()
	userUpdateService := s_user.NewUpdateUserService(updateUserRepo)

	userDeleteRepo := r_user.NewDeleteUserRepository()
	userDeleteService := s_user.NewDeleteUserService(userDeleteRepo)

	//Home
	getHomeRepo := r_home.NewGetHomeRepository()
	homeGetService := s_home.NewGetHomeService(getHomeRepo)

	addHomeRepo := r_home.NewAddHomeRepository()
	homeAddService := s_home.NewAddHomeService(addHomeRepo)

	updateHomeRepo := r_home.NewUpdateHomeRepository()
	homeUpdateService := s_home.NewUpdateHomeService(updateHomeRepo)

	homeDeleteRepo := r_home.NewDeleteHomeRepository()
	homeDeleteService := s_home.NewDeleteHomeService(homeDeleteRepo)

	//About
	getAboutRepo := r_about.NewGetAboutRepository()
	aboutGetService := s_about.NewGetAboutService(getAboutRepo)

	addAboutRepo := r_about.NewAddAboutRepository()
	aboutAddService := s_about.NewAddAboutService(addAboutRepo)

	updateAboutRepo := r_about.NewUpdateAboutRepository()
	aboutUpdateService := s_about.NewUpdateAboutService(updateAboutRepo)

	aboutDeleteRepo := r_about.NewDeleteAboutRepository()
	aboutDeleteService := s_about.NewDeleteAboutService(aboutDeleteRepo)

	//Produk
	getProdukRepo := r_produk.NewGetProdukRepository()
	produkGetService := s_produk.NewGetProdukService(getProdukRepo)

	addProdukRepo := r_produk.NewAddProdukRepository()
	produkAddService := s_produk.NewAddProdukService(addProdukRepo)

	updateProdukRepo := r_produk.NewUpdateProdukRepository()
	produkUpdateService := s_produk.NewUpdateProdukService(updateProdukRepo)

	produkDeleteRepo := r_produk.NewDeleteProdukRepository()
	produkDeleteService := s_produk.NewDeleteProdukService(produkDeleteRepo)

	//New Arrival
	getNewArrivalRepo := r_newarrival.NewGetNewArrivalRepository()
	newarrivalGetService := s_newarrival.NewGetNewArrivalService(getNewArrivalRepo)

	addNewArrivalRepo := r_newarrival.NewAddNewArrivalRepository()
	newarrivalAddService := s_newarrival.NewAddNewArrivalService(addNewArrivalRepo)

	updateNewArrivalRepo := r_newarrival.NewUpdateNewArrivalRepository()
	newarrivalUpdateService := s_newarrival.NewUpdateNewArrivalService(updateNewArrivalRepo)

	newarrivalDeleteRepo := r_newarrival.NewDeleteNewArrivalRepository()
	newarrivalDeleteService := s_newarrival.NewDeleteNewArrivalService(newarrivalDeleteRepo)

	//Best Seller
	getBestsellerRepo := r_bestseller.NewGetBestsellerRepository()
	bestsellerGetService := s_bestseller.NewGetBestsellerService(getBestsellerRepo)

	addBestsellerRepo := r_bestseller.NewAddBestsellerRepository()
	bestsellerAddService := s_bestseller.NewAddBestsellerService(addBestsellerRepo)

	updateBestsellerRepo := r_bestseller.NewUpdateBestsellerRepository()
	bestsellerUpdateService := s_bestseller.NewUpdateBestsellerService(updateBestsellerRepo)

	bestsellerDeleteRepo := r_bestseller.NewDeleteBestsellerRepository()
	bestsellerDeleteService := s_bestseller.NewDeleteBestsellerService(bestsellerDeleteRepo)

	//Best Seller
	getTestimoniRepo := r_testimoni.NewGetTestimoniRepository()
	testimoniGetService := s_testimoni.NewGetTestimoniService(getTestimoniRepo)

	addTestimoniRepo := r_testimoni.NewAddTestimoniRepository()
	testimoniAddService := s_testimoni.NewAddTestimoniService(addTestimoniRepo)

	updateTestimoniRepo := r_testimoni.NewUpdateTestimoniRepository()
	testimoniUpdateService := s_testimoni.NewUpdateTestimoniService(updateTestimoniRepo)

	testimoniDeleteRepo := r_testimoni.NewDeleteTestimoniRepository()
	testimoniDeleteService := s_testimoni.NewDeleteTestimoniService(testimoniDeleteRepo)

	//Send Email
	emailRepo := r_send_email.NewGmailRepository("smtp.gmail.com", "587", "test@gmail.com", "test123")
	emailService := s_send_email.NewEmailService(emailRepo)

	// Create controller instance
	userGetController := c_user.NewGetUserController(userGetService)
	userUpdateController := c_user.NewUpdateUserController(userUpdateService)
	userDeleteController := c_user.NewUserDeleteController(userDeleteService)

	homeGetController := c_home.NewGetHomeController(homeGetService)
	homeAddController := c_home.NewHomeAddController(homeAddService)
	homeUpdateController := c_home.NewUpdateHomeController(homeUpdateService)
	homeDeleteController := c_home.NewHomeDeleteController(homeDeleteService)

	aboutGetController := c_about.NewGetAboutController(aboutGetService)
	aboutAddController := c_about.NewAboutAddController(aboutAddService)
	aboutUpdateController := c_about.NewUpdateAboutController(aboutUpdateService)
	aboutDeleteController := c_about.NewAboutDeleteController(aboutDeleteService)

	produkGetController := c_produk.NewGetProdukController(produkGetService)
	produkAddController := c_produk.NewProdukAddController(produkAddService)
	produkUpdateController := c_produk.NewUpdateProdukController(produkUpdateService)
	produkDeleteController := c_produk.NewProdukDeleteController(produkDeleteService)

	bestsellerGetController := c_bestseller.NewGetBestsellerController(bestsellerGetService)
	bestsellerAddController := c_bestseller.NewBestsellerAddController(bestsellerAddService)
	bestsellerUpdateController := c_bestseller.NewUpdateBestsellerController(bestsellerUpdateService)
	bestsellerDeleteController := c_bestseller.NewBestsellerDeleteController(bestsellerDeleteService)

	newarrivalGetController := c_newarrival.NewGetNewArrivalController(newarrivalGetService)
	newarrivalAddController := c_newarrival.NewNewArrivalAddController(newarrivalAddService)
	newarrivalUpdateController := c_newarrival.NewUpdateNewArrivalController(newarrivalUpdateService)
	newarrivalDeleteController := c_newarrival.NewNewArrivalDeleteController(newarrivalDeleteService)

	testimoniGetController := c_testimoni.NewGetTestimoniController(testimoniGetService)
	testimoniAddController := c_testimoni.NewTestimoniAddController(testimoniAddService)
	testimoniUpdateController := c_testimoni.NewUpdateTestimoniController(testimoniUpdateService)
	testimoniDeleteController := c_testimoni.NewTestimoniDeleteController(testimoniDeleteService)

	emailController := c_send_email.NewEmailController(emailService)

	// Apply to public routes
	r.GET("/", controller.Helloworld)
	r.POST("/login", handler.LoginHandler)
	r.POST("/register", registerUser.RegisterUser)
	r.GET("public/users", userGetController.GetAllUser)
	r.GET("public/produk", produkGetController.GetAllProduk)
	r.GET("public/newarrival", newarrivalGetController.GetAllNewArrival)
	r.GET("/public/home", homeGetController.GetAllHome)
	r.GET("/public/about", aboutGetController.GetAllAbout)
	r.GET("/public/bestseller", bestsellerGetController.GetAllBestseller)
	r.GET("public/testimoni", testimoniGetController.GetAllTestimoni)
	r.POST("/public/send-email", func(c *gin.Context) {
		emailController.SendEmail(c.Writer, c.Request)
	})
	// r.GET("/users/:id_user", userGetController.GetUserByID)
	// Apply auth middleware to routes
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		// auth.GET("/users", userGetController.GetAllUser)
		auth.GET("/users/:id_user", userGetController.GetUserByID)
		auth.PUT("/users/:id_user", userUpdateController.UpdateUser)
		auth.DELETE("/users/:id_user", userDeleteController.DeleteUser)

		auth.GET("/home", homeGetController.GetAllHome)
		auth.POST("home", homeAddController.AddHome)
		auth.PUT("home/:home_id", homeUpdateController.UpdateHome)
		auth.DELETE("home/:home_id", homeDeleteController.DeleteHome)

		auth.GET("/about", aboutGetController.GetAllAbout)
		auth.POST("about", aboutAddController.AddAbout)
		auth.PUT("about/:about_id", aboutUpdateController.UpdateAbout)
		auth.DELETE("about/:about_id", aboutDeleteController.DeleteAbout)

		auth.GET("/produk", produkGetController.GetAllProduk)
		auth.POST("produk", produkAddController.AddProduk)
		auth.PUT("produk/:produk_id", produkUpdateController.UpdateProduk)
		auth.DELETE("produk/:produk_id", produkDeleteController.DeleteProduk)

		auth.GET("/newarrival", newarrivalGetController.GetAllNewArrival)
		auth.POST("newarrival", newarrivalAddController.AddNewArrival)
		auth.PUT("newarrival/:new_arrival_id", newarrivalUpdateController.UpdateNewArrival)
		auth.DELETE("newarrival/:new_arrival_id", newarrivalDeleteController.DeleteNewArrival)

		auth.GET("/bestseller", bestsellerGetController.GetAllBestseller)
		auth.POST("bestseller", bestsellerAddController.AddBestseller)
		auth.PUT("bestseller/:bestseller_id", bestsellerUpdateController.UpdateBestseller)
		auth.DELETE("bestseller/:bestseller_id", bestsellerDeleteController.DeleteBestseller)

		auth.GET("/testimoni", testimoniGetController.GetAllTestimoni)
		auth.POST("testimoni", testimoniAddController.AddTestimoni)
		auth.PUT("testimoni/:testimoni_id", testimoniUpdateController.UpdateTestimoni)
		auth.DELETE("testimoni/:testimoni_id", testimoniDeleteController.DeleteTestimoni)
	}
	return r
}
