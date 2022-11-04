package main

import (
	"encoding/gob"
	"fmt"
	"github.com/a4anthony/go_bookings/internal/config"
	"github.com/a4anthony/go_bookings/internal/handlers"
	"github.com/a4anthony/go_bookings/internal/models"
	"github.com/a4anthony/go_bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is the main function
func main() {
	gob.Register(models.Reservation{})
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // true for production
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache: ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println(fmt.Sprintf("Server is running on port %s", port))

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
