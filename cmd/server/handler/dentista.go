package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/dentista"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/web"
	"github.com/gin-gonic/gin"
)

type DentistaHandler struct {
	s dentista.Service
}

func NewDentistaHandler(s dentista.Service) *DentistaHandler {
	return &DentistaHandler{s: s}
}

// GetByID godoc
// @Summary List dentists
// @Tags Dentista
// @Description get dentista
// @Produce  json
// @Success 200 {object} web.response
// @Router /dentistas [get]
func (h *DentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista no encontrado"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

// CreateDentista godoc
// @Summary Crear dentista
// @Tags Dentista
// @Description create dentista
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /dentistas [post]
func (h *DentistaHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d domain.Dentista
		if err := c.ShouldBindJSON(&d); err != nil {
			web.Failure(c, 400, errors.New("datos inválidos"))
			return
		}
		dentista, err := h.s.Create(d)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, dentista)
	}
}

// UpdateDentista godoc
// @Summary Actualizar dentista
// @Tags Dentista
// @Description update dentista
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /dentistas/{id} [put]
func (h *DentistaHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}
		var d domain.Dentista
		if err := c.ShouldBindJSON(&d); err != nil {
			web.Failure(c, 400, errors.New("datos inválidos"))
			return
		}
		dentista, err := h.s.Update(id, d)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, dentista)
	}
}

// UpdateByField godoc
// @Summary Actualizar dentista por alguno de sus campos
// @Tags Dentista
// @Description update dentista
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /dentistas/{id} [patch]
func (h *DentistaHandler) UpdateByField() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}
		exists, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		if exists.Id == 0 {
			web.NotFound(c)
			return
		}
		var d domain.Dentista
		if err := c.ShouldBindJSON(&d); err != nil {
			web.Failure(c, 400, errors.New(err.Error()))
			return
		}

		d.Id = exists.Id

		if d.Apellido == "" {
			d.Apellido = exists.Apellido
		}

		if d.Nombre == "" {
			d.Nombre = exists.Nombre
		}

		if d.Matricula == "" {
			d.Matricula = exists.Matricula
		}

		dentista, err := h.s.Update(id, d)
		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, 200, dentista)

	}
}

// DeleteDentista godoc
// @Summary Eliminar dentista
// @Tags Dentista
// @Description delete dentista
// @Param token header string true "token"
// @Success 204 {object} web.response
// @Router /dentistas/{id} [delete]
func (h *DentistaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 204, nil)
	}
}
