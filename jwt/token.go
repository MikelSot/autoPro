package jwt

import (
	"errors"
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(data *dto.LoginClient) (string, error) {
	claim := dto.Claim{
		//ID:       data.ID,
		//Name:     data.Name,
		//LastName: data.LastName,
		Email:    data.Email,
		//State:    data.State,
		//Role:     data.Role,

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

func ValidateToken(t string) (dto.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &dto.Claim{}, verifyFunction)
	if err != nil {
		return dto.Claim{}, err
	}

	if !token.Valid{
		return dto.Claim{}, errors.New("token no valid")
	}

	// optener el claim de los claim
	claim , ok := token.Claims.(*dto.Claim)
	if !ok {
		return dto.Claim{}, errors.New("no se puedo optener los claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	// retorna la informacion de nuestro archivo publico
	return verifyKey, nil
}