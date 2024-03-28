package turno

import (
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/store"
)

type Repository interface {
	Add(turno *domain.TurnoAbstract) (int, error)
	Update(turno *domain.TurnoAbstract) error
	Delete(id int) error
	GetByID(id int) (*domain.Turno, error)
	GetAll() ([]*domain.Turno, error)
	GetByDNI(dni string) (*domain.TurnoResponse, error)
}

type repository struct {
	storage store.StoreTurno
}

func NewRepository(storage store.StoreTurno) Repository {
	return &repository{storage}
}

func (r *repository) Add(turno *domain.TurnoAbstract) (int, error) {
	return r.storage.Add(turno)
}

func (r *repository) Update(turno *domain.TurnoAbstract) error {
	return r.storage.Update(turno)
}

func (r *repository) Delete(id int) error {
	return r.storage.Delete(id)
}

func (r *repository) GetByID(id int) (*domain.Turno, error) {
	return r.storage.GetByID(id)
}

func (r *repository) GetAll() ([]*domain.Turno, error) {
	return r.storage.GetAll()
}

func (r *repository) GetByDNI(dni string) (*domain.TurnoResponse, error) {
	return r.storage.GetByDNI(dni)
}
