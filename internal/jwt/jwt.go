package jwt

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	*jwt.StandardClaims
	CityId    uint     `json:"cityId"`
	City      string   `json:"city"`
	Roles     []string `json:"roles"`
	Authority string   `json:"authority"`
}

func parseToken(token string, secret []byte) (*CustomClaims, error) {
	claims := &CustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return claims, err
}
