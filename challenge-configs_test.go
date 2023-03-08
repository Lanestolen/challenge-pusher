package main

import (
	"testing"
)

func Test_validateConfig(t *testing.T) {
	tests := []struct {
		name         string
		exerciseFile string
		wantErr      error
	}{
		{
			name:         "Success - static one flag",
			exerciseFile: "testdata/success-static-1-flag.yml",
		},
		{
			name:         "Success - static two flags",
			exerciseFile: "testdata/success-static-2-flags.yml",
		},
		{
			name:         "fail - missing flag",
			exerciseFile: "testdata/fail-missing-flag.yml",
			wantErr:      ErrExerciseNotFound,
		},
		{
			name:         "fail - missing solution",
			exerciseFile: "testdata/fail-missing-solution.yml",
			wantErr:      ErrSolutionNotFound,
		},
		{
			name:         "fail - missing tag",
			exerciseFile: "testdata/fail-missing-tag.yml",
			wantErr:      ErrMissingDefaultInfo,
		},
		{
			name:         "fail - missing category",
			exerciseFile: "testdata/fail-missing-category.yml",
			wantErr:      ErrMissingDefaultInfo,
		},
		{
			name:         "fail - missing name",
			exerciseFile: "testdata/fail-missing-name.yml",
			wantErr:      ErrMissingDefaultInfo,
		},
		{
			name:         "fail - missing description",
			exerciseFile: "testdata/fail-missing-description.yml",
			wantErr:      ErrMissingDefaultInfo,
		},
		{
			name:         "fail - missing platforms",
			exerciseFile: "testdata/fail-missing-privacyuniverse.yml",
			wantErr:      ErrMissingPrivacyUniverseInfo,
		},
		{
			name:         "success - privacy universe",
			exerciseFile: "testdata/success-privacyuniverse.yml",
		},
		{
			name:         "success - dynamic",
			exerciseFile: "testdata/success-dynamic.yml",
		},
		{
			name:         "success - dynamic multiple hosts",
			exerciseFile: "testdata/success-dynamic-multiple.yml",
		},
		{
			name:         "fail - missing image",
			exerciseFile: "testdata/fail-missing-image.yml",
			wantErr:      ErrMissingImage,
		},
		{
			name:         "fail - missing flag env",
			exerciseFile: "testdata/fail-missing-flagdef.yml",
			wantErr:      ErrMissingFlag,
		},
		{
			name:         "fail - tag mismatch",
			exerciseFile: "testdata/fail-tag-mismatch.yml",
			wantErr:      ErrTagMismatch,
		},
		{
			name:         "fail - missing points",
			exerciseFile: "testdata/fail-missing-points.yml",
			wantErr:      ErrMissingPoints,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			challengeConf, err := parseConf(tt.exerciseFile)
			if err != nil {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := validateConfig(challengeConf); err != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
