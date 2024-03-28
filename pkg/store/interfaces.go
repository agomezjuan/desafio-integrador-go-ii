package store

import "github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"

type DentistaStore interface {
	// GetAll devuelve todos los dentistas
	GetAll() ([]domain.Dentista, error)
	// Read devuelve un dentista por su id
	Read(id int) (domain.Dentista, error)
	// ReadByMatricula devuelve un dentista por su matricula
	ReadByMatricula(matricula string) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(dentista domain.Dentista) error
	// Update actualiza un dentista
	Update(dentista domain.Dentista) error
	// Delete elimina un dentista
	Delete(id int) error
	// Exists verifica si un dentista existe
	Exists(matricula string) bool
}

type PacienteStore interface {
	// GetAll devuelve todos los pacientes
	GetAll() ([]domain.Paciente, error)
	// Read devuelve un paciente por su id
	Read(id int) (domain.Paciente, error)
	// ReadByDNI devuelve un paciente por su DNI
	ReadByDNI(dni string) (domain.Paciente, error)
	// Create agrega un nuevo paciente
	Create(paciente domain.Paciente) error
	// Update actualiza un paciente
	Update(paciente domain.Paciente) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(dni string) bool
}


type StoreTurno interface {
	Add(turno *domain.TurnoAbstract) (int, error)
	Update(turno *domain.TurnoAbstract) error
	Delete(id int) error
	GetByID(id int) (*domain.Turno, error)
	GetAll() ([]*domain.Turno, error)
	GetByDNI(dni string) (*domain.TurnoResponse, error)
}
