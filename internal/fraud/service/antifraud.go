package service

import (
	"strings"

	"Gateway/internal/fraud/models/"
)

func IsFraudulent(payment models.Payment) bool {
	// regras simples
	if payment.Amount > 10000 {
		return true
	}

	if strings.HasPrefix(payment.CardNumber, "0000") {
		return true
	}

	if payment.DeviceID == "" || payment.IpAddress == "" {
		return true
	}
	return false
}
