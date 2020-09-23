package vo

type (
	// TransactionRequestDTO ...
	TransactionRequestDTO struct {
		TransactionID string `json:"transaction_id"`
		Name          string `json:"name"`
		PayerID       string `json:"payer_id"`
		PayerEmail    string `json:"payer_email" validate:"required"`
		PayerType     string `json:"payer_type"`
		PartnerID     string `json:"partner_id"`
		PaymentType   string `json:"payment_type"`
		ProductType   string `json:"product_type"`
		Flag          string `json:"flag" validate:"required"`
		Channel       string `json:"channel" validate:"required"`
		Country       string `json:"country" validate:"required"`
		Currency      string `json:"currency" validate:"required"`
	}

	// TransactionResponseDTO ...
	TransactionResponseDTO struct {
		Status  string      `json:"status,omitempty"`
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message,omitempty"`
	}
)
