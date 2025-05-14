package models

type Payment struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	CardNumber   string  `json:"card_number"`
	CardHolder   string  `json:"card_holder"`
	CardExpiry   string  `json:"card_expiry"`
	CVV          string  `json:"cvv"`
	IpAddress    string  `json:"ip_address"`
	DeviceID     string  `json:"device_id"`
	IsFraudulent string  `json:"is_fraudulent"`
}
