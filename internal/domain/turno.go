package domain

type Turno struct {
	Id          int      `json:"id"`
	Dentista    Dentista `json:"dentista" binding:"required"`
	Paciente    Paciente `json:"paciente_id" binding:"required"`
	Fecha       string   `json:"fecha" binding:"required"`
	Hora        string   `json:"hora" binding:"required"`
	Descripcion string   `json:"descripcion"`
}

type TurnoAbstract struct {
	Id          int    `json:"id"`
	DentistaId  int    `json:"dentista_id"`
	PacienteId  int    `json:"paciente_id"`
	Fecha       string `json:"fecha"`
	Hora        string `json:"hora"`
	Descripcion string `json:"descripcion"`
}

type TurnoResponse struct {
	Id          int
	Paciente    string
	Dentista    string
	Fecha       string
	Hora        string
	Descripcion string
}
