package docs

import (
	SharedVO "github.com/github-profile/go-swagger/shared/vo"
)

// swagger:route POST /api/transaction Transaction idOfTransactionEndpoint
// Transaction explanation.
// responses:
//   200: VirtualAccountResponse

// swagger:parameters idOfTransactionEndpoint
type TransactionParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body SharedVO.TransactionRequestDTO
}

// This text will appear as description of your response body.
// swagger:response TransactionResponse
type TransactionResponseWrapper struct {
	// in:body
	Body SharedVO.TransactionResponseDTO
}
