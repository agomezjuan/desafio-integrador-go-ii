package main

import (
	"database/sql"
	"log"

	"github.com/agomezjuan/desafio-integrador-go-ii/cmd/server/handler"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/dentista"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	bd, err := sql.Open("mysql", "root:root@tcp(localhost:3346)/my_db")
	if err != nil {
		log.Fatal(err)
	}

	// Creamos el servicio
	storage := store.NewSqlStore(bd)
	repo := dentista.NewRepository(storage)
	service := dentista.NewService(repo)
	dentistaHandler := handler.NewDentistaHandler(service)

	// Creamos el router
	r := gin.Default() 

	// Definimos una ruta para consultar el estado del servidor
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Definimos una ruta para consultar un dentista por su id
	dentistas := r.Group("/dentistas")
	{
		dentistas.POST("/", dentistaHandler.Create())
		dentistas.GET("/:id", dentistaHandler.GetByID())
		dentistas.PUT("/:id", dentistaHandler.Update())
	}

	// Ejecutamos el servidor en el puerto 8081
	r.Run(":8081")
}
