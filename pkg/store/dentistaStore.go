package store

import (
	"database/sql"
	"log"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
)

type dentistaStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) DentistaStoreInterface {
	return &dentistaStore{
		db: db,
	}
}

func (s *dentistaStore) Read(id int) (domain.Dentista, error) {
	var dentista domain.Dentista
	query := "SELECT id, apellido, nombre, matricula FROM dentistas WHERE id = ?;"
	err := s.db.QueryRow(query, id).Scan(&dentista.ID, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
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

// Update actualiza un producto
func (s *dentistaStore) Update(product domain.Dentista) error {
	query := "UPDATE dentistas SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(product.Apellido, product.Nombre, product.Matricula, product.ID)
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	
	return nil
}

// Delete elimina un producto
func (s *dentistaStore) Delete(id int) error {
	return nil
}

// Exists verifica si un producto existe
func (s *dentistaStore) Exists(codeValue string) bool {
	return true
}
