package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"gitlab.com/nbdgocean6/clockwerk/gvar"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
	"strings"
)

func BasicAuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.New("authentication required")
			}
			authData, ok := md["authorization"]
			if !ok {
				return nil, errors.New("authentication required")
			}
			var username, password string
			var basicAuth, data []string
			if basicAuth = strings.Split(authData[0], " "); len(basicAuth) != 2 {
				return nil, errors.New("authentication required")
			}
			if data = strings.Split(basicAuth[1], ":"); len(basicAuth) != 2 {
				return nil, errors.New("authentication required")
			}
			username = data[0]
			password = data[1]

			key := fmt.Sprintf("%s_%s", gvar.HashKeyMap, username)
			hashedPassword, ok := gvar.SyncMapHashStorage.Load(key)
			if !ok {
				return nil, errors.New("username not found")
			}
			err := bcrypt.CompareHashAndPassword([]byte(hashedPassword.(string)), []byte(password))
			if err != nil {
				return nil, err
			}
			return next(ctx, request)
		}
	}
}
