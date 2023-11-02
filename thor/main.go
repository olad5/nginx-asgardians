package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getMjolnir(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3700/gift_thor_hammer", nil)
	if err != nil {
		panic(err)
	}

	odinResponse := make(map[string]interface{})

	res, err := http.DefaultClient.Do(req)
	if err := json.NewDecoder(res.Body).Decode(&odinResponse); err != nil {
		panic(err)
	}

	type response struct {
		Status          string `json:"status"`
		Message         string `json:"message"`
		MjolnirCreators string `json:"mjolnir_creators"`
	}

	createdBy := odinResponse["created_by"].(string)

	resp := response{
		Status:          "OK",
		Message:         "Received Mjölnir from Odin",
		MjolnirCreators: createdBy,
	}

	responseBody, err := json.Marshal(resp)
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
		Message: `I, Thor, the god of thunder, bid you welcome! May your journey through this digital realm be as epic as the storm I command in the heavens.`,
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
	port := "3900"
	http.HandleFunc("/get_mjolnir", getMjolnir)
	http.HandleFunc("/", home)
	fmt.Println("Serving on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
