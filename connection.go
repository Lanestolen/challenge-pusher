package main

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/lanestolen/challenge-pusher/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Creds struct {
	Token    string
	Insecure bool
}

func (c Creds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"token": string(c.Token),
	}, nil
}

func (c Creds) RequireTransportSecurity() bool {
	return !c.Insecure
}

func NewExServiceConn(config config) (proto.ExerciseStoreClient, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"au": config.AuthKey,
	})
	tokenString, err := token.SignedString([]byte(config.SignKey))
	if err != nil {
		return nil, err
	}

	authCreds := Creds{Token: tokenString, Insecure: true}

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})),
		grpc.WithPerRPCCredentials(authCreds),
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", config.Endpoint, config.Port), dialOpts...)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to challenge store")
		return nil, err
	}
	c := proto.NewExerciseStoreClient(conn)
	return c, nil

}
