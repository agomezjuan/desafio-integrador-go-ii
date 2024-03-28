package main

import (
	"database/sql"
	"log"

	"github.com/agomezjuan/desafio-integrador-go-ii/cmd/server/handler"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/dentista"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/paciente"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/turno"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	docs "github.com/agomezjuan/desafio-integrador-go-ii/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Certified Tech Developer
// @version 1.0
// @description This API Handle Dentists, Patients and Appointments
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	bd, err := sql.Open("mysql", "root:root@tcp(localhost:3346)/my_db")
	if err != nil {
		log.Fatal(err)
	}

	// Verificamos que la conexión a la base de datos sea exitosa
	err = bd.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Creamos la tabla dentistas
	_, err = bd.Exec("CREATE TABLE IF NOT EXISTS dentistas (id INT AUTO_INCREMENT PRIMARY KEY, apellido VARCHAR(100), nombre VARCHAR(100), matricula VARCHAR(100));")
	if err != nil {
		log.Fatal(err)
	}

	// Creamos la tabla pacientes
	_, err = bd.Exec("CREATE TABLE IF NOT EXISTS pacientes (id INT AUTO_INCREMENT PRIMARY KEY, apellido VARCHAR(100), nombre VARCHAR(100), domicilio VARCHAR(100), dni VARCHAR(100), fecha_alta DATE);")
	if err != nil {
		log.Fatal(err)
	}

	// Creamos la tabla turnos
	_, err = bd.Exec("CREATE TABLE IF NOT EXISTS turnos (id INT AUTO_INCREMENT PRIMARY KEY, fecha DATE, hora TIME, descripcion VARCHAR(250), id_paciente INT, id_dentista INT, FOREIGN KEY (id_paciente) REFERENCES pacientes(id), FOREIGN KEY (id_dentista) REFERENCES dentistas(id));")
	if err != nil {
		log.Fatal(err)
	}

	// Creamos el servicio de dentistas
	dentistaStorage := store.NewDentistaStore(bd)
	repo := dentista.NewRepository(dentistaStorage)
	service := dentista.NewService(repo)
	dentistaHandler := handler.NewDentistaHandler(service)

	// Creamos el servicio de pacientes
	pacienteStorage := store.NewPacienteStore(bd)
	repoPaciente := paciente.NewRepository(pacienteStorage)
	servicePaciente := paciente.NewService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	// Creamos el servicio Turno
	storageTurno := store.NewSqlTurno(bd)
	turnoRepo := turno.NewRepository(storageTurno)
	turnoService := turno.NewService(turnoRepo)
	turnoHandler := handler.NewTurnoHandler(turnoService, service, servicePaciente)

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
		dentistas.PATCH("/:id", dentistaHandler.UpdateByField())
		dentistas.DELETE("/:id", dentistaHandler.Delete())
	}

	// Definimos una ruta para consultar un paciente por su id
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("/", pacienteHandler.GetAll())
		pacientes.POST("/", pacienteHandler.Create())
		pacientes.GET("/:id", pacienteHandler.GetByID())
		pacientes.PUT("/:id", pacienteHandler.Update())
		pacientes.DELETE("/:id", pacienteHandler.Delete())
	}

	turnoRoutes := r.Group("/turno")
	{
		turnoRoutes.POST("/", turnoHandler.Add)
		turnoRoutes.POST("/:dni/:matricula", turnoHandler.AddByDniMatricula)
		turnoRoutes.GET("/:id", turnoHandler.GetByID)
		turnoRoutes.GET("/query", turnoHandler.GetByDNI)
		turnoRoutes.GET("/", turnoHandler.GetAll)
		turnoRoutes.PUT("/:id", turnoHandler.Update)
		turnoRoutes.PATCH("/:id", turnoHandler.Patch)
		turnoRoutes.DELETE("/:id", turnoHandler.Delete)
	}

	// Documentación de la API
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Ejecutamos el servidor en el puerto 8081
	r.Run(":8081")
}
