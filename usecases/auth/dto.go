package auth_usecase

type LoginSpec struct {
	Email    string
	Password string
}

type LoginResult struct {
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type SignUpSpec struct {
	Name     string
	Email    string
	Password string
}

type SignUpResult struct {
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
}
