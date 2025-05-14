package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/payment", handler.PaymentHandler)

	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
