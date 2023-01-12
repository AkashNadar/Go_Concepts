package app

import (
	"log"
	"net/http"

	"github.com/banking/domain"
	"github.com/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	// Wiring
	// handler needs service and service needs repo
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// Routes
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// Server
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
