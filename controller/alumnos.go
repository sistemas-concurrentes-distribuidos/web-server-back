package controller

import (
	"encoding/json"
	"net/http"
	"server/model"
	s "server/service"
	util "server/utils"

	"github.com/gorilla/mux"
)

//GatAllStudents the name of all students
func GatAllStudents(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, s.GetAllStudents())
}

//GatAllClasses the name of all classes
func GatAllClasses(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, s.GetAllClasses())
}

//AddRegister Add new register in the sistem
func AddRegister(w http.ResponseWriter, r *http.Request) {
	var body model.Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	s.AgregarCalificacionPorMateria(body)
	util.RespondWithJSON(w, http.StatusOK, body)
}

//GetStudentAvegare fasdfas
func GetStudentAvegare(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student := vars["name"]
	average := s.ObtenerPromedioAlumno(student)
	util.RespondWithJSON(w, http.StatusOK, map[string]float64{"average": average})
}

//GetGeneralAverage sdsdg
func GetGeneralAverage(w http.ResponseWriter, r *http.Request) {
	average := s.ObtenerPromedioAlumnos()
	util.RespondWithJSON(w, http.StatusOK, map[string]float64{"average": average})
}

//GetAverageByClass sdfasdf
func GetAverageByClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	class := vars["name"]
	average := s.ObtenerPromedioPorMateria(class)
	util.RespondWithJSON(w, http.StatusOK, map[string]float64{"average": average})
}
