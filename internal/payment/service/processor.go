package service

import (
	"fmt"

	fraudModels "github.com/yannkistenmacker/gateway/internal/fraud/models"
	"github.com/yannkistenmacker/gateway/internal/fraud/service"
)

func ProcessPayment(payment fraudModels.Payment) error {
	if service.IsFraudulent(payment) {
		return fmt.Errorf("Pagamento identificado como fraude")
	}

	// Chamada para adquirente
	fmt.Println("Pagamento processado com sucesso")
	return nil
}
