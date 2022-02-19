package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/frzam/affordability-mock/controllers"
)

func main() {
	http.HandleFunc("/api/v2/enquiry/consumer/affordability/save", saveHandler)
	http.ListenAndServe(":8080", nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {

	req := &controllers.AffordabilityRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Error in decoding: ", err)
	}
	log.Printf("Request:%+v ", *req)
	resp := &controllers.AffordabilityResponse{}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
