package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

// Success escribe una respuesta exitosa
func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}

// Failure escribe una respuesta fallida
func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Status:  status,
		Code:    http.StatusText(status),
		Message: err.Error(),
	})
}

// Not Found escribe una respuesta fallida en caso de no encontrar Datos
func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, errorResponse{
		Message: "Sin registros",
		Status:  http.StatusNotFound,
		Code:    http.StatusText(http.StatusNotFound),
	})
}
