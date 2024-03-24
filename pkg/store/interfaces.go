package store

import "github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"

type DentistaStoreInterface interface {
	// Read devuelve un dentista por su id
	Read(id int) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(dentista domain.Dentista) error
	// Update actualiza un dentista
	Update(dentista domain.Dentista) error
	// Delete elimina un dentista
	Delete(id int) error
	// Exists verifica si un dentista existe
	Exists(codeValue string) bool
}
