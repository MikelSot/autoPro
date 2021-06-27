package jwt

import (
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(data *dto.DataClient) (string, error) {
	claim := dto.Claim{
		ID:       data.ID,
		Name:     data.Name,
		LastName: data.LastName,
		Email:    data.Email,
		State:    data.State,
		Role:     data.Role,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "autoPRo",
		},
	}

	// preparar para firmar el token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	// se firma el token
	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
