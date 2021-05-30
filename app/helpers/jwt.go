package helpers

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Token struct {
	Hash   string
	Expire int64
}
type AuthConfiguration struct {
	App_Jwt_Secret string
	Api_Jwt_Secret string
	Jwt_Expire     int
}

var AuthConfig *AuthConfiguration //nolint:gochecknoglobals

func CreateToken(ctx *fiber.Ctx, userID uint) (Token, error) {
	var t Token
	secret := Config("SECRET_JWT")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = userID
	expiresIn := time.Now().Add(time.Duration(60*60*24) * time.Second).Unix()
	claims["exp"] = expiresIn

	tokenHash, err := token.SignedString([]byte(secret))
	if err != nil {
		return t, err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "userid",
		Value:    tokenHash,
		HTTPOnly: true,
	})
	t.Hash = tokenHash
	t.Expire = expiresIn
	return t, nil
}

func ExtraToken(ctx *fiber.Ctx) string {
	strArr := strings.Split(ctx.Get("Authorization"), " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func CookieToken(ctx *fiber.Ctx) string {
	tokenString := ctx.Cookies("userid")
	if tokenString == "" {
		return ""
	}

	return tokenString
}

func ParseToken(ctx *fiber.Ctx) (uint, error) {
	secret := Config("SECRET_JWT")
	tokenString := CookieToken(ctx)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	err2 := claims.Valid()

	if err2 != nil {
		DeleteToken(ctx)
		return 0, err2
	}

	return uint(claims["id"].(float64)), nil
}

//DeleteToken deletes the jwt token
func DeleteToken(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name: "userid",
		// Set expiry date to the past
		Expires: time.Now().Add(-(time.Hour * 2)),
	})
}

//RefreshToken refreshes the token
func RefreshToken(c *fiber.Ctx) (Token, error) {

	var t Token
	u, err := ParseToken(c)

	if err != nil {
		return t, err
	}
	return CreateToken(c, u)
}
