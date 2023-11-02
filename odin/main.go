package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func giveThorMjolnir(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status    string `json:"status"`
		Message   string `json:"message"`
		CreatedBy string `json:"created_by"`
	}

	res := response{
		Status:    "OK",
		Message:   "Here is Mjölnir. You are now the God of thunder",
		CreatedBy: "The dwarves brothers, Eitri and Brokkr",
	}

	responseBody, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func home(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	res := response{
		Status:  "OK",
		Message: "I am Odin, the All-Father, and I grant you my wisdom as you traverse this digital realm. Seek knowledge and wield it with honor.",
	}

	responseBody, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

func main() {
	port := "3700"
	http.HandleFunc("/", home)
	http.HandleFunc("/gift_thor_hammer", giveThorMjolnir)
	fmt.Println("Serving on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
