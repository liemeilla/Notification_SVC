package entities

import "time"

type (
	EmptyResponse struct{}
)

// internal entities
type (
	Auth struct {
		CustomerId int       `json:"customer_id"`
		ApiKey     string    `json:"api_key"`
		Status     string    `json:"status"`
		CreatedAt  time.Time `json:"created_at"`
	}

	CustomerNotification struct {
		CustomerId      int       `json:"customer_id"`
		NotificationURL string    `json:"notification_url"`
		Status          string    `json:"status"`
		CreatedAt       time.Time `json:"created_at"`
	}

	NotificationLog struct {
		IdempotencyId    string    `json:"id"`
		CustomerId       int       `json:"customer_id"`
		NotificationUrl  string    `json:"notification_url"`
		NotificationData string    `json:"notification_data"`
		RequestJson      string    `json:"request_json"`
		ResponseJson     string    `json:"response_json"`
		StatusSent       string    `json:"status_sent"`
		CreatedAt        time.Time `json:"created_at"`
	}
)

// external entities
type (
	ReqPaymentNotification struct {
		Notification Notification `json:"notification"`
	}

	ResPaymentNotification struct {
		Status string `json:"status"`
	}

	Notification struct {
		TransactionTime time.Time `json:"transaction_time"`
		ReferenceID     string    `json:"reference_id"`
		TransactionID   string    `json:"transaction_id"`
		PaymentStatus   string    `json:"payment_status"`
		Currency        string    `json:"currency"`
		Amount          float64   `json:"amount"`
		ChannelCode     string    `json:"channel_code"`
	}
)
