package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Creamos el router
	router := gin.Default() 

	// Definimos una ruta para consultar el estado del servidor
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Ejecutamos el servidor en el puerto 8081
	router.Run(":8081")
}
