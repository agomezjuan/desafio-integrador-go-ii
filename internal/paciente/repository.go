package paciente

import (
	"errors"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
	"github.com/agomezjuan/desafio-integrador-go-ii/pkg/store"
)

type Repository interface {
	// GetAll devuelve todos los pacientes
	GetAll() ([]domain.Paciente, error)
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Paciente, error)
	// GetByDNI busca un paciente por su DNI
	GetByDNI(dni string) (domain.Paciente, error)
	// Create agrega un nuevo paciente
	Create(p domain.Paciente) (domain.Paciente, error)
	// Update actualiza un paciente
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage store.PacienteStore
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.PacienteStore) Repository {
	return &repository{storage}
}

// GetAll devuelve todos los pacientes
func (r *repository) GetAll() ([]domain.Paciente, error) {
	pacientes, err := r.storage.GetAll()
	if err != nil {
		return nil, errors.New("no se encontraron pacientes")
	}

	return pacientes, nil
}

// GetByID busca un paciente por su id
func (r *repository) GetByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.Read(id)
	if err != nil {
		return domain.Paciente{}, errors.New("no se encontró al paciente")
	}

	return paciente, nil
}

// GetByDNI busca un paciente por su DNI
func (r *repository) GetByDNI(dni string) (domain.Paciente, error) {
	paciente, err := r.storage.ReadByDNI(dni)
	if err != nil {
		return domain.Paciente{}, errors.New("no se encontró al paciente")
	}

	return paciente, nil
}

// Create agrega un nuevo paciente
func (r *repository) Create(p domain.Paciente) (domain.Paciente, error) {
	if r.storage.Exists(p.DNI) {
		return domain.Paciente{}, errors.New("el paciente ya existe")
	}

	err := r.storage.Create(p)
	if err != nil {
		return domain.Paciente{}, errors.New("no se pudo agregar al paciente")
	}

	return p, nil
}

// Update actualiza un paciente
func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	if !r.storage.Exists(p.DNI) {
		return domain.Paciente{}, errors.New("el paciente ya existe")
	}

	err := r.storage.Update(p)
	if err != nil {
		return domain.Paciente{}, errors.New("no se pudo actualizar al paciente")
	}

	return p, nil
}

// Delete elimina un paciente
func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return errors.New("no se pudo eliminar al paciente")
	}

	return nil
}
