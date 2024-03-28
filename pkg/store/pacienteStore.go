package store

import (
	"database/sql"
	"log"

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
)

type pacienteStore struct {
	db *sql.DB
}

func NewPacienteStore(db *sql.DB) PacienteStore {
	return &pacienteStore{
		db: db,
	}
}

// GetAll devuelve todos los pacientes
func (s *pacienteStore) GetAll() ([]domain.Paciente, error) {
	var pacientes []domain.Paciente
	query := "SELECT * FROM pacientes;"
	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var paciente domain.Paciente
		err := rows.Scan(&paciente.Id, &paciente.Apellido, &paciente.Nombre, &paciente.Domicilio, &paciente.DNI, &paciente.FechaAlta)
		if err != nil {
			log.Fatal(err)
		}
		pacientes = append(pacientes, paciente)
	}

	return pacientes, nil
}

// Read devuelve un paciente por su id
func (s *pacienteStore) Read(id int) (domain.Paciente, error) {
	var paciente domain.Paciente
	query := "SELECT id, apellido, nombre, dni FROM pacientes WHERE id = ?;"
	err := s.db.QueryRow(query, id).Scan(&paciente.Id, &paciente.Apellido, &paciente.Nombre, &paciente.DNI, &paciente.FechaAlta, &paciente.Domicilio)
	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// ReadByDNI devuelve un paciente por su dni
func (s *pacienteStore) ReadByDNI(dni string) (domain.Paciente, error) {
	var paciente domain.Paciente
	query := "SELECT id, apellido, nombre, dni, fecha_alta, domicilio FROM pacientes WHERE dni = ?;"
	err := s.db.QueryRow(query, dni).Scan(&paciente.Id, &paciente.Apellido, &paciente.Nombre, &paciente.DNI, &paciente.FechaAlta, &paciente.Domicilio)
	if err != nil {
		return domain.Paciente{}, err
	}

	return paciente, nil
}

// Create agrega un nuevo paciente
func (s *pacienteStore) Create(paciente domain.Paciente) error {
	query := "INSERT INTO pacientes (apellido, nombre, dni, domicilio, fecha_alta) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(paciente.Apellido, paciente.Nombre, paciente.DNI, paciente.Domicilio, paciente.FechaAlta)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Update actualiza un paciente
func (s *pacienteStore) Update(paciente domain.Paciente) error {
	query := "UPDATE pacientes SET apellido = ?, nombre = ?, dni = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(paciente.Apellido, paciente.Nombre, paciente.DNI, paciente.Id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Delete elimina un paciente
func (s *pacienteStore) Delete(id int) error {
	query := "DELETE FROM pacientes WHERE id = ?;"
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

// Exists verifica si un paciente existe
func (s *pacienteStore) Exists(dni string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM pacientes WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
