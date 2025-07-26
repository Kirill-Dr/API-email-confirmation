package verify

type EmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyRequest struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}
