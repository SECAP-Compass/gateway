package jwt

type Handler interface {
	FetchClaims(token string) (*CustomClaims, error)
}

type jwtHandler struct {
	secret []byte
}

func NewJwtHandler() Handler {
	// Should read secret from file.
	return &jwtHandler{secret: []byte("secret")}
}

func (j *jwtHandler) FetchClaims(token string) (*CustomClaims, error) {
	return parseToken(token, j.secret)
}
