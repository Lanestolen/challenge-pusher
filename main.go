package main

import (
	"context"
	"errors"

	"os"
	"strconv"

	"github.com/lanestolen/challenge-pusher/proto"
	"github.com/rs/zerolog/log"
)

var (
	ErrMissingEnvironment = errors.New("missing parameters needed for connecting to the exercise database")
	ErrInvalidFormat      = errors.New("environment set is and invalid format")
)

func main() {
	conf, err := readConfigFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to find required environments")
	}

	client, err := NewExServiceConn(*conf)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to exercise database")
	}

	challenges, err := getChallengeConfigs()
	if err != nil {
		log.Fatal().Err(err).Msg("failed get challenge configs from dir")
	}
	if _, err := client.AddExercises(context.TODO(), &proto.AddExercisesRequest{Exercises: challenges}); err != nil {
		log.Fatal().Err(err).Msg("failed to insert/update challenges")
	}
}

type config struct {
	Endpoint string
	Port     uint64
	AuthKey  string
	SignKey  string
}

func readConfigFromEnv() (*config, error) {
	var conf config

	conf.Endpoint = os.Getenv("ENDPOINT")
	if conf.Endpoint == "" {
		return nil, ErrMissingEnvironment
	}

	port, err := strconv.ParseUint(os.Getenv("PORT"), 0, 64)
	if err != nil {
		return nil, ErrInvalidFormat
	}
	conf.Port = port

	conf.AuthKey = os.Getenv("AUTH_KEY")
	if conf.AuthKey == "" {
		return nil, ErrMissingEnvironment
	}

	conf.SignKey = os.Getenv("SIGN_KEY")
	if conf.SignKey == "" {
		return nil, ErrMissingEnvironment
	}

	return &conf, nil
}
