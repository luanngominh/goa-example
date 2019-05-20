package jwt

import (
	"context"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	uuid "github.com/satori/go.uuid"

	"github.com/luanngominh/goa-example/app"
	"github.com/luanngominh/goa-example/ext/config"
)

// NewJWTMiddleware ...
func NewJWTMiddleware(jwtSecurity *goa.JWTSecurity) (goa.Middleware, error) {
	validationHandler, _ := goa.NewMiddleware(func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		token := jwt.ContextJWT(ctx)
		_, ok := token.Claims.(jwtgo.MapClaims)
		if !ok {
			return jwt.ErrJWTError("unsupported claims shape")
		}
		return nil
	})

	return jwt.New(config.Cfg.SignKey, validationHandler, app.NewJWTSecurity()), nil
}

//GenerateJWT ...
func GenerateJWT(userID, email string) *jwtgo.Token {
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	oneMonth := time.Now().Add(time.Duration(24*30) * time.Hour).Unix()
	uuid := uuid.NewV4()

	token.Claims = jwtgo.MapClaims{
		"iss":        "Issuer",          // who creates the token and signs it
		"aud":        "Audience",        // to whom the token is intended to be sent
		"exp":        oneMonth,          // time when the token will expire (10 minutes from now)
		"jti":        uuid.String(),     // a unique identifier for the token
		"iat":        time.Now().Unix(), // when the token was issued/created (now)
		"nbf":        2,                 // time before which the token is not yet valid (2 minutes ago)
		"sub":        "llra api token",  // the subject/principal is whom the token is about
		"user.id":    userID,            // user id - not a standard claim
		"user.email": email,             // user email - not a standard claim
		"scopes":     "api:access company_account:true",
	}

	return token
}

func GenerateJWTToken(token *jwtgo.Token) string {
	return ""
}
