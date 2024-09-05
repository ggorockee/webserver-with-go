package serializers

type LoginSerializer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
