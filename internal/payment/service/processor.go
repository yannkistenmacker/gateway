package service

import (
	"fmt"

	fraudModels "gateway/internal/fraud/models"

	"gateway/internal/fraud/service"
)

func ProcessPayment(payment fraudModels.Payment) error {
	if service.IsFraudulent(payment) {
		return fmt.Errorf("Pagamento identificado como fraude")
	}

	// Chamada para adquirente
	fmt.Println("Pagamento processado com sucesso")
	return nil
}
