package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/banking/service"
)

// type Customer struct {
// 	Name    string `json:"name"`
// 	City    string `json:"city"`
// 	Pincode int    `json:"pincode"`
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, world!")
// }

type CustomerHandlers struct {
	Service service.CustomerService
}

func (h *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	// customers := []Customer{
	// 	{"Akash", "Mumbai", 400482},
	// 	{"Rahul", "Delhi", 200482},
	// }

	customers, _ := h.Service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
