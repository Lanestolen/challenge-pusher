package main

import (
	"errors"
	"os"
	"strings"

	"github.com/lanestolen/challenge-pusher/proto"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	ErrExerciseNotFound           = errors.New("exercise with name not found")
	ErrSolutionNotFound           = errors.New("solution for exercise not found")
	ErrTagMismatch                = errors.New("challenge tags does not match")
	ErrMissingFlag                = errors.New("missing flag for one of the exercises")
	ErrMissingImage               = errors.New("missing image for non static challenge")
	ErrMissingPoints              = errors.New("exercise should be given some points")
	ErrMissingDefaultInfo         = errors.New("missing tag, name, description or category")
	ErrMissingPrivacyUniverseInfo = errors.New("missing required info for privacy universe")

	defaultDir = "challenge-config/"
)

func validateConfig(exercise *proto.Exercise) error {
	exercises := make(map[string]bool)

	if exercise.Name == "" || exercise.Category == "" || exercise.Tag == "" || exercise.Description == "" {
		return ErrMissingDefaultInfo
	}

	if exercise.Category == "Privacy Universe" && len(exercise.Platforms) == 0 {
		return ErrMissingPrivacyUniverseInfo
	}

	for _, h := range exercise.Hosts {
		if !exercise.Static && h.Image == "" {
			return ErrMissingImage
		}

		for _, f := range h.Flags {
			exercises[f.Name] = false
			if strings.Split(f.Tag, "-")[0] != exercise.Tag {
				return ErrTagMismatch
			}

			if f.Env == "" && f.Static == "" {
				log.Error().Str("env", f.Env).Str("static", f.Static).Msg("Missing environment variable or static flag")
				return ErrMissingFlag
			}

			if f.Points == 0 {
				return ErrMissingPoints
			}
		}
	}
	for _, s := range exercise.Solutions {
		_, ok := exercises[s.ChallengeName]
		if !ok {
			log.Error().Str("solution name", s.ChallengeName).Msg("Solution without challenge found")
			return ErrExerciseNotFound
		}
		exercises[s.ChallengeName] = true
	}

	for ex, solution := range exercises {
		if !solution {
			log.Error().Str("challenge", ex).Msg("Solution missing for exercise")
			return ErrSolutionNotFound
		}
	}

	return nil
}

func parseConf(path string) (*proto.Exercise, error) {
	var exercise proto.Exercise

	fileContent, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, os.ErrNotExist
		}
		return nil, err
	}

	if err := yaml.Unmarshal(fileContent, &exercise); err != nil {
		return nil, err
	}

	return &exercise, nil
}

func getChallengeConfigs() ([]*proto.Exercise, error) {
	var challs []*proto.Exercise

	challFiles, err := os.ReadDir(defaultDir)
	if err != nil {
		return nil, err
	}
	for _, conf := range challFiles {
		chall, err := parseConf(defaultDir + conf.Name())
		if err != nil {
			return nil, err
		}
		if err := validateConfig(chall); err != nil {
			return nil, err
		}
		challs = append(challs, chall)
	}

	return challs, nil
}
