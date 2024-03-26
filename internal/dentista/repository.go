package dentista

import (
	"errors"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/store"
)

type Repository interface {
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo producto
	Create(p domain.Dentista) (domain.Dentista, error)
	// Update actualiza un producto
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	// Delete elimina un producto
	Delete(id int) error
}

type repository struct {
	storage store.DentistaStoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.DentistaStoreInterface) Repository {
	return &repository{storage}
}

// GetByID busca un dentista por su id
func (r *repository) GetByID(id int) (domain.Dentista, error) {
	dentista, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, errors.New("no se encontró al dentista")
	}

	return dentista, nil
}

// Create agrega un nuevo dentista
func (r *repository) Create(d domain.Dentista) (domain.Dentista, error) {
	if !r.storage.Exists(d.Matricula) {
		return domain.Dentista{}, errors.New("ya existe un dentista con esa matrícula")
	}
	err := r.storage.Create(d)
	if err != nil {
		return domain.Dentista{}, errors.New("no se pudo agregar al dentista")
	}

	return d, nil
}

// Update actualiza un dentista
func (r *repository) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	if !r.storage.Exists(d.Matricula) {
		return domain.Dentista{}, errors.New("ya existe un dentista con esa matrícula")
	}
	err := r.storage.Update(d)
	if err != nil {
		return domain.Dentista{}, errors.New("no se pudo actualizar al dentista")
	}

	return d, nil
}

// Delete elimina un dentista
func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return errors.New("no se pudo eliminar al dentista")
	}

	return nil
}