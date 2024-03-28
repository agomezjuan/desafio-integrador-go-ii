package domain

// Dentista representa a un dentista
type Dentista struct {
	Id        int    `json:"id"`
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}

type DentistaAbstract struct {
	Id        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
}
