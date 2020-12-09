package model

//Body to handle server petitions
type Body struct {
	Alumno       string  `json:"alumno"`
	Materia      string  `json:"materia"`
	Calificacion float64 `json:"calificacion"`
}
