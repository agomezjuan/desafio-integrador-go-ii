package paciente

import (
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
)

type Service interface {
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

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll devuelve todos los pacientes
func (s *service) GetAll() ([]domain.Paciente, error) {
	pacientes, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}

	return pacientes, nil
}

// GetByID busca un paciente por su id
func (s *service) GetByID(id int) (domain.Paciente, error) {
	paciente, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Create agrega un nuevo paciente
func (s *service) Create(p domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}

	return p, nil
}

// Update actualiza un paciente
func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}

	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.DNI != "" {
		p.DNI = u.DNI
	}

	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}

	return p, nil
}

// Delete elimina un paciente
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetByDNI(dni string) (domain.Paciente, error) {
	return s.r.GetByDNI(dni)
}
