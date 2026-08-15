package jwtutils

import (
	"path/filepath"
	"testing"

	"github.com/HemlockPham7/common-libs/pkg/errorutils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewJWTValidator(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		publicKeyPath string
		expectedErr   bool
	}{
		{
			name:          "case with invalid public key path",
			publicKeyPath: filepath.FromSlash("./invalid_public_key.pem"),
			expectedErr:   true,
		},
		{
			name:          "case with valid public key path",
			publicKeyPath: filepath.FromSlash("./public_key.test.pem"),
			expectedErr:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			jwtValidator, err := NewJWTValidator(tc.publicKeyPath)
			if tc.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, jwtValidator)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, jwtValidator)
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	t.Parallel()

	JWTValidator, err := NewJWTValidator("./public_key.test.pem")
	require.NoError(t, err)

	testCases := []struct {
		name string

		inputTokenStr string

		expectedErr    error
		expectedClaims jwt.MapClaims
	}{
		{
			name: "case with fully jwt claims",

			inputTokenStr: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsIm5hbWUiOiJUZXN0IFVzZXIiLCJzdWIiOiIxMjM0NTY3ODkifQ.GFZZptrUg_308WUc0imbUqgbvkvmfmrQs02sC3uHlxAKFw6p72qYfCuse3jGxeLp8npNJbDHXqUQNnOcbbCKQtg7ltd6WHFwblnnAiV6m5jKG-sFQPLKe0YMcooDvFd13VdqrMdpO-njWIq7cJ4CHlS1WWXgabM1pfbCsM61IhAy9VuEY63jqYy-uxla8tRDNspEte7InsgD65zHNy5MK57qURQJx4qHQKKwHZefyDceZvYhLrCvHMom5FnHqul-tEJN4QKMrSgFTDOwxpHA4IySC_GrNyW63-3-If913BFcRDxMf87EeZ6PZPVH2GRwMH3Xd8_QyQ_uJw3iiSQq2A",

			expectedErr: nil,
			expectedClaims: jwt.MapClaims{
				"sub":   "123456789",
				"name":  "Test User",
				"admin": false,
			},
		},
		{
			name: "case with fully jwt claims",

			inputTokenStr: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.e30.jzySL95DktCKhDz-lHzPbuCmNB-kc6NXq8liILtdfa4ZnzT33653OyixO4NIeDMLLF8dFIXHzrodcJbH7dUKcmSibddOG999AuDetB8Xp51Bv3GNHSxhYF6xi9_HGtZsO0oKkS5fbRjQWm98MX170hGPCfAdUN08-Kr4UCovgLLdsAWQ5BccE4HP47ktOVLJEWmH8veKDnyunEgvFGGx3HN--H8hGBzozNfBwpqA4UJhGT8u5eTVqSl6VYWHPdHo8WgOWq7VJjEx_DCI6_20cvzMXCHxYSeTYpyKDn5rUL-iTKK4yldWLG_FN22DW31du9xiegSV6qEqLEiCLj5yLg",

			expectedErr:    nil,
			expectedClaims: jwt.MapClaims{},
		},
		{
			name: "invalid token",

			inputTokenStr: "abdfac.dfaf.231",

			expectedErr:    errorutils.ErrInvalidToken,
			expectedClaims: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			claims, err := JWTValidator.ValidateJWT(tc.inputTokenStr)

			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedClaims, claims)
		})
	}
}
