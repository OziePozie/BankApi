package middleware

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
)

type authKey struct{}

type Claims struct {
	jwt.RegisteredClaims
	UserId uuid.UUID `json:"user"`
}

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")

		var claims Claims

		token, err := jwt.ParseWithClaims(auth, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(""), nil
		})

		if claims, ok := token.Claims.(*Claims); ok {

			log.Println(claims.UserId)
		}
		userID := claims.UserId

		log.Print(err)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		ctx := contextWithUserID(request.Context(), userID)
		newRequest := request.WithContext(ctx)

		log.Print(ctx.Value(authKey{}))

		handler.ServeHTTP(writer, newRequest)
	})

}

func contextWithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, authKey{}, userID)
}

func UserIDFromContext(ctx context.Context) uuid.UUID {
	userID, _ := ctx.Value(authKey{}).(uuid.UUID)

	return userID
}
