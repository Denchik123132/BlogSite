package server

import (
	"BlogSite/internal/handlers"
	"BlogSite/internal/storage/database"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	app := &App{
		Router: gin.Default(),
	}
	app.setupRouter()
	app.setupDatabase()
	return app
}

func (app *App) setupRouter() {
	app.Router.LoadHTMLGlob("web/templates/*")

	app.Router.GET("/", handlers.HomeHandler)
	app.Router.GET("/add", handlers.AddPostPageHandler)
	app.Router.GET("/posts/:id", handlers.PostHandler)
	//app.Router.GET("/avatar", handlers.RandomAvatar)

	app.Router.POST("/addPost", handlers.AddPostHandler)

	app.Router.Static("/assets", "./internal/assets")
}

func (app *App) setupDatabase() {
	database.CreateTable(database.GetDB())
}

func (app *App) Run(addr string) {
	if err := app.Router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
