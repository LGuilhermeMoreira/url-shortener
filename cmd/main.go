package main

import (
	"fmt"
	"net/http"

	cnfg "github.com/LGuilhermeMoreira/url-shortener/config"
	"github.com/LGuilhermeMoreira/url-shortener/internal/infra/database"
	"github.com/LGuilhermeMoreira/url-shortener/router"
)

func main() {
	c, err := cnfg.NewConfig()
	if err != nil {
		panic(err)
	}
	conn, err := database.NewConnection(c)
	if err != nil {
		panic(err)
	}
	urlDb := database.UrlDb{Db: conn}
	mux := router.CreateRouter(urlDb)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", c.Port),
		Handler: mux,
	}

	fmt.Printf("server on http://localhost:%v/ 🔥\n", c.Port)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}