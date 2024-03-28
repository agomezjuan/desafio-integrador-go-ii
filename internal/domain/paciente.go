package domain

// Paciente representa a un dentista
type Paciente struct {
	Id        int    `json:"id"`
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Domicilio string `json:"domicilio" binding:"required"`
	DNI       string `json:"dni" binding:"required"`
	FechaAlta string `json:"fecha_alta" binding:"required"`
}

type PacienteAbstract struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Domicilio string `json:"domicilio"`
	DNI       string `json:"dni"`
	FechaAlta string `json:"fecha_alta"`
}
