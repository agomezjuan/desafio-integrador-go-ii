package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/paciente"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/web"
	"github.com/gin-gonic/gin"
)

type PacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *PacienteHandler {
	return &PacienteHandler{s: s}
}

// validateEmptys valida que los campos no esten vacios
func validarVacios(paciente *domain.Paciente) (bool, error) {

	if paciente.Nombre == "" || paciente.Apellido == "" || paciente.Domicilio == "" || paciente.FechaAlta == "" {
		return false, errors.New("los campos no pueden estar vacíos")
	}
	return true, nil
}

// validarFecha valida que la fecha de alta sea valida
func validarFecha(fecha_alta string) (bool, error) {
	dates := strings.Split(fecha_alta, "-")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("la fecha de alta es inválida. Debe tener el formato aaaa-mm-dd")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("la fecha de alta es inválida. Deben ser números")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("ingrese una fecha de alta válida entre 1900-01-01 y 9999-12-31")
	}
	return true, nil
}

// GetAll godoc
// @Summary List patients
// @Tags Paciente
// @Description get patients
// @Produce  json
// @Success 200 {object} web.response
// @Router /pacientes [get]
func (h *PacienteHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		pacientes, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 404, errors.New("no se encontraron pacientes"))
			return
		}
		web.Success(c, 200, pacientes)
	}
}

// GetByID godoc
// @Summary List patient by id
// @Tags Paciente
// @Description get patient by id
// @Param id path int true "ID"
// @Produce  json
// @Success 200 {object} web.response
// @Router /pacientes/{id} [get]
func (h *PacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}
		paciente, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("paciente no encontrado"))
			return
		}
		web.Success(c, 200, paciente)
	}
}

// CreatePaciente godoc
// @Summary Crear paciente
// @Tags Paciente
// @Description create paciente
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response
func (h *PacienteHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var d domain.Paciente
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		if err := c.ShouldBindJSON(&d); err != nil {
			web.Failure(c, 400, errors.New("todos los campos son requeridos: nombre, apellido, domicilio, fecha_alta"))
			return
		}

		valid, err := validarVacios(&d)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		valid, err = validarFecha(d.FechaAlta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		// Convertir fecha de alta a formato time
		fechaAlta := strings.Split(d.FechaAlta, "-")
		dia, _ := strconv.Atoi(fechaAlta[2])
		mes, _ := strconv.Atoi(fechaAlta[1])
		anio, _ := strconv.Atoi(fechaAlta[0])
		d.FechaAlta = fmt.Sprintf("%d-%02d-%02d", anio, mes, dia)

		paciente, err := h.s.Create(d)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, paciente)
	}
}

// UpdatePaciente godoc
// @Summary Actualizar paciente
// @Tags Paciente
// @Description update paciente
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /pacientes/{id} [put]
func (h *PacienteHandler) Update() gin.HandlerFunc {
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
		var d domain.Paciente
		if err := c.ShouldBindJSON(&d); err != nil {
			web.Failure(c, 400, errors.New("datos inválidos"))
			return
		}
		paciente, err := h.s.Update(id, d)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, paciente)
	}
}

// UpdateByField godoc
// @Summary Actualizar paciente por alguno de sus campos
// @Tags Paciente
// @Description update paciente
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /pacientes/{id} [patch]
func (h *PacienteHandler) UpdateByField() gin.HandlerFunc {
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
			web.Failure(c, 404, errors.New("paciente no encontrado"))
			return
		}
		var d domain.Paciente
		if err := c.ShouldBindJSON(&d); err != nil {
			web.Failure(c, 400, errors.New(err.Error()))
			return
		}

		d.Id = exists.Id

		if d.Nombre == "" {
			d.Nombre = exists.Nombre
		}

		if d.Apellido == "" {
			d.Apellido = exists.Apellido
		}

		if d.Domicilio == "" {
			d.Domicilio = exists.Domicilio
		}

		if d.FechaAlta == "" {
			d.FechaAlta = exists.FechaAlta
		}

		paciente, err := h.s.Update(id, d)

		if err != nil {
			web.Failure(c, 500, err)
			return
		}
		web.Success(c, 200, paciente)
	}
}

// Delete godoc
// @Summary Eliminar paciente
// @Tags Paciente
// @Description delete paciente
// @Param id path int true "ID"
// @Success 204 {object} web.response
// @Router /pacientes/{id} [delete]
func (h *PacienteHandler) Delete() gin.HandlerFunc {
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
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 204, nil)
	}
}
