package dentista

import (
	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
)

type Service interface {
	// GetByID busca un dentista por su id
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(d domain.Dentista) (domain.Dentista, error)
	// Update actualiza un dentista
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	// Delete elimina un dentista
	Delete(id int) error
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetByID busca un dentista por su id
func (s *service) GetByID(id int) (domain.Dentista, error) {
	dentista, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}

	return dentista, nil
}

// Create agrega un nuevo dentista
func (s *service) Create(d domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}

	return d, nil
}

// Update actualiza un dentista
func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}

	if u.Apellido != "" {
		d.Apellido = u.Apellido
	}
	if u.Nombre != "" {
		d.Nombre = u.Nombre
	}
	if u.Matricula != "" {
		d.Matricula = u.Matricula
	}

	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Dentista{}, err
	}

	return d, nil
}

// Delete elimina un dentista
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}