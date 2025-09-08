package controller

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func AppInitialize() {
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "Origin"})
	originsOk := handlers.AllowedOrigins([]string{
		// development
		"http://127.0.0.1:8000",
		// production
		"https://aplikasijambi.kemenkum.go.id",
	})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})

	router := mux.NewRouter()
	router.HandleFunc("/provinsi", FetchProvinces).Methods(http.MethodGet)
	router.HandleFunc("/kabupaten/{provincecode}", FetchRegenciesByProvinceId).Methods(http.MethodGet)
	router.HandleFunc("/kecamatan/{regencycode}", FetchDistrictsByRegencyId).Methods(http.MethodGet)
	router.HandleFunc("/kelurahan/{districtcode}", FetchVillagesByDistrictId).Methods(http.MethodGet)

	http.ListenAndServe("127.0.0.1:8182", handlers.CORS(headersOk, originsOk, methodsOk)(router))
}
