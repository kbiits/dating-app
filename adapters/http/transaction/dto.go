package http_transaction

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BuyRequest struct {
	PackageID string `json:"package_id" validate:"required"`
}
