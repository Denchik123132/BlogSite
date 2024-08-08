package main

import (
	"BlogSite/internal/app"
)

func main() {
	app := server.NewApp()
	app.Run(":8080")
}
