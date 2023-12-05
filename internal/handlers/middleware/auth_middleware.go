package middleware

import (
	"context"
	"github.com/gofrs/uuid"
	"net/http"
)

type authKey struct{}

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		userID, err := uuid.FromString(auth)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		ctx := contextWithUserID(request.Context(), userID)
		newRequest := request.WithContext(ctx)

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
