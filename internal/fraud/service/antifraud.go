package service

import (
	"strings"
	"time"

	"gateway/internal/fraud/models"
)

// Global history storage payment card
var cardPaymentHistory = map[string][]time.Time{}

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

	if isCardUsedTooFrequently(payment) {
		return true
	}
	return false
}

// Check if card was used too many times in a short time
func isCardUsedTooFrequently(payment models.Payment) bool {

	now := time.Now()

	// Set time windows
	const timeWindow = 1 * time.Minute

	// Set max allowed number of payments in window
	const maxAllowedPayments = 5

	// Get card history
	times := cardPaymentHistory[payment.CardNumber]

	// Sort payments inside time window
	var recent []time.Time
	for _, t := range times {
		if now.Sub(t) <= timeWindow {
			recent = append(recent, t)
		}
	}

	// Add payment in history
	recent = append(recent, now)
	cardPaymentHistory[payment.CardNumber] = recent

	// Check if have exceeded the limit
	return len(recent) > maxAllowedPayments
}
