package routes

import(
	"github.com/gorilla/mux"
	"github.com/rootturk/words-storage/pkg/controllers"
)

var RegisterWordRoutes = func(router *mux.Router){
	router.HandleFunc("/dictionary/", controllers.CreateDictionary).Methods("POST")
	router.HandleFunc("/dictionary/", controllers.GetDictionaries).Methods("GET")
	router.HandleFunc("/dictionary/{id}", controllers.GetDictionaryById).Methods("GET")
	router.HandleFunc("/dictionary/{id}", controllers.UpdateDictionary).Methods("PUT")
	router.HandleFunc("/dictionary/{id}", controllers.DeleteDictionary).Methods("DELETE")
}