package handler

import (
	"errors"
	"net/http"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/dentista"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/paciente"
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/turno"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/util"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/web"
	"github.com/gin-gonic/gin"
)

type TurnoHandler interface {
	Add(c *gin.Context)
	AddByDniMatricula(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
	Delete(c *gin.Context)
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	GetByDNI(c *gin.Context)
}

type turnoHandler struct {
	service         turno.Service
	serviceDentista dentista.Service
	servicePaciente paciente.Service
}

func NewTurnoHandler(service turno.Service, serviceDentista dentista.Service, servicePaciente paciente.Service) TurnoHandler {
	return &turnoHandler{service, serviceDentista, servicePaciente}
}

// Add godoc
// @Summary Agrega un nuevo turno
// @Description Agrega un nuevo turno
// @Tags Turnos
// @Accept json
// @Produce json
// @Param body body domain.TurnoAbstract true "Turno"
// @Success 201 {int} int
// @Router /turnos [post]
func (h *turnoHandler) Add(c *gin.Context) {
	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	id, err := h.service.Add(&turno)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusCreated, id)
}

// Update godoc
// @Summary Actualiza un turno
// @Description Actualiza un turno
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param body body domain.TurnoAbstract true "Turno"
// @Success 200 {object} domain.TurnoAbstract
// @Router /turnos/{id} [put]
func (h *turnoHandler) Update(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	existeTurno, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if existeTurno == nil {
		web.NotFound(c)
		return
	}

	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	turno.Id = existeTurno.Id

	if err := h.service.Update(&turno); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, turno)
}

// Patch godoc
// @Summary Actualiza un turno
// @Description Actualiza un turno
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param body body domain.TurnoAbstract true "Turno"
// @Success 200 {object} domain.TurnoAbstract
// @Router /turnos/{id} [patch]
func (h *turnoHandler) Patch(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	existe, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if existe.Id == 0 {
		web.NotFound(c)
		return
	}

	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	turno.Id = existe.Id

	if turno.Descripcion == "" {
		turno.Descripcion = existe.Descripcion
	}

	if turno.Fecha == "" {
		turno.Fecha = existe.Fecha
	}

	if turno.Hora == "" {
		turno.Hora = existe.Hora
	}

	if turno.DentistaId == 0 {
		turno.DentistaId = existe.Dentista.Id
	}

	if turno.PacienteId == 0 {
		turno.DentistaId = existe.Paciente.Id
	}

	if err := h.service.Update(&turno); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, turno)
}

// Delete godoc
// @Summary Elimina un turno
// @Description Elimina un turno
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 204 {object} web.response
// @Router /turnos/{id} [delete]
func (h *turnoHandler) Delete(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	if err := h.service.Delete(id); err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, nil)
}

// GetByID godoc
// @Summary Obtiene un turno por ID
// @Description Obtiene un turno por ID
// @Tags Turnos
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} domain.Turno
// @Router /turnos/{id} [get]
func (h *turnoHandler) GetByID(c *gin.Context) {
	id, err := util.GetIdFromParam(c)
	if err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}
	turno, err := h.service.GetByID(id)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if turno == nil {
		web.NotFound(c)
		return
	}
	web.Success(c, http.StatusOK, turno)
}

// GetAll godoc
// @Summary Obtiene todos los turnos
// @Description Obtiene todos los turnos
// @Tags Turnos
// @Accept json
// @Produce json
// @Success 200 {object} domain.Turno
// @Router /turnos [get]
func (h *turnoHandler) GetAll(c *gin.Context) {
	dentists, err := h.service.GetAll()
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusOK, dentists)
}

// GetByDNI godoc
// @Summary Obtiene un turno por DNI
// @Description Obtiene un turno por DNI
// @Tags Turnos
// @Accept json
// @Produce json
// @Param dni path string true "DNI"
// @Success 200 {object} domain.Turno
// @Router /turnos/query [get]
func (h *turnoHandler) GetByDNI(c *gin.Context) {
	dni := util.GetDNIFromParam(c)
	if dni == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("DNI faltante"))
		return
	}
	turno, err := h.service.GetByDNI(dni)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if turno.Id == 0 {
		web.NotFound(c)
		return
	}
	web.Success(c, http.StatusOK, turno)
}

// AddByDniMatricula godoc
// @Summary Agrega un nuevo turno por DNI y Matricula
// @Description Agrega un nuevo turno por DNI y Matricula
// @Tags Turnos
// @Accept json
// @Produce json
// @Param dni path string true "DNI"
// @Param matricula path string true "Matricula"
// @Param body body domain.TurnoAbstract true "Turno"
// @Success 201 {int} int
// @Router /turnos/{dni}/{matricula} [post]
func (h *turnoHandler) AddByDniMatricula(c *gin.Context) {
	dni, matricula := util.GetDniMatriculaFromParam(c)
	if dni == "" || matricula == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("parametro Dni o Matricula no encontrado"))
		return
	}

	var turno domain.TurnoAbstract
	if err := c.ShouldBindJSON(&turno); err != nil {
		web.Failure(c, http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	if turno.Fecha == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("fecha requerida"))
		return
	}

	if turno.Hora == "" {
		web.Failure(c, http.StatusBadRequest, errors.New("hora requerida"))
		return
	}

	dentista, err := h.serviceDentista.GetByMatricula(matricula)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}

	if dentista.Matricula != matricula {
		web.NotFound(c)
		return
	}

	paciente, err := h.servicePaciente.GetByDNI(dni)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	if paciente.DNI != dni {
		web.NotFound(c)
		return
	}

	turno.DentistaId = dentista.Id
	turno.PacienteId = paciente.Id

	id, err := h.service.Add(&turno)
	if err != nil {
		web.Failure(c, http.StatusInternalServerError, errors.New(err.Error()))
		return
	}
	web.Success(c, http.StatusCreated, id)
}
