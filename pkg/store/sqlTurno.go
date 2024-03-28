package store

import (
	"database/sql"
	/*"log"*/

	"github.com/agomezjuan/desafio-integrador-go-ii/internal/domain"
)

type sqlTurno struct {
	db *sql.DB
}

func NewSqlTurno(db *sql.DB) StoreTurno {
	return &sqlTurno{
		db: db,
	}
}

func (s *sqlTurno) Add(turno *domain.TurnoAbstract) (int, error) {
	res, err := s.db.Exec("INSERT INTO turnos (id_dentista, id_paciente, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?)", turno.DentistaId, turno.PacienteId, turno.Fecha, turno.Hora, turno.Descripcion)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	turno.Id = int(id)

	return turno.Id, nil
}

func (s *sqlTurno) Update(turno *domain.TurnoAbstract) error {
	_, err := s.db.Exec("UPDATE turnos SET dentista_id = ?, paciente_id = ?, fecha = ?, hora = ?, descripcion = ? WHERE id = ?", turno.DentistaId, turno.PacienteId, turno.Fecha, turno.Hora, turno.Descripcion, turno.Id)
	return err
}

func (s *sqlTurno) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM turnos WHERE id = ?", id)
	return err
}

func (s *sqlTurno) GetByID(id int) (*domain.Turno, error) {
	row := s.db.QueryRow("SELECT id, dentista_id, paciente_id, fecha, hora, descripcion FROM turnos WHERE id = ?", id)

	var turno domain.Turno
	err := row.Scan(&turno.Id, &turno.Dentista.Id, &turno.Paciente.Id, &turno.Fecha, &turno.Hora, &turno.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return &turno, nil
		}
		return &turno, err
	}

	return &turno, nil
}

func (s *sqlTurno) GetAll() ([]*domain.Turno, error) {
	rows, err := s.db.Query("SELECT id, dentista_id, paciente_id, fecha, hora, descripcion FROM turnos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	turnos := make([]*domain.Turno, 0)
	for rows.Next() {
		var turno domain.Turno
		err := rows.Scan(&turno.Id, &turno.Dentista.Id, &turno.Paciente.Id, &turno.Fecha, &turno.Hora, &turno.Descripcion)
		if err != nil {
			return nil, err
		}
		turnos = append(turnos, &turno)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return turnos, nil
}

func (s *sqlTurno) GetByDNI(dni string) (*domain.TurnoResponse, error) {
	row := s.db.QueryRow("SELECT turnos.id, CONCAT_WS(' ',dentistas.nombre, dentistas.apellido) AS dentista, CONCAT_WS(' ',pacientes.nombre, pacientes.apellido) AS paciente, turnos.fecha, turnos.hora, turnos.descripcion FROM turnos INNER JOIN pacientes ON (turnos.paciente_id = pacientes.id) INNER JOIN dentistas ON (turnos.dentista_id = dentistas.id) WHERE (pacientes.dni =?)", dni)

	var turno domain.TurnoResponse
	err := row.Scan(&turno.Id, &turno.Dentista, &turno.Paciente, &turno.Fecha, &turno.Hora, &turno.Descripcion)
	if err != nil {
		if err == sql.ErrNoRows {
			return &turno, nil
		}
		return &turno, err
	}

	return &turno, nil
}
