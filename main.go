package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	port = ":8080"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})

	serv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	if err := serv.ListenAndServe(); err != nil {
		logrus.Info("stop serv.ListenAndServe: %v", err)
	}

	logrus.Info("end...")
}
