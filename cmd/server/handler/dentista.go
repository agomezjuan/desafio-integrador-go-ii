package handler

import (
	"errors"
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

func (h *DentistaHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d domain.Dentista
		if err := c.ShouldBindJSON(&d)
		err != nil {
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

func (h *DentistaHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}
		var d domain.Dentista
		if err := c.ShouldBindJSON(&d)
		err != nil {
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