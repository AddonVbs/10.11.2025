package main

import (
	"log"
	"modul/internal/ServiceLinks"
	"modul/internal/hendler"

	"github.com/labstack/echo/v4"
)

func main() {
	repo := ServiceLinks.RepositoryLinsk{}
	service := ServiceLinks.NewLinksService(repo)
	h := hendler.NewLinksServiHendler(service)

	e := echo.New()
	e.GET("/links", h.GetLink)
	e.GET("/links/:id", h.GetLinkByID)

	log.Println("Server started at :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
