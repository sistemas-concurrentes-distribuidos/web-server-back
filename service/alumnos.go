package service

import (
	"errors"
	"fmt"
	"server/model"
)

var materias = make(map[string]map[string]float64)
var alumnos = make(map[string]map[string]float64)

//AgregarCalificacionPorMateria para agregar registros
func AgregarCalificacionPorMateria(body model.Body) error {
	alumno := alumnos[body.Alumno]
	if alumno == nil {
		alumno = make(map[string]float64)
	}
	if alumno[body.Materia] != 0 {
		return errors.New("La calificación ya fue asignada")
	}
	alumno[body.Materia] = body.Calificacion
	alumnos[body.Alumno] = alumno

	materia := materias[body.Materia]
	if materia == nil {
		materia = make(map[string]float64)
	}
	materia[body.Alumno] = body.Calificacion
	materias[body.Materia] = materia
	return nil
}

//GetAllStudents get all the names of the students
func GetAllStudents() []string {
	keys := make([]string, 0, len(alumnos))
	for k := range alumnos {
		keys = append(keys, k)
	}
	return keys
}

//GetAllClasses get all the names of the classes
func GetAllClasses() []string {
	keys := make([]string, 0, len(materias))
	for k := range materias {
		keys = append(keys, k)
	}
	return keys
}

//ObtenerPromedioAlumno para obtener el promedio de un alumno
func ObtenerPromedioAlumno(alumno string) float64 {
	return calcularPromedioDeAlumno(alumno)
}

//ObtenerPromedioAlumnos para obtener el promedio general
func ObtenerPromedioAlumnos() float64 {
	var total float64
	for alumno := range alumnos {
		total += calcularPromedioDeAlumno(alumno)
	}
	promedio := total / float64(len(alumnos))
	fmt.Println("total: ", total)
	fmt.Println("alumnos: ", float64(len(alumnos)))
	fmt.Println("promedio: ", promedio)
	fmt.Println("******************************")
	return promedio
}

//ObtenerPromedioPorMateria para obtener el promedio por materia
func ObtenerPromedioPorMateria(materia string) float64 {
	var total float64
	for _, calificacion := range materias[materia] {
		total += calificacion
	}
	promedio := total / float64(len(materias[materia]))
	fmt.Println("Materia: ", materia)
	fmt.Println("total: ", total)
	fmt.Println("calificaciones: ", float64(len(materias[materia])))
	fmt.Println("promedio: ", promedio)
	fmt.Println("-----------------------------")
	return promedio
}

func calcularPromedioDeAlumno(alumno string) float64 {
	var total float64 = 0.0
	for _, calificacion := range alumnos[alumno] {
		total = total + calificacion
	}

	promedio := total / float64(len(alumnos[alumno]))
	fmt.Println("Alumno: ", alumno)
	fmt.Println("total: ", total)
	fmt.Println("materias: ", float64(len(alumnos[alumno])))
	fmt.Println("promedio: ", promedio)
	fmt.Println("-----------------------------")
	return float64(promedio)
}
