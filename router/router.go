package router

import (
	"net/http"
	c "server/controller"

	"github.com/gorilla/mux"
)

//InitRouter config the router and add all the endpoint
func InitRouter() *mux.Router {
	//Create the router
	r := mux.NewRouter()

	//Add endpoint
	r.HandleFunc("/student", c.GatAllStudents).Methods(http.MethodGet)
	r.HandleFunc("/student", c.AddRegister).Methods(http.MethodPost)
	r.HandleFunc("/class", c.GatAllClasses).Methods(http.MethodGet)
	r.HandleFunc("/average", c.GetGeneralAverage).Methods(http.MethodGet)
	r.HandleFunc("/average/student/{name}", c.GetStudentAvegare).Methods(http.MethodGet)
	r.HandleFunc("/average/class/{name}", c.GetAverageByClass).Methods(http.MethodGet)

	return r
}
