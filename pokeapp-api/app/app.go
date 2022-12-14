package app

import (
	"fmt"
	"net/http"

	"github.com/alexparco/pokeapp-api/middleware"

	cmtController "github.com/alexparco/pokeapp-api/comment/controllers"
	cmtRepo "github.com/alexparco/pokeapp-api/comment/repository"
	cmtService "github.com/alexparco/pokeapp-api/comment/services"

	favController "github.com/alexparco/pokeapp-api/fav/controllers"
	favRepo "github.com/alexparco/pokeapp-api/fav/repository"
	favService "github.com/alexparco/pokeapp-api/fav/services"

	userController "github.com/alexparco/pokeapp-api/user/controllers"
	userRepo "github.com/alexparco/pokeapp-api/user/repository"
	userService "github.com/alexparco/pokeapp-api/user/services"

	authController "github.com/alexparco/pokeapp-api/auth/controllers"
	authService "github.com/alexparco/pokeapp-api/auth/services"

	"github.com/alexparco/pokeapp-api/config"
	"github.com/alexparco/pokeapp-api/database"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Config   *config.Config
	Postgres *database.SqlClient
	Router   *gin.Engine
}

func NewApi(config *config.Config) *Api {

	postgres := database.NewSqlClient(&config.Postgres)

	return &Api{
		Config:   config,
		Postgres: postgres,
		Router:   gin.Default(),
	}
}

func (a *Api) Run() {
	a.Handler()
	a.Router.Run(fmt.Sprintf(":%v", a.Config.Server.Port))
}

func (a *Api) Handler() {

	a.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Requested-With", "Credentials", "Origin"},
		ExposeHeaders:    []string{"Content-Length", "credentials"},
		AllowCredentials: true,
	}))

	v1 := a.Router.Group("/api/v1")
	v1.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"text": "hola mundo"})
	})
	// fav
	fRepo := favRepo.NewFavRepo(a.Postgres)
	fService := favService.NewFavService(fRepo)
	fController := favController.NewFavController(fService)

	// comment
	cRepo := cmtRepo.NewCommentRepo(a.Postgres)
	cService := cmtService.NewCommentService(cRepo)
	cController := cmtController.NewCommentController(cService)

	// user
	uRepo := userRepo.NewUserRepo(a.Postgres)
	uService := userService.NewUserService(uRepo)
	uController := userController.NewUserController(uService)

	// auth, jwt
	aService := authService.NewAuthService(uRepo)
	jwtService := authService.NewJwtService(a.Config.Server.JwtKey)
	aController := authController.NewAuthController(aService, jwtService)

	// middleware
	m := middleware.NewMiddleware(jwtService)

	{
		auth := v1.Group("/auth")
		auth.POST("/register", aController.Register())
		auth.POST("/login", aController.Login())

		comment := v1.Group("/comment")
		comment.GET("", cController.GetCommentsByPokeId())

		user := v1.Group("/user")
		user.GET("/", uController.GetUsers())

		// auth user
		auth.Use(m.AuthSessionMiddleware())
		auth.GET("/profile", uController.Profile())
		auth.PATCH("/", uController.Update())
		auth.DELETE("/", uController.Delete())

		// auth comments
		comment.Use(m.AuthSessionMiddleware())
		comment.POST("", cController.Create())
		comment.PATCH("/:id", cController.UpdateMessage())
		comment.DELETE("/:id", cController.Delete())

		// auth favs
		fav := v1.Group("/fav")
		fav.Use(m.AuthSessionMiddleware())
		fav.GET("/", fController.GetAllFavsByUserId())
		fav.POST("/", fController.PostFav())
		fav.DELETE("/:id", fController.DeleteFav())
	}

}
