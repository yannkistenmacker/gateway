package handler

import (

	"encoding/json"
	"net/http"
	"github.com/yannkistenmacker/gateway/internal/fraud/models"
	"github.com/yannkistenmacker/gateway/internal/fraud/service"

)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var payment fraudModels.payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Erro ao decodificar o pagamento", http.StatusBadRequest)
		return
	}

	err = service.ProcessPayment(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pagamento aceito"))
}
