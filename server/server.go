package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/zkpf/handler"
)

func Serve() *http.Server {
	router := gin.Default()
	router.GET("/", handler.UI)
	router.POST("/proof", handler.Proof)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	return srv
}

var Server = Serve()
