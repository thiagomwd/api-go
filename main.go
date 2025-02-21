package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
	Age  int    `json:"age_in_seconds"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Bem-vindo à nossa API!"}
	message.Age = 30 * 24 * 60 * 60
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", nil)
}
