package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	res := response{
		Status:  "OK",
		Message: "Oh, how amusing it is to have caught your attention, mortal. Loki, the god of mischief, welcomes you to this digital realm of tricks and marvels!",
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
	port := "3500"
	http.HandleFunc("/", home)
	fmt.Println("Serving on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
