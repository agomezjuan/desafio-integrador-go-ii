package store

import (
	"database/sql"
	"log"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
)

type dentistaStore struct {
	db *sql.DB
}

func NewDentistaStore(db *sql.DB) DentistaStore {
	return &dentistaStore{
		db: db,
	}
}

// GetAll devuelve todos los dentistas
func (s *dentistaStore) GetAll() ([]domain.Dentista, error) {
	var dentistas []domain.Dentista
	query := "SELECT id, apellido, nombre, matricula FROM dentistas;"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dentista domain.Dentista
		err := rows.Scan(&dentista.Id, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
		if err != nil {
			return nil, err
		}
		dentistas = append(dentistas, dentista)
	}

	return dentistas, nil
}

// Read devuelve un dentista por su id
func (s *dentistaStore) Read(id int) (domain.Dentista, error) {
	var dentista domain.Dentista
	query := "SELECT id, apellido, nombre, matricula FROM dentistas WHERE id = ?;"
	err := s.db.QueryRow(query, id).Scan(&dentista.Id, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}

	return dentista, nil
}

// ReadByMatricula devuelve un dentista por su matricula
func (s *dentistaStore) ReadByMatricula(matricula string) (domain.Dentista, error) {
	var dentista domain.Dentista
	query := "SELECT id, apellido, nombre, matricula FROM dentistas WHERE matricula = ?;"
	err := s.db.QueryRow(query, matricula).Scan(&dentista.Id, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}

	return dentista, nil
}

// Create agrega un nuevo dentista
func (s *dentistaStore) Create(dentista domain.Dentista) error {
	query := "INSERT INTO dentistas (apellido, nombre, matricula) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Update actualiza un dentista
func (s *dentistaStore) Update(dentista domain.Dentista) error {
	query := "UPDATE dentistas SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula, dentista.Id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Delete elimina un dentista
func (s *dentistaStore) Delete(id int) error {
	query := "DELETE FROM dentistas WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Exists verifica si un dentista existe
func (s *dentistaStore) Exists(matricula string) bool {
	var exists bool
	var id int

	query := "SELECT id FROM dentistas WHERE matricula = ?;"
	err := s.db.QueryRow(query, matricula).Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
