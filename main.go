package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/catcurd/http/cat"
	services "github.com/catcurd/services/cat"
	store "github.com/catcurd/store/cat"
)

func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false

	st := store.New()
	s := services.New(st)
	h := cat.Handler{Services: s}

	app.GET("/cat", h.Get)
	app.GET("/cat/{id}", h.GetByID)
	app.POST("/cat", h.Create)
	app.PUT("/cat/{id}", h.Update)
	app.DELETE("/cat/{id}", h.Delete)

	app.Start()
}