package main

import (
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/database"
	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/handler"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	conf.NewEnv()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	db := database.NewDatabase()
	cr := repo.NewCollectionRepo(db)
	ar := repo.NewArticleRepo(db)
	ur := repo.NewAccountRepo(db)
	og := gateway.NewOGPGateway()
	uu := usecase.NewAccountUseCase(ur)
	h := handler.NewHandler(ur, ar, cr, og, uu)

	private := e.Group("/account", h.AuthMiddleware)
	private.GET("/private/profile/:id", h.GetAccountProfile)
	private.POST("/collection", h.CreateCollection)
	private.POST("/bookmark", h.DoBookmark)
	e.GET("account/profile/:id", h.GetAccountProfile)
	e.GET("account/collection/:id", h.GetCollections)

	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)

	e.GET("/ogp", h.GetOGP)
	e.GET("/article", h.GetArticles)
	e.GET("/token", h.GenerateToken)
	e.POST("/article", h.CreateArticle)
	e.Logger.Fatal(e.Start(":8080"))
}
