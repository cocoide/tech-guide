package main

import (
	"context"
	"os"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/database"
	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/handler"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/service"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	conf.NewEnv()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("BACKEDN_URL"), os.Getenv("FRONTNED_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	db := database.NewDatabase()
	ctx := context.Background()
	cr := repo.NewCollectionRepo(db)
	ar := repo.NewArticleRepo(db)
	ur := repo.NewAccountRepo(db)
	tr := repo.NewTopicRepo(db)

	og := gateway.NewOGPGateway()
	ag := gateway.NewOpenAIGateway(ctx)
	uu := usecase.NewAccountUseCase(ur)
	ts := service.NewTopicAnalysisService(ag, tr)
	h := handler.NewHandler(ur, ar, cr, og, uu, ts)

	private := e.Group("/account", h.AuthMiddleware)
	private.GET("/private/profile/:id", h.GetAccountProfile)
	e.GET("account/profile/:id", h.GetAccountProfile)
	e.GET("account/collection/:id", h.GetCollections)

	private.POST("/bookmark", h.DoBookmark)
	private.POST("/collection", h.CreateCollection)
	private.GET("/collection", h.GetCollectionForBookmark)
	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)
	e.GET("/email", h.CheckEmail)

	e.GET("/ogp", h.GetOGP)
	e.GET("/article", h.GetArticles)
	e.GET("/article/related/:id", h.GetRelatedArticles)
	e.GET("/token", h.GenerateToken)
	e.POST("/article", h.CreateArticle)
	e.Logger.Fatal(e.Start(":8080"))
}
